package pkgmgr

//  Создание архивов

// CreatePackageIn представляет собой мета информацию о пакете для создания архива и отправки его на сервер
type CreatePackageIn struct {
	SshConfig SshConfig
	Name      string
	Ver       string
	Targets   []Target
	Packages  []PackageDependency
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

// UpdatePackagesIn представляет собой запрос пакетов к скачиванию
type UpdatePackagesIn struct {
	SshConfig SshConfig
	Packages  []PackageRequest
}

// PackageRequest пакет
type PackageRequest struct {
	Name string
	Ver  string
}

// ssh username@your_server_ip
// ssh Pebble
// .ssh/config
type SshConfig struct {
	Host         string
	User         string
	PackagesDir  string
	Passwd       string
	IdentityFile string
}
