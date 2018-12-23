package main

import (
	"os"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "dereutils"
	app.Usage = "The interface of im@sparql"
	app.Version = "0.0.1"

	app.Run(os.Args)
}