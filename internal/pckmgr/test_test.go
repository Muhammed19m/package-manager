package pkgmgr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreatePackage(t *testing.T) {
	sshCfg := SshConfig{
		Host:         "",
		User:         "",
		Passwd:       "",
		IdentityFile: "",
	}

	t.Run("Name должен быть валидным", func(t *testing.T) {
		err := CreatePackage(CreatePackageIn{
			SshConfig: sshCfg,
			Name:     "",
			Ver:      "",
			Targets:  []Target{},
			Packages: []PackageDependency{},
		})

		assert.Error(t, err)
	})

	t.Run("Ver должен быть валидным", func(t *testing.T) {
		err := CreatePackage(CreatePackageIn{
			SshConfig: sshCfg,
			Name:     "",
			Ver:      "",
			Targets:  []Target{},
			Packages: []PackageDependency{},
		})

		assert.Error(t, err)
	})

	t.Run("Targets не должен быть пустым", func(t *testing.T) {
		err := CreatePackage(CreatePackageIn{
			SshConfig: sshCfg,
			Name:     "",
			Ver:      "",
			Targets:  []Target{},
			Packages: []PackageDependency{},
		})

		assert.Error(t, err)
	})

	t.Run("успешная загрузка по шаблону", func(t *testing.T) {
		err := CreatePackage(CreatePackageIn{
			SshConfig: sshCfg,
			Name: "package-1",
			Ver:  "1.0",
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
			Name: "package-1",
			Ver:  "1.0",
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
			Name: "package-1",
			Ver:  "1.0",
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
			Name: "package-1",
			Ver:  "1.0",
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

func CreatePackage(a CreatePackageIn) error { return nil }

func Test_UpdatePackages(t *testing.T) {
	sshCfg := SshConfig{
		Host:         "",
		User:         "",
		Passwd:       "",
		IdentityFile: "",
	}

	t.Run("Packages не должно быть пустым", func(t *testing.T) {
		err := UpdatePackages(UpdatePackagesIn{
			SshConfig: sshCfg,
			Packages: []PackageRequest{},
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
}

func UpdatePackages(a UpdatePackagesIn) error { return nil }
