package pkgmgr

import "encoding/json"

func JsonToPackagesInfo(j string) (PackagesInfo, error) {
	var jsonPkgsInfo jsonPkgsInfo

	if err := json.Unmarshal([]byte(j), &jsonPkgsInfo); err != nil {
		return PackagesInfo{}, err
	}

	return PackagesInfo{
		Packages: jsonPkgsInfo.Packages,
	}, nil
}

type jsonPkgsInfo struct {
	Packages []PackageDependency `json:"packages"`
}
