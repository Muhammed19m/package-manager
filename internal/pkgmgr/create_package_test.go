package pkgmgr

import (
	"path"

	"github.com/appleboy/easyssh-proxy"
)

func (suite *testSuite) Test_CreatePackage() {
	// 1. создать тестконтейнер системы с запущенныи ssh
	// 2. создать SshConfig к этому контейнеру

	suite.Run("Name не должен быть пустым", func() {
		err := CreatePackage(CreatePackageIn{
			SshConfig: suite.sshConfig,
			Name:      "",
			Ver:       "1.0",
			Targets:   []Target{{}},
			Packages:  []PackageDependency{},
		})
		suite.ErrorIs(err, ErrNameEmpty)
	})

	suite.Run("Ver должен быть валидным", func() {
		err := CreatePackage(CreatePackageIn{
			SshConfig: suite.sshConfig,
			Name:      "somename",
			Ver:       "",
			Targets:   []Target{{}},
			Packages:  []PackageDependency{},
		})

		suite.ErrorIs(err, ErrVerInvalid)
	})

	suite.Run("Targets не должен быть пустым", func() {
		err := CreatePackage(CreatePackageIn{
			SshConfig: suite.sshConfig,
			Name:      "somename",
			Ver:       "1.0",
			Targets:   []Target{},
			Packages:  []PackageDependency{},
		})

		suite.ErrorIs(err, ErrTargetsEmpty)
	})

	suite.Run("успешная загрузка по шаблону", func() {
		err := CreatePackage(CreatePackageIn{
			SshConfig: suite.sshConfig,
			Name:      "package-1",
			Ver:       "1.0",
			Targets: []Target{
				{
					Path: "./funny*.png",
				},
			},
			Packages: []PackageDependency{},
		})

		suite.NoError(err)
		suite.assertRemoteFileExists("package-1-1.0.tar")
	})

	suite.Run("успешная загрузка по шаблону с исключением", func() {
		err := CreatePackage(CreatePackageIn{
			SshConfig: suite.sshConfig,
			Name:      "package-1",
			Ver:       "1.0",
			Targets: []Target{
				{
					Path:    "./funny*.png",
					Exclude: "*exluded.png",
				},
			},
			Packages: []PackageDependency{},
		})

		suite.NoError(err)
		suite.assertRemoteFileExists("package-1-1.0.tar")
	})
}

func (suite *testSuite) assertRemoteFileExists(f string) {
	// Create MakeConfig instance with remote username, server address and path to private key.
	ssh := &easyssh.MakeConfig{
		User:     suite.sshConfig.User,
		Server:   suite.sshConfig.Server,
		Password: suite.sshConfig.Passwd,
		Port:     suite.sshConfig.Port,
	}
	cmdCheckFile := "[ -e " + path.Join(suite.sshConfig.PackagesDir, f) + " ]"
	_, _, _, err := ssh.Run(cmdCheckFile)
	suite.NoError(err)
}
