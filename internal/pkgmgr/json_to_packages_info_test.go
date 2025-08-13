package pkgmgr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonToPackagesInfo(t *testing.T) {
	t.Run("не валдиный json", func(t *testing.T) {
		json := "hi: 1"
		pkgInfo, err := JsonToPackagesInfo(json)
		assert.Error(t, err)
		assert.Zero(t, pkgInfo)
	})

	t.Run("валидный json", func(t *testing.T) {
		json := `{
		"packages": [
			{"name": "packet-1", "ver": ">=1.10"}
			]
		}`
		pkgInfo, err := JsonToPackagesInfo(json)
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

		json := `{
		"packages": [
			{"name": "packet-1", "ver": ">=1.10", "test": 1}
		],
		"test": 1
		}`
		pkgInfo, err := JsonToPackagesInfo(json)
		assert.NoError(t, err)
		assert.Equal(t, expectedPkgsInfo, pkgInfo)
	})
	t.Run("незаданные поля равны нулю", func(t *testing.T) {
		pkgInfo, err := JsonToPackagesInfo("{}")
		assert.NoError(t, err)
		assert.Zero(t, pkgInfo)
	})
}
