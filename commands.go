package main

import "github.com/codegangsta/cli"

var commands = []cli.Command{
	{
		Name:  "records",
		Usage: "manage records",
		Subcommands: []cli.Command{
			{
				Name:  "list",
				Usage: "list records for a domain",
				Action: func(c *cli.Context) {
					domain := c.GlobalString("domain")
					if domain != "" {
						// list records domain
					}
				},
			},
		},
	},
}
