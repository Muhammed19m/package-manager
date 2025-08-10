package pkgmgr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYamlToPackageInfo(t *testing.T) {
	t.Run("не валдиный yaml", func(t *testing.T) {
		yaml := `{"hi": 1}`
		pkgInfo, err := YamlToPackageInfo(yaml)
		assert.NoError(t, err)
		assert.Zero(t, pkgInfo)
	})

	t.Run("валидный yaml", func(t *testing.T) {
		yaml := `name: package
ver: "1.0"
targets: []`
		pkgInfo, err := YamlToPackageInfo(yaml)
		assert.NoError(t, err)
		assert.NotZero(t, pkgInfo)
	})
	t.Run("лишние поля будут отброшены", func(t *testing.T) {
		expectedPkgInfo := PackageInfo{
			Name:     "package",
			Ver:      "1.0",
			Targets:  []Target{},
			Packages: nil,
		}

		yaml := `name: package
ver: "1.0"
targets: []
test: 1`
		pkgInfo, err := YamlToPackageInfo(yaml)
		assert.NoError(t, err)
		assert.Equal(t, expectedPkgInfo, pkgInfo)
	})
	t.Run("незаданные поля равны нулю", func(t *testing.T) {
		pkgInfo, err := YamlToPackageInfo("some: q")
		assert.NoError(t, err)
		assert.Zero(t, pkgInfo)
	})
}
