package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "dnscli"
	app.Usage = "DNSimple records from the CLI"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "domain, d",
			Value: "example.com",
			Usage: "select domain name",
		},
	}
	app.Commands = commands
	app.Run(os.Args)
}
