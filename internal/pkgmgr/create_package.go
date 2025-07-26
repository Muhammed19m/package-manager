package pkgmgr

import (
	"errors"
	"strings"
)

var (
	ErrNameEmpty = errors.New("имя пустое")
)

func CreatePackage(a CreatePackageIn) error {
	a.Name = strings.TrimSpace(a.Name)
	if a.Name == "" {
		return ErrNameEmpty
	}

	return nil

}
