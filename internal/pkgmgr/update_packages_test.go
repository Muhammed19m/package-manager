package pkgmgr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UpdatePackages(t *testing.T) {
	sshCfg := SshConfig{
		User:         "",
		Passwd:       "",
		IdentityFile: "",
	}

	t.Run("Packages не должно быть пустым", func(t *testing.T) {
		err := UpdatePackages(UpdatePackagesIn{
			SshConfig: sshCfg,
			Packages:  []PackageRequest{},
		})

		assert.Error(t, err)
	})

	t.Run("Пакет успешно загружен и существет в директории", func(t *testing.T) {
		expectedFile := "./funny1.png"
		err := UpdatePackages(UpdatePackagesIn{
			SshConfig: sshCfg,
			Packages: []PackageRequest{{
				Name: "package-1",
				Ver:  "1.0",
			}},
		})

		assert.FileExists(t, expectedFile)
		assert.NoError(t, err)
	})

	t.Run("Пакета не существует на сервере и он не был загружен", func(t *testing.T) {
		expectedFile := "./funny1.png"
		err := UpdatePackages(UpdatePackagesIn{
			SshConfig: sshCfg,
			Packages: []PackageRequest{{
				Name: "package-1",
				Ver:  "1.0",
			}},
		})

		assert.NoFileExists(t, expectedFile)
		assert.Error(t, err)
	})
}
