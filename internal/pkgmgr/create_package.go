package pkgmgr

import (
	"errors"
	"log/slog"
	"strings"

	"github.com/Masterminds/semver"
)

var (
	ErrNameEmpty = errors.New("имя пустое")
	ErrVerInvalid = errors.New("неправильная версия")
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

	return nil

}
