package pkgmgr

//  Создание архивов

// PackageForCreate представляет собой мета информацию о пакете для создания архива и отправки его на сервер
type PackageForCreate struct {
	Name     string
	Ver      string
	Targets  []Target
	Packages []PackageDependency
}

// Target это шаблон файлов для создания пакетов
type Target struct {
	Path    string
	Exclude string // Опциональное поле. Содержит шаблон исключения
}

// PackageDependency представляет с собой информацию о зависимости пакета
type PackageDependency struct {
	Name string
	Ver  string
}

//  Загрузка архивов

// PackagesForUpdate представляет собой запрос пакетов к скачиванию
type PackagesForUpdate struct {
	Packages []PackageRequest
}

// PackageRequest пакет
type PackageRequest struct {
	Name string
	Ver  string
}
