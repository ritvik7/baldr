package command

import (
	"github.com/codegangsta/cli"
)

var Pump = cli.Command{
	Name:    "pump",
	Aliases: []string{"p"},
	Usage:   "Serve site locally and watch for API changes.",
	Action: func(c *cli.Context) {
		println("starting server...")
	},
}
