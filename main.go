package main

import (
	"command"
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	var app = cli.NewApp()
	app.Usage = "Generates API documentation site from swagger spec."
	app.Commands = []cli.Command{
		command.Init,
		command.Build,
		command.Serve,
		command.Pump,
	}
	app.Run(os.Args)

}
