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

var containerSshConfig = pkgmgr.SshConfig{
	User:        "testuser",
	Server:      "localhost",
	Port:        "2222",
	Passwd:      "testpass",
	PackagesDir: "/tmp/pkgs",
}

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

					if err := pkgmgr.CreatePackage(containerSshConfig, packageInfo); err != nil {
						return fmt.Errorf("создание пакета: %w", err)
					}

					return nil
				},
			},
			{
				Name:  "update",
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

					if err := pkgmgr.UpdatePackages(containerSshConfig, ".", packagesInfo.Packages); err != nil {
						return fmt.Errorf("создание пакета: %w", err)
					}

					return nil
				},
			},
		},
		Name:  "pm",
		Usage: "пакетный менеджер для ....",
		Flags: []cli.Flag{},
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

// возвращает тип файла
func packageInfoParserByFilePath(path string) (func(y string) (pkgmgr.PackageInfo, error), error) {
	if strings.HasSuffix(path, ".json") {
		return pkgmgr.JsonToPackageInfo, nil
	}
	if strings.HasSuffix(path, ".yaml") || strings.HasSuffix(path, ".yml") {
		return pkgmgr.YamlToPackageInfo, nil
	}

	return nil, ErrUnknownFiletype
}

// возвращает функцию парсера
func packagesInfoParserByFilePath(path string) (func(y string) (pkgmgr.PackagesInfo, error), error) {
	if strings.HasSuffix(path, ".json") {
		return pkgmgr.JsonToPackagesInfo, nil
	}
	if strings.HasSuffix(path, ".yaml") || strings.HasSuffix(path, ".yml") {
		// return pkgmgr.YamlToPackagesInfo, nil
	}

	return nil, ErrUnknownFiletype
}
