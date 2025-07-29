package pkgmgr

import (
	"path/filepath"

	"github.com/appleboy/easyssh-proxy"
)

func (suite *testSuite) Test_UpdatePackages() {
	sshCfg := SshConfig{
		User:         "",
		Passwd:       "",
		IdentityFile: "",
	}

	suite.Run("Packages не должно быть пустым", func() {
		err := UpdatePackages(UpdatePackagesIn{
			SshConfig: sshCfg,
			Packages:  []PackageRequest{},
		})

		suite.Error(err)
	})

	suite.Run("Пакет успешно загружен и существет в директории", func() {
		expectedFile := "./funny1.png"
		err := UpdatePackages(UpdatePackagesIn{
			SshConfig: sshCfg,
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
		err := UpdatePackages(UpdatePackagesIn{
			SshConfig: sshCfg,
			Packages: []PackageRequest{{
				Name: "package-1",
				Ver:  "1.0",
			}},
		})

		suite.FileExists(expectedFile)
		suite.Error(err)
	})
}

func (suite *testSuite) copyTestFile(localArchive string, sshConfig SshConfig) {
	ssh := &easyssh.MakeConfig{
		User:     sshConfig.User,
		Server:   sshConfig.Server,
		Password: sshConfig.Passwd,
		Port:     sshConfig.Port,
	}

	remoteTargetAbs := filepath.Join(sshConfig.PackagesDir, localArchive)
	err := ssh.Scp(localArchive, remoteTargetAbs)
	suite.Require().NoError(err)
}
