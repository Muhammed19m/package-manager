package pkgmgr

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/appleboy/easyssh-proxy"
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

	ssh := &easyssh.MakeConfig{
		User:     a.SshConfig.User,
		Server:   a.SshConfig.Server,
		Password: a.SshConfig.Passwd,
		Port:     a.SshConfig.Port,
	}

	for _, pkg := range a.Packages {

		archiveName := pkg.Name + "-" + pkg.Ver + ".tar"
		remoteTargetAbs := filepath.Join(a.SshConfig.PackagesDir, archiveName)

		if err := ssh.Scp(remoteTargetAbs, a.DownloadDir); err != nil {
			return err
		}

	}

	return nil
}
