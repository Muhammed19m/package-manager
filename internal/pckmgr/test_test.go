package pkgmgr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreatePackage(t *testing.T) {
	t.Run("Name должен быть валидным", func(t *testing.T) {
		err := CreatePackage(PackageForCreate{
		Name:     "",
		Ver:      "",
		Targets:  []Target{},
		Packages: []PackageDependency{},
		})

		assert.Error(t, err)
	})

	t.Run("Ver должен быть валидным", func(t *testing.T) {
		err := CreatePackage(PackageForCreate{
		Name:     "",
		Ver:      "",
		Targets:  []Target{},
		Packages: []PackageDependency{},
		})

		assert.Error(t, err)
	})
	
	t.Run("Targets не должен быть пустым", func(t *testing.T) {
		err := CreatePackage(PackageForCreate{
			Name:     "",
			Ver:      "",
			Targets:  []Target{},
			Packages: []PackageDependency{},
		})

		assert.Error(t, err)
	})

	t.Run("успешная загрузка по строковому шаблону", func(t *testing.T) {
		err := CreatePackage(PackageForCreate{
			Name:     "package-1",
			Ver:      "1.0",
			Targets:  []Target{
				"./funny.png",
			},
			Packages: []PackageDependency{},
		})

		assert.NoError(t, err)
	})

	t.Run("успешная загрузка по расширенному шаблону", func(t *testing.T) {
		err := CreatePackage(PackageForCreate{
			Name:     "package-1",
			Ver:      "1.0",
			Targets:  []Target{
				TargetExtended{
					Path: "./funny*.png",
					Exclude: "",
				},
			},
			Packages: []PackageDependency{},
		})

		assert.NoError(t, err)
	})

	t.Run("успешная загрузка вместе с зависимостью", func(t *testing.T) {
		err := CreatePackage(PackageForCreate{
			Name:     "package-1",
			Ver:      "1.0",
			Targets:  []Target{
				"./funny.png",
			},
			Packages: []PackageDependency{
				{Name: "package-3", Ver: "<=2.0"},
			},
		})

		assert.NoError(t, err)
	})

	t.Run("зависимость должна существовать", func(t *testing.T) {
		err := CreatePackage(PackageForCreate{
			Name:     "package-1",
			Ver:      "1.0",
			Targets:  []Target{
				"./funny.png",
			},
			Packages: []PackageDependency{
				{Name: "package-not-exits", Ver: "<=2.0"},
			},
		})

		assert.Error(t, err)
	})
}

func CreatePackage(a PackageForCreate) error{return nil}

func Test_UpdatePackage(t *testing.T) {
	t.Run("Name должен быть валидным", func(t *testing.T) {
		err := UpdatePackage(PackagesForUpdate{
		Name:     "",
		Ver:      "",
		Targets:  []Target{},
		Packages: []PackageDependency{},
		})

		assert.Error(t, err)
	})
}

func UpdatePackage(a PackagesForUpdate) error {return nil}