package main

import (
	"os"

	"github.com/mkideal/cli"
)

func main() {
	os.Exit(cli.Run(new(ArgT), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*ArgT)

		if argv.Tag == "" {
			argv.Tag = "DEV"
		}

		if _, err := os.Stat("./VERSION"); os.IsNotExist(err) {
			createFileWithContent("VERSION", "0.0.0")
		}

		data := getVersionFileContent()
		err := ParseFile(data)
		CheckFatalError(err, false, "Error: Bad format in VERSION file")

		Core(data, argv)

		return nil
	}))
}
