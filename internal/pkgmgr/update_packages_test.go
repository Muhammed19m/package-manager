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
		expectedFile := "./testtmp/package-1-1.0.tar"
		
		suite.copyTestFile("./testdata/package-1-1.0.tar")

		err := UpdatePackages(UpdatePackagesIn{
			SshConfig: suite.sshConfig,
			DownloadDir: "./testtmp",
			Packages: []PackageRequest{{
				Name: "package-1",
				Ver:  "1.0",
			}},
		})

		suite.FileExists(expectedFile)
		suite.Error(err)
	})

	suite.Run("Пакета не существует на сервере и он не был загружен", func() {
		expectedFile := "./funny1.png"
		
		suite.copyTestFile("./testdata/package-1-1.0.tar")

		err := UpdatePackages(UpdatePackagesIn{
			SshConfig: suite.sshConfig,
			Packages: []PackageRequest{{
				Name: "package-1",
				Ver:  "1.0",
			}},
		})

		suite.FileExists(expectedFile)
		suite.Error(err)
	})
}

func (suite *testSuite) copyTestFile(localArchive string) {
	ssh := &easyssh.MakeConfig{
		User:     suite.sshConfig.User,
		Server:   suite.sshConfig.Server,
		Password: suite.sshConfig.Passwd,
		Port:     suite.sshConfig.Port,
	}

	remoteTargetAbs := filepath.Join(suite.sshConfig.PackagesDir, localArchive)
	err := ssh.Scp(localArchive, remoteTargetAbs)
	suite.Require().NoError(err)
}
