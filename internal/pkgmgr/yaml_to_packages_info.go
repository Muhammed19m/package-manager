package pkgmgr

import "gopkg.in/yaml.v3"

func YamlToPackagesInfo(j string) (PackagesInfo, error) {
	var yamlPkgsInfo yamlPkgsInfo

	if err := yaml.Unmarshal([]byte(j), &yamlPkgsInfo); err != nil {
		return PackagesInfo{}, err
	}

	return PackagesInfo{
		Packages: yamlPkgsInfo.Packages,
	}, nil
}

type yamlPkgsInfo struct {
	Packages []PackageRequest `yaml:"packages"`
}
