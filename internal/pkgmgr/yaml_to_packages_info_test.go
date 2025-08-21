package pkgmgr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYamlToPackagesInfo(t *testing.T) {
	t.Run("не валдиный yaml", func(t *testing.T) {
		yaml := `{"hi": 1}`
		pkgInfo, err := YamlToPackageInfo(yaml)
		assert.NoError(t, err)
		assert.Zero(t, pkgInfo)
	})

	t.Run("валидный yaml", func(t *testing.T) {
		yaml := `packages: [
{"name": "packet-1", "ver": ">=1.10"}
]`
		pkgInfo, err := YamlToPackagesInfo(yaml)
		assert.NoError(t, err)
		assert.NotZero(t, pkgInfo)
	})
	t.Run("лишние поля будут отброшены", func(t *testing.T) {
		expectedPkgsInfo := PackagesInfo{
			Packages: []PackageRequest{
				{
					Name: "packet-1",
					Ver:  ">=1.10",
				},
			},
		}

		yaml := `packages: [{"name": "packet-1", "ver": ">=1.10", "test": 1}]
test: 1`
		pkgInfo, err := YamlToPackagesInfo(yaml)
		assert.NoError(t, err)
		assert.Equal(t, expectedPkgsInfo, pkgInfo)
	})
	t.Run("незаданные поля равны нулю", func(t *testing.T) {
		pkgInfo, err := YamlToPackagesInfo(`test1: 1
test2: 2`)
		assert.NoError(t, err)
		assert.Zero(t, pkgInfo)
	})
}
