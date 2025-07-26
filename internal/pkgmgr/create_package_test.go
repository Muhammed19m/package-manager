package pkgmgr

func (suite *testSuite) Test_CreatePackage() {
	// 1. создать тестконтейнер системы с запущенныи ssh
	// 2. создать SshConfig к этому контейнеру
	
	suite.Run("Name не должен быть пустым", func() {
		err := CreatePackage(CreatePackageIn{
			SshConfig: suite.sshConfig,
			Name:      "",
			Ver:       "",
			Targets:   []Target{},
			Packages:  []PackageDependency{},
		})
		suite.Error(err)
	})

	suite.Run("Ver должен быть валидным", func() {
		err := CreatePackage(CreatePackageIn{
			SshConfig: suite.sshConfig,
			Name:      "",
			Ver:       "",
			Targets:   []Target{},
			Packages:  []PackageDependency{},
		})

		suite.Error(err)
	})

	suite.Run("Targets не должен быть пустым", func() {
		err := CreatePackage(CreatePackageIn{
			SshConfig: suite.sshConfig,
			Name:      "",
			Ver:       "",
			Targets:   []Target{},
			Packages:  []PackageDependency{},
		})

		suite.Error(err)
	})

	suite.Run("успешная загрузка по шаблону", func() {
		err := CreatePackage(CreatePackageIn{
			SshConfig: suite.sshConfig,
			Name:      "package-1",
			Ver:       "1.0",
			Targets: []Target{
				{
					Path: "./funny.png",
				},
			},
			Packages: []PackageDependency{},
		})

		suite.NoError(err)
	})

	suite.Run("успешная загрузка по шаблону с исключением", func() {
		err := CreatePackage(CreatePackageIn{
			SshConfig: suite.sshConfig,
			Name:      "package-1",
			Ver:       "1.0",
			Targets: []Target{
				Target{
					Path:    "./funny*.png",
					Exclude: "*exluded.png",
				},
			},
			Packages: []PackageDependency{},
		})

		suite.NoError(err)
	})
}
