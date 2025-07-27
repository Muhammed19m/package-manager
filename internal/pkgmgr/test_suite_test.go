package pkgmgr

import (
	"context"
	"path"
	"testing"
	"time"

	"github.com/docker/go-connections/nat"
	testifySuite "github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type testSuite struct {
	testifySuite.Suite
	sshServerCloser func()
	sshConfig       SshConfig
	sshCleanup      func()
}

func Test_TestSuite(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	testifySuite.Run(t, new(testSuite))
}

func (suite *testSuite) newSshContainer() {
	// Конфигурация для подключения
	suite.sshConfig = SshConfig{
		Server:      "localhost",
		Port:        "",// late
		User:        "testuser",
		Passwd:      "testpass",
		PackagesDir: "/tmp/pkgs",
	}
	internalPort := nat.Port("2222/tcp")
	// Создать контейнер
	sshContainer, err := testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
		Started: true,
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "linuxserver/openssh-server:version-10.0_p1-r7",
			// ExposedPorts: []string{suite.sshConfig.Port + "/tcp"},
			ExposedPorts: []string{internalPort.Port()},
			Env: map[string]string{
				"USER_NAME":       suite.sshConfig.User,
				"USER_PASSWORD":   suite.sshConfig.Passwd,
				"PASSWORD_ACCESS": "true",
			},
			WaitingFor: wait.ForListeningPort(internalPort).WithStartupTimeout(5 * time.Second),
		},
	})
	suite.Require().NoError(err)
	suite.sshServerCloser = func() {
		suite.Require().NotNil(sshContainer)
		_ = sshContainer.Terminate(context.Background())
	}

	// Связать внешний порт с внутренним
	suitesshConfigPort, err := sshContainer.MappedPort(context.Background(), internalPort)
	suite.Require().NoError(err)
	suite.sshConfig.Port = suitesshConfigPort.Port()

	// Создать директорию для пакетов
	ctxMkdir, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	_, _, err = sshContainer.Exec(ctxMkdir, []string{"mkdir", "-p", suite.sshConfig.PackagesDir})
	suite.Require().NoError(err)

	suite.sshCleanup = func() {
		// Удалить все из директории пакетов
		ctxCleanup, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		_, _, err := sshContainer.Exec(ctxCleanup, []string{"rm", "-r", path.Join(suite.sshConfig.PackagesDir, "/*")})
		suite.Require().NoError(err)
	}
}

// SetupTest выполняется перед каждым тестом, связанным с suite
func (suite *testSuite) SetupTest() {
	suite.newSshContainer()
}

// TearDownSubTest выполняется после каждого подтеста, связанного с suite
func (suite *testSuite) TearDownSubTest() {
	suite.sshCleanup()
}

// TearDownSubTest выполняется после каждого подтеста, связанного с suite
func (suite *testSuite) TearDownTest() {
	suite.sshServerCloser()
}
