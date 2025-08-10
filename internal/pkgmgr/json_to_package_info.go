package pkgmgr

import "encoding/json"

func JsonToPackageInfo(j string) (PackageInfo, error) {
	var jsonPkgInfo jsonPkgInfo

	if err := json.Unmarshal([]byte(j), &jsonPkgInfo); err != nil {
		return PackageInfo{}, err
	}

	return PackageInfo{
		Name:     jsonPkgInfo.Name,
		Ver:      jsonPkgInfo.Ver,
		Targets:  jsonPkgInfo.Targets,
		Packages: jsonPkgInfo.Packages,
	}, nil
}

type jsonPkgInfo struct {
	Name     string              `json:"name"`
	Ver      string              `json:"ver"`
	Targets  []Target            `json:"targets"`
	Packages []PackageDependency `json:"packages"`
}
