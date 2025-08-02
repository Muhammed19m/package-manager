package pkgmgr

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/Masterminds/semver"
	"github.com/appleboy/easyssh-proxy"
	"github.com/mholt/archives"
)

var (
	ErrNameEmpty     = errors.New("имя пустое")
	ErrVerInvalid    = errors.New("неправильная версия")
	ErrTargetsEmpty  = errors.New("нет файлов для упаковки")
	ErrCreateArchive = errors.New("не удалось создать архив")
)

func CreatePackage(sshCfg SshConfig, pkgInfo PackageInfo) error {
	pkgInfo.Name = strings.TrimSpace(pkgInfo.Name)
	if pkgInfo.Name == "" {
		return ErrNameEmpty
	}

	_, err := semver.NewVersion(pkgInfo.Ver)
	if err != nil {
		slog.Error(err.Error())
		return ErrVerInvalid
	}

	if len(pkgInfo.Targets) == 0 {
		return ErrTargetsEmpty
	}

	var allFileNames []string
	for _, target := range pkgInfo.Targets {
		fileNames, err := filenamesByTarget(target)
		if err != nil {
			return err
		}
		allFileNames = append(allFileNames, fileNames...)
	}

	archiveName := pkgInfo.Name + "-" + pkgInfo.Ver + ".tar"
	archiveAbs := filepath.Join(os.TempDir(), "pkgmgr", fmt.Sprint(time.Now().Unix()), archiveName)

	if err = createArchive(allFileNames, archiveAbs); err != nil {
		return ErrCreateArchive
	}
	// defer os.Remove(archiveAbs)

	ssh := &easyssh.MakeConfig{
		User:     sshCfg.User,
		Server:   sshCfg.Server,
		Password: sshCfg.Passwd,
		Port:     sshCfg.Port,
	}

	remoteTargetAbs := path.Join(sshCfg.PackagesDir, archiveName)

	if err = ssh.Scp(archiveAbs, remoteTargetAbs); err != nil {
		return err
	}

	return nil
}

func filenamesByTarget(target Target) ([]string, error) {
	var res []string
	filenames, err := filepath.Glob(target.Path)
	if err != nil {
		return nil, err
	}

	for _, name := range filenames {
		exludeClean := filepath.Clean(target.Exclude)
		matched, err := filepath.Match(exludeClean, name)
		if err != nil {
			return nil, err
		}
		if !matched {
			res = append(res, name)
		}
	}

	return res, nil
}

func createArchive(filenames []string, outputArchive string) error {
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}

	if err = os.MkdirAll(filepath.Dir(outputArchive), 0755); err != nil {
		return err
	}

	archiveFilenames := make(map[string]string)
	for _, name := range filenames {
		fileAbs, err := filepath.Abs(name)
		if err != nil {
			return err
		}
		archiveFilenames[fileAbs], err = filepath.Rel(currentDir, fileAbs)
		if err != nil {
			return err
		}
		if runtime.GOOS == "windows" {
			archiveFilenames[fileAbs] = filepath.ToSlash(archiveFilenames[fileAbs])
		}
	}

	ctx := context.Background()

	// map files on disk to their paths in the archive using default settings (second arg)
	var fileInfos []archives.FileInfo
	fileInfos, err = archives.FilesFromDisk(ctx, nil, archiveFilenames)
	if err != nil {
		return err
	}

	// create the output file we'll write to
	out, err := os.Create(outputArchive)
	if err != nil {
		return err
	}
	defer out.Close()

	// create the archive
	if err = (archives.Tar{}).Archive(ctx, out, fileInfos); err != nil {
		return err
	}

	return nil
}
