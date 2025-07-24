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
				"./funny.png"
			}
			Packages: []PackageDependency{},
		})

		assert.NoError(t, err)
	})

	t.Run("успешная загрузка по строковому шаблону", func(t *testing.T) {
		err := CreatePackage(PackageForCreate{
			Name:     "package-1",
			Ver:      "1.0",
			Targets:  []Target{
				TargetExtended{
					Path: "",
					Exclude: ""
				}
			}
			Packages: []PackageDependency{},
		})

		assert.NoError(t, err)
	})
}

func CreatePackage() error{return nil}