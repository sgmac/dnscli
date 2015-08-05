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
					listRecordsDomain(domain)
				},
			},
			{
				Name:  "update",
				Usage: "update record for a domain",
				Action: func(c *cli.Context) {
					domain := c.GlobalString("domain")
					id := c.String("id")
					content := c.String("content")
					updateRecordDomain(domain, content, id)
				},
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Usage: "id record",
					},
					cli.StringFlag{
						Name:  "content, c",
						Usage: "update record value",
					},
				},
			},
			{
				Name:  "get",
				Usage: "get record for a domain",
				Action: func(c *cli.Context) {
					domain := c.GlobalString("domain")
					id := c.String("id")
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
				Name:  "delete",
				Usage: "delete record for a domain",
				Action: func(c *cli.Context) {
					domain := c.GlobalString("domain")
					id := c.String("id")
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
				Name:  "add",
				Usage: "create record for a domain",
				Action: func(c *cli.Context) {
					domain := c.GlobalString("domain")
					recordType := c.String("type")
					content := c.String("content")
					name := c.String("name")
					createRecordDomain(domain, name, content, recordType)
				},
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "type, t",
						Usage: "record type (i.e: A, NS, MX)",
					},
					cli.StringFlag{
						Name:  "content, c",
						Usage: "define record content",
					},
					cli.StringFlag{
						Name:  "name, n",
						Usage: "record name",
					},
				},
			},
		},
	},
}
