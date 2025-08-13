package pkgmgr

//  Создание архивов

// PackageInfo представляет собой мета информацию о пакете для создания архива и отправки его на сервер
type PackageInfo struct {
	Name     string
	Ver      string
	Targets  []Target
	Packages []PackageDependency
}

// PackagesInfo представляет собой мета информацию о списке пакет для скачивания с сервера и распаковки
type PackagesInfo struct {
	Packages []PackageRequest
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

// PackageRequest пакет
type PackageRequest struct {
	Name string
	Ver  string
}

// ssh username@your_server_ip
// ssh Pebble
// .ssh/config
type SshConfig struct {
	Server       string
	Port         string
	User         string
	PackagesDir  string
	Passwd       string
	IdentityFile string
}
