package main

import (
	"golang-cli/app"
	"os"
)

func main() {
	appCli := app.GeneratorCli()
	appCli.Run(os.Args)

}
