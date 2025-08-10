package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Muhammed19m/package-manager/internal/pkgmgr"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Commands: []*cli.Command{
			{
				Name:  "create",
				Usage: "упаковывать файлы в архив, и залить их на сервер по SSH",

				Action: func(ctx context.Context, cmd *cli.Command) error {
					path := cmd.Args().Get(0)
					fmt.Println("completed task: ", cmd.Args().First())

					contentToConfig, err := contentParserByFilePath(path)
					if err != nil {
						return err
					}

					contentFile, err := os.ReadFile(path)
					if err != nil {
						return err
					}
					packageInfo, err := contentToConfig(string(contentFile))

					pkgmgr.CreatePackage(shhConfig, packageInfo)

					return nil
				},
			},
			{
				Name:  "update",
				Usage: "скачать файлы архивов по SSH и распаковать",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					// pm update ./packages.json
					path := cmd.Args().Get(0)
					fmt.Println("completed task: ", cmd.Args().First())

					contentToConfig, err := contentParserByFilePath(path)
					if err != nil {
						return err
					}

					contentFile, err := os.ReadFile(path)
					if err != nil {
						return err
					}
					packageInfo, err := contentToConfig(string(contentFile))

					packages := []pkgmgr.PackageRequest{}

					pkgmgr.CreatePackage(shhConfig, packageInfo)

					// cfg := ...
					pkgmgr.UpdatePackages(shhConfig, downloadDir, packages)
					return nil
				},
			},
		},
		Name:  "pm",
		Usage: "пакетный менеджер для ....",
		Flags: []cli.Flag{},
		Action: func(context.Context, *cli.Command) error {

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

var ErrUnknownFiletype = errors.New("неизвестный тип файла")

// возвращает тип файла
func contentParserByFilePath(path string) (func(y string) (pkgmgr.PackageInfo, error), error) {
	if strings.HasSuffix(path, ".json") {
		return pkgmgr.JsonToPackageInfo, nil
	}
	if strings.HasSuffix(path, ".yaml") || strings.HasSuffix(path, ".yml") {
		return pkgmgr.YamlToPackageInfo, nil
	}

	return nil, ErrUnknownFiletype
}
