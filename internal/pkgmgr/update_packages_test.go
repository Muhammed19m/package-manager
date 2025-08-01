package pkgmgr

import (
	"path/filepath"

	"github.com/appleboy/easyssh-proxy"
)

func (suite *testSuite) Test_UpdatePackages() {
	suite.Run("Packages не должно быть пустым", func() {
		err := UpdatePackages(UpdatePackagesIn{
			SshConfig: suite.sshConfig,
			Packages:  []PackageRequest{},
		})

		suite.Error(err)
	})

	suite.Run("Пакет успешно загружен и существет в директории", func() {
		expectedFile := filepath.Join("testtmp","package-1-1.0.tar")

		suite.copyTestFile(filepath.Join("testdata","package-1-1.0.tar"))
		err := UpdatePackages(UpdatePackagesIn{
			SshConfig:   suite.sshConfig,
			DownloadDir: "testtmp",
			Packages: []PackageRequest{{
				Name: "package-1",
				Ver:  "1.0",
			}},
		})

		suite.FileExists(expectedFile)
		suite.NoError(err)
	})

	suite.Run("Пакета не существует на сервере и он не был загружен", func() {
		expectedFile := filepath.Join("testtmp","package-1-1.0.tar")

		suite.copyTestFile(filepath.Join("testdata","package-1-1.0.tar"))

		err := UpdatePackages(UpdatePackagesIn{
			SshConfig:   suite.sshConfig,
			DownloadDir: "testtmp",
			Packages: []PackageRequest{{
				Name: "package-1",
				Ver:  "1.0",
			}},
		})

		suite.FileExists(expectedFile)
		suite.NoError(err)
	})
}

func (suite *testSuite) copyTestFile(file string) {
	ssh := &easyssh.MakeConfig{
		User:     suite.sshConfig.User,
		Server:   suite.sshConfig.Server,
		Password: suite.sshConfig.Passwd,
		Port:     suite.sshConfig.Port,
	}

	filename := filepath.Base(file)
	remoteTargetAbs := filepath.Join(suite.sshConfig.PackagesDir, filename)

	err := ssh.Scp(file, remoteTargetAbs)
	suite.Require().NoError(err)
}
