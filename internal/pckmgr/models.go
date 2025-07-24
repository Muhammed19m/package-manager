package pckmgr


type PackageMeta struct {
	Name string
	Ver string
	Targets []Target
 	Packages []PackageDependency
}

type Target any // Может быть строкой или типом TargetWithExclude

type TargetWithExclude struct {
	Path string
	Exclude string
}

type PackageDependency struct {
	Name string
	Ver string
}