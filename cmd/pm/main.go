package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli/v3"

	"github.com/Muhammed19m/package-manager/internal/pkgmgr"
)

var (
	ErrEmptyPath = errors.New("путь к файлу пустой")
)

func main() {
	cmd := &cli.Command{
		Commands: []*cli.Command{
			{
				Name:  "create",
				Usage: "упаковывать файлы в архив и залить их на сервер по SSH",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					// Получить путь к файлу
					path := cmd.Args().Get(0)
					if path == "" {
						return ErrEmptyPath
					}
					//  Подобрать нужную функцию преобразования содержимого файла в PackageInfo
					contentToPackageInfo, err := packageInfoParserByFilePath(path)
					if err != nil {
						return err
					}
					// Прочитать файл
					content, err := os.ReadFile(path)
					if err != nil {
						return err
					}

					// Преобразовать содержимое файла PackageInfo
					packageInfo, err := contentToPackageInfo(string(content))
					if err != nil {
						return err
					}
					sshConfig := ComplateSshConfigFromCommand(cmd)

					if err := pkgmgr.CreatePackage(sshConfig, packageInfo); err != nil {
						return fmt.Errorf("создание пакета: %w", err)
					}

					return nil
				},
			},
			{
				Name: "update",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:      "DownloadDir",
						Value:     ".",
						Usage:     "Директория куда будет скачиваться файлы",
						Aliases:   []string{"dir"},
						TakesFile: true,
					},
				},
				Usage: "скачать файлы архивов по SSH и распаковать",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					// pkgmgr.UpdatePackages(shhConfig, downloadDir, packages)
					// Получить путь к файлу
					path := cmd.Args().Get(0)
					if path == "" {
						return ErrEmptyPath
					}
					//  Подобрать нужную функцию преобразования содержимого файла в PackageInfo
					contentToPackageInfo, err := packagesInfoParserByFilePath(path)
					if err != nil {
						return err
					}
					// Прочитать файл
					content, err := os.ReadFile(path)
					if err != nil {
						return err
					}

					// Преобразовать содержимое файла PackageInfo
					packagesInfo, err := contentToPackageInfo(string(content))
					if err != nil {
						return err
					}
					// Достать значение из флага DownloadDir(dir)
					DownloadDir := cmd.String("DownloadDir")

					sshConfig := ComplateSshConfigFromCommand(cmd)

					if err := pkgmgr.UpdatePackages(sshConfig, DownloadDir, packagesInfo.Packages); err != nil {
						return fmt.Errorf("загрузка пакета: %w", err)
					}

					return nil
				},
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "User",
				Usage:    "пользователь подключения к ssh",
				Required: true,
				Sources:  cli.EnvVars("USER_SSH_PM"),
			},
			&cli.StringFlag{
				Name:     "Hostname",
				Usage:    "сервер подключения к ssh",
				Required: true,
				Value:    "localhost",
				Sources:  cli.EnvVars("SERVER_SSH_PM"),
			},
			&cli.StringFlag{
				Name:    "Port",
				Usage:   "порт подключения к ssh",
				Value:   "22",
				Sources: cli.EnvVars("PORT_SSH_PM"),
			},
			&cli.StringFlag{
				Name:    "Passwrd",
				Usage:   "пароль подключения к ssh",
				Sources: cli.EnvVars("PASSWORD_SSH_PM"),
			},
			&cli.StringFlag{
				Name:    "PackagesDir",
				Usage:   "директория для работы спакетами на сервере",
				Value:   "/tmp/pkgs",
				Sources: cli.EnvVars("PACKAGES_DIR_SSH_PM"),
			},
			&cli.StringFlag{
				Name:    "IdentityFile",
				Usage:   "идентификационный файл",
				Value:   "",
				Sources: cli.EnvVars("IDENTITY_FILE_SSH_PM"),
			},
		},
		Name:  "pm",
		Usage: "пакетный менеджер для ....",
		Action: func(_ context.Context, cmd *cli.Command) error {
			cli.ShowAppHelpAndExit(cmd, 2)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

var ErrUnknownFiletype = errors.New("неизвестный тип файла")

// возвращает функцию парсера для парсинга файла упаковки файла
func packageInfoParserByFilePath(path string) (func(y string) (pkgmgr.PackageInfo, error), error) {
	if strings.HasSuffix(path, ".json") {
		return pkgmgr.JsonToPackageInfo, nil
	}
	if strings.HasSuffix(path, ".yaml") || strings.HasSuffix(path, ".yml") {
		return pkgmgr.YamlToPackageInfo, nil
	}

	return nil, ErrUnknownFiletype
}

// возвращает функцию парсера для парсинга файла распаковки файлов
func packagesInfoParserByFilePath(path string) (func(y string) (pkgmgr.PackagesInfo, error), error) {
	if strings.HasSuffix(path, ".json") {
		return pkgmgr.JsonToPackagesInfo, nil
	}
	if strings.HasSuffix(path, ".yaml") || strings.HasSuffix(path, ".yml") {
		return pkgmgr.YamlToPackagesInfo, nil
	}

	return nil, ErrUnknownFiletype
}

// возвращает SshConfig собирая из флагов или EnvVars
func ComplateSshConfigFromCommand(cmd *cli.Command) pkgmgr.SshConfig {
	sshConfig := pkgmgr.SshConfig{
		Server:       cmd.String("Hostname"),
		Port:         cmd.String("Port"),
		User:         cmd.String("User"),
		PackagesDir:  cmd.String("PackagesDir"),
		Passwd:       cmd.String("Passwrd"),
		IdentityFile: cmd.String("IdentityFile"),
	}
	return sshConfig
}
