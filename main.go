package main

import (
	"fmt"
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
			Value: "",
			Usage: "select domain name",
		},
	}
	app.Commands = commands
	if isConfigEmpty(*config) {
		fmt.Println("Verify your have the required parameters properly set.")
		fmt.Println("Review ~/.dnscli/config.json")
		os.Exit(1)
	}
	app.Run(os.Args)
}
