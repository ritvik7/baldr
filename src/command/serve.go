package command

import (
	"flag"
	"fmt"
	"github.com/codegangsta/cli"
	"log"
	"net/http"
)

func Servec(port string, host string) {
	var addr = flag.String("addr", host+":"+port, "http service address")
	flag.Parse()
	http.Handle("/", http.FileServer(http.Dir("dist")))
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

var Serve = cli.Command{
	Name:    "serve",
	Aliases: []string{"s"},
	Usage:   "Serve your site locally.",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:   "o,host",
			Usage:  "Input address for serving the site",
			Value:  "0.0.0.0",
		},
		cli.StringFlag{
			Name:   "p,port",
			Usage:  "Input port for serving the site",
			Value:  "8080",
		},
	},
	Action: func(c *cli.Context) {
		fmt.Println("Starting Server...")
		Servec(c.String("port"),c.String("host"))
	},
}
