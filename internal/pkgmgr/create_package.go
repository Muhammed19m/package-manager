package pkgmgr

import (
	"errors"
	"log/slog"
	"strings"

	"github.com/Masterminds/semver"
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

	// // Call Scp method with file you want to upload to remote server.
	// // Please make sure the `tmp` floder exists.
	// err := ssh.Scp("/root/source.csv", "/tmp/target.csv")

	// // Handle errors
	// if err != nil {
	// 	panic("Can't run remote command: " + err.Error())
	// } else {
	// 	fmt.Println("success")
	// }

	return nil
}
