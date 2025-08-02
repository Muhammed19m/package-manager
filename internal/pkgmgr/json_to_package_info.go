package pkgmgr

func jsonToPackageInfo(json string) (PackageInfo, error) {
	return PackageInfo{}, nil
}

type jsonPkgInfo struct {
	Name     string              `json:"name,omitempty"`
	Ver      string              `json:"ver,omitempty"`
	Targets  []Target            `json:"targets,omitempty"`
	Packages []PackageDependency `json:"packages,omitempty"`
}
