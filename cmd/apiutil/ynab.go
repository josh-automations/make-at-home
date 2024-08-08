package main

import (
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

var ynabApiToken string

var ynabFlags []cli.Flag = []cli.Flag{
	&cli.PathFlag{
		Name:    "config",
		Usage:   "path to config file",
		Value:   "/etc/ynab/config.yaml",
		Aliases: []string{"c"},
	},
	altsrc.NewStringFlag(&cli.StringFlag{
		Name:        "ynab_api_token",
		Destination: &ynabApiToken,
	}),
}

var ynabCommand *cli.Command = &cli.Command{
	Name:  "ynab",
	Usage: "access ynab rest api client",
	Subcommands: []*cli.Command{
		ynabBudgetsCommand,
		ynabAccountsCommand,
		ynabCategoriesCommand,
	},
	Flags: ynabFlags,
	Before: altsrc.InitInputSourceWithContext(ynabFlags,
		altsrc.NewYamlSourceFromFlagFunc("config")),
}
