package pkgmgr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreatePackage(t *testing.T) {
	// 1. создать тестконтейнер системы с запущенныи ssh
	// 2. создать SshConfig к этому контейнеру
	sshCfg := SshConfig{
		Host:         "",
		User:         "",
		PackagesDir:  "",
		Passwd:       "",
		IdentityFile: "",
	}

	t.Run("Name должен быть валидным", func(t *testing.T) {
		err := CreatePackage(CreatePackageIn{
			SshConfig: sshCfg,
			Name:      "",
			Ver:       "",
			Targets:   []Target{},
			Packages:  []PackageDependency{},
		})

		assert.Error(t, err)
	})

	t.Run("Ver должен быть валидным", func(t *testing.T) {
		err := CreatePackage(CreatePackageIn{
			SshConfig: sshCfg,
			Name:      "",
			Ver:       "",
			Targets:   []Target{},
			Packages:  []PackageDependency{},
		})

		assert.Error(t, err)
	})

	t.Run("Targets не должен быть пустым", func(t *testing.T) {
		err := CreatePackage(CreatePackageIn{
			SshConfig: sshCfg,
			Name:      "",
			Ver:       "",
			Targets:   []Target{},
			Packages:  []PackageDependency{},
		})

		assert.Error(t, err)
	})

	t.Run("успешная загрузка по шаблону", func(t *testing.T) {
		err := CreatePackage(CreatePackageIn{
			SshConfig: sshCfg,
			Name:      "package-1",
			Ver:       "1.0",
			Targets: []Target{
				{
					Path: "./funny.png",
				},
			},
			Packages: []PackageDependency{},
		})

		assert.NoError(t, err)
	})

	t.Run("успешная загрузка по шаблону с исключением", func(t *testing.T) {
		err := CreatePackage(CreatePackageIn{
			SshConfig: sshCfg,
			Name:      "package-1",
			Ver:       "1.0",
			Targets: []Target{
				Target{
					Path:    "./funny*.png",
					Exclude: "*exluded.png",
				},
			},
			Packages: []PackageDependency{},
		})

		assert.NoError(t, err)
	})

	t.Run("успешная загрузка вместе с зависимостью", func(t *testing.T) {
		err := CreatePackage(CreatePackageIn{
			SshConfig: sshCfg,
			Name:      "package-1",
			Ver:       "1.0",
			Targets: []Target{
				Target{
					Path: "./funny*.png",
				},
			},
			Packages: []PackageDependency{
				{Name: "package-3", Ver: "<=2.0"},
			},
		})

		assert.NoError(t, err)
	})

	t.Run("зависимость должна существовать", func(t *testing.T) {
		err := CreatePackage(CreatePackageIn{
			SshConfig: sshCfg,
			Name:      "package-1",
			Ver:       "1.0",
			Targets: []Target{
				Target{
					Path: "./funny*.png",
				},
			},
			Packages: []PackageDependency{
				{Name: "package-not-exits", Ver: "<=2.0"},
			},
		})

		assert.Error(t, err)
	})
}
