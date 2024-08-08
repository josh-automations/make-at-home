package main

import (
	"context"

	"github.com/josh-automations/make-at-home/pkg/ynab"
	"github.com/urfave/cli/v2"
)

var ynabCategoriesCommand *cli.Command = &cli.Command{
	Name:  "categories",
	Usage: "access categories resource",
	Subcommands: []*cli.Command{
		{
			Name:  "get",
			Usage: "get category/ies",
			Action: func(ctx *cli.Context) error {
				args := ctx.Args()
				if args.Len() > 1 {
					return cli.Exit("too many arguments", 1)
				}

				args0 := args.First()
				if args0 == "" || args0 == "all" {
					return ynabGetCategories(ctx)
				} else {
					return ynabGetCategoryById(ctx, args0)
				}
			},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "budget",
					Aliases:  []string{"b"},
					Required: true,
				},
				&cli.Int64Flag{
					Name:    "last-server-knowledge",
					Aliases: []string{"k"},
					Value:   -1,
				},
			},
		},
	},
}

func ynabGetCategories(ctx *cli.Context) error {
	var client *ynab.ClientWithResponses
	client, err := ynab.NewYnabClient(ynab.DefaultServer, ynabApiToken)
	if err != nil {
		return err
	}

	params := &ynab.GetCategoriesParams{}
	val := ctx.Int64("last-server-knowledge")
	if val > 0 {
		params.LastKnowledgeOfServer = &val
	}

	resp, err := client.GetCategoriesWithResponse(context.TODO(), ctx.String("budget"), params)
	if err != nil {
		return err
	}

	switch resp.StatusCode() {
	case 200:
		return printJson(*resp.JSON200)
	case 404:
		return printJson(*resp.JSON404)
	default:
		return printJson(*resp.JSONDefault)
	}
}

func ynabGetCategoryById(ctx *cli.Context, catId string) error {
	var client *ynab.ClientWithResponses
	client, err := ynab.NewYnabClient(ynab.DefaultServer, ynabApiToken)
	if err != nil {
		return err
	}

	resp, err := client.GetCategoryByIdWithResponse(context.TODO(), ctx.String("budget"), catId)
	if err != nil {
		return err
	}

	switch resp.StatusCode() {
	case 200:
		return printJson(*resp.JSON200)
	case 404:
		return printJson(*resp.JSON404)
	default:
		return printJson(*resp.JSONDefault)
	}
}
