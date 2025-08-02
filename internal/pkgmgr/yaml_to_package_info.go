package pkgmgr

import (
	"gopkg.in/yaml.v3"
)

func yamlToPackageInfo(y string) (PackageInfo, error) {
	var yamlPkgInfo yamlPkgInfo

	if err := yaml.Unmarshal([]byte(y), &yamlPkgInfo); err != nil {
		return PackageInfo{}, err
	}

	return PackageInfo{
		Name:     yamlPkgInfo.Name,
		Ver:      yamlPkgInfo.Ver,
		Targets:  yamlPkgInfo.Targets,
		Packages: yamlPkgInfo.Packages,
	}, nil
}

type yamlPkgInfo struct {
	Name     string              `yaml:"name"`
	Ver      string              `yaml:"ver"`
	Targets  []Target            `yaml:"targets"`
	Packages []PackageDependency `yaml:"packages"`
}
