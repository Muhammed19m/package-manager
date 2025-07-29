package pkgmgr

import (
	"errors"
	"net"
	"os"
	"path/filepath"

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

	for _, pkg := range a.Packages {

		archiveName := pkg.Name + "-" + pkg.Ver + ".tar"
		remoteTargetAbs := filepath.Join(a.SshConfig.PackagesDir, archiveName)

		if err := scpClient.CopyFileFromRemote(remoteTargetAbs, a.DownloadDir, &scp.FileTransferOption{}); err != nil {
			return err
		}
	}

	return nil
}
