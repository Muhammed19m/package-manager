package pkgmgr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonToPackageInfo(t *testing.T) {
	t.Run("не валдиный json", func(t *testing.T) {
		json := "hi: 1"
		pkgInfo, err := jsonToPackageInfo(json)
		assert.Error(t, err)
		assert.Zero(t, pkgInfo)
	})

	t.Run("валидный json", func(t *testing.T) {
		json := `{
		"name": "package",
		"ver": "1.0",
		"targets": []
		}`
		pkgInfo, err := jsonToPackageInfo(json)
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

		json := `{
		"name": "package",
		"ver": "1.0",
		"targets": [],
		"test": 1
		}`
		pkgInfo, err := jsonToPackageInfo(json)
		assert.NoError(t, err)
		assert.Equal(t, expectedPkgInfo, pkgInfo)
	})
	t.Run("незаданные поля равны нулю", func(t *testing.T) {
		pkgInfo, err := jsonToPackageInfo("{}")
		assert.NoError(t, err)
		assert.Zero(t, pkgInfo)
	})
}
