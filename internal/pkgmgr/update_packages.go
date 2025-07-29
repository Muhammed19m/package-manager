package pkgmgr

import "errors"

var (
	ErrPackagesEmpty = errors.New("список пакетов пустой")
)

func UpdatePackages(a UpdatePackagesIn) error {
	if len(a.Packages) == 0 {
		return ErrPackagesEmpty
	}
	
	return nil
}
