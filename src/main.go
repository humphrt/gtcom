package main

import (
	"os"
	"path/filepath"

	"github.com/mkideal/cli"
)

func main() {
	os.Exit(cli.Run(new(ArgT), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*ArgT)

		if argv.Tag == "" {
			argv.Tag = "DEV"
		}

		absPath, err := filepath.Abs("./")
		CheckFatalError(err, true)
		rootPath := GetRootPath(absPath, 0)
		if _, err := os.Stat(rootPath + "/VERSION"); os.IsNotExist(err) {
			createFileWithContent(rootPath+"/VERSION", "0.0.0")
		}

		data, err := GetVersionFileContent(rootPath + "/VERSION")
		CheckFatalError(err, true)
		err = ParseFile(data)
		CheckFatalError(err, false, "Error: Bad format in VERSION file")

		Core(data, argv, rootPath)
		return nil
	}))
}
