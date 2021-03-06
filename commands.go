package main

import (
	"os"

	"github.com/codegangsta/cli"
)

var commands = []cli.Command{
	{
		Name:  "records",
		Usage: "manage records",
		Subcommands: []cli.Command{
			{
				Name:  "add",
				Usage: "create record",
				Action: func(c *cli.Context) {
					domain := c.GlobalString("domain")
					recordType := c.String("type")
					content := c.String("content")
					name := c.String("name")
					if content == "" || recordType == "" {
						cli.ShowSubcommandHelp(c)
						return
					}
					createRecordDomain(domain, name, content, recordType)
				},
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "type, t",
						Usage: "record type (i.e: A, NS, MX)",
					},
					cli.StringFlag{
						Name:  "content, c",
						Usage: "record content",
					},
					cli.StringFlag{
						Name:  "name, n",
						Usage: "record name",
					},
				},
			},
			{
				Name:  "delete",
				Usage: "delete record",
				Action: func(c *cli.Context) {
					domain := c.GlobalString("domain")
					id := c.String("id")
					if id == "" {
						cli.ShowSubcommandHelp(c)
						return
					}
					deleteRecordDomain(domain, id)
				},
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Usage: "id record",
					},
				},
			},
			{
				Name:  "get",
				Usage: "get record",
				Action: func(c *cli.Context) {
					domain := c.GlobalString("domain")
					id := c.String("id")
					if id == "" {
						cli.ShowSubcommandHelp(c)
						return
					}
					getRecordDomain(domain, id)
				},
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Usage: "id record",
					},
				},
			},
			{
				Name:  "list",
				Usage: "list records",
				Action: func(c *cli.Context) {
					domain := c.GlobalString("domain")
					listRecordsDomain(domain)
				},
			},
			{
				Name:  "update",
				Usage: "update record",
				Action: func(c *cli.Context) {
					domain := c.GlobalString("domain")
					id := c.String("id")
					content := c.String("content")
					if id == "" || content == "" {
						cli.ShowSubcommandHelp(c)
						return
					}
					updateRecordDomain(domain, content, id)
				},
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Usage: "id record",
					},
					cli.StringFlag{
						Name:  "content, c",
						Usage: "record value",
					},
				},
			},
		},
	},
	{
		Name:  "autorenewal",
		Usage: "enable auto renewal",
		Action: func(c *cli.Context) {
			enabled := c.Bool("e")
			disabled := c.Bool("d")
			if !enabled {
				if !disabled {
					cli.ShowSubcommandHelp(c)
					os.Exit(1)
				}
			}
			domain := c.GlobalString("domain")
			enableAutoRenewal(domain, enabled)
		},
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "e",
				Usage: "enable auto renewal",
			},
			cli.BoolFlag{
				Name:  "d",
				Usage: "disable auto renewal",
			},
		},
	},
}
