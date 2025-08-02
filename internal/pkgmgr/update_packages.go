package pkgmgr

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"time"

	"github.com/mholt/archives"
	"github.com/povsister/scp"
)

var (
	ErrPackagesEmpty = errors.New("список пакетов пустой")
)

func UpdatePackages(a UpdatePackagesIn) error {
	if len(a.Packages) == 0 {
		return ErrPackagesEmpty
	}

	if err := os.MkdirAll(a.DownloadDir, 0755); err != nil {
		return err
	}

	clientConf := scp.NewSSHConfigFromPassword(a.SshConfig.User, a.SshConfig.Passwd)
	host := net.JoinHostPort(a.SshConfig.Server, a.SshConfig.Port)
	scpClient, err := scp.NewClient(host, clientConf, &scp.ClientOption{})
	if err != nil {
		return err
	}
	defer scpClient.Close()

	tmpDir := filepath.Join(os.TempDir(), "pkgmgr", fmt.Sprint(time.Now().Unix()))
	if err := os.MkdirAll(tmpDir, 0755); err != nil {
		return err
	}

	for _, pkg := range a.Packages {

		archiveName := pkg.Name + "-" + pkg.Ver + ".tar"
		remoteTargetAbs := path.Join(a.SshConfig.PackagesDir, archiveName)

		if err := scpClient.CopyFileFromRemote(remoteTargetAbs, tmpDir, &scp.FileTransferOption{}); err != nil {
			return err
		}

		input, err := os.Open(filepath.Join(tmpDir, archiveName))
		if err != nil {
			return err
		}
		defer input.Close()

		var format archives.Tar

		err = format.Extract(context.TODO(), input, func(ctx context.Context, f archives.FileInfo) error {
			archiveFile, err := f.Open()
			if err != nil {
				return err
			}
			defer archiveFile.Close()

			nameInArchive := f.NameInArchive
			if runtime.GOOS == "windows" {
				nameInArchive = filepath.FromSlash(nameInArchive)
			}

			newFile, err := os.Create(filepath.Join(a.DownloadDir, nameInArchive))
			if err != nil {
				return err
			}
			defer newFile.Close()

			if _, err = io.Copy(newFile, archiveFile); err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}
