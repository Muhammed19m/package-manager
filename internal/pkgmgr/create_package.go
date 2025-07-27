package pkgmgr

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/mholt/archives"
)

var (
	ErrNameEmpty    = errors.New("имя пустое")
	ErrVerInvalid   = errors.New("неправильная версия")
	ErrTargetsEmpty = errors.New("нет файлов для упаковки")
)

func CreatePackage(a CreatePackageIn) error {
	a.Name = strings.TrimSpace(a.Name)
	if a.Name == "" {
		return ErrNameEmpty
	}

	_, err := semver.NewVersion(a.Ver)
	if err != nil {
		slog.Error(err.Error())
		return ErrVerInvalid
	}

	if len(a.Targets) == 0 {
		return ErrTargetsEmpty
	}

	// Call Scp method with file you want to upload to remote server.
	// Please make sure the `tmp` floder exists.
	err := ssh.Scp("/root/source.csv", "/tmp/target.csv")

	// Handle errors
	if err != nil {
		panic("Can't run remote command: " + err.Error())
	} else {
		fmt.Println("success")
	}

	return nil
}

func filenamesByTarget(target Target) ([]string, error){
	var res []string 
	filenames, err := filepath.Glob(target.Path)
	if err != nil {
		return nil, err
	}
	
	for _, name := range filenames {
		matched, err := filepath.Match(target.Exclude, name)
		if err != nil {
			return nil, err
		}
		if !matched {
			res = append(res, name)
		}
	}

	return res, nil
}

func createArchive(filenames []string) (string, error)  {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	
	var archiveFilenames map[string]string
	for _,name  := range filenames {
		fileAbs, err := filepath.Abs(name)
		if err != nil {
			return "", err
		}
		archiveFilenames[fileAbs], err = filepath.Rel(currentDir, fileAbs)
		if err != nil {
			return "", err
		}
	}

	ctx := context.TODO()

	// map files on disk to their paths in the archive using default settings (second arg)
	filenames, err := archives.FilesFromDisk(ctx, nil,archiveFilenames)
	if err != nil {
		return err
	}



	// create the output file we'll write to
	out, err := os.Create("/tmp/example.tar.gz")
	if err != nil {
		return err
	}
	defer out.Close()

	// we can use the CompressedArchive type to gzip a tarball
	// (since we're writing, we only set Archival, but if you're
	// going to read, set Extraction)
	format := archives.CompressedArchive{
		Compression: archives.Gz{},
		Archival:    archives.Tar{},
	}

	// create the archive
	err = format.Archive(ctx, out, filenames)
	if err != nil {
		return err
	}
}