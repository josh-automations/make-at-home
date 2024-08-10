// MIT License
//
// Copyright (c) 2024 josh-automations
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package ynabcmd

import (
	"context"

	"github.com/josh-automations/make-at-home/pkg/ynab"
	"github.com/urfave/cli/v2"
)

func (y *YnabCmd) getBudgetsCommand() *cli.Command {
	return &cli.Command{
		Name:  "budgets",
		Usage: "access budget resource",
		Subcommands: []*cli.Command{
			{
				Name:  "get",
				Usage: "get budget(s)",
				Action: func(ctx *cli.Context) error {
					args := ctx.Args()
					if args.Len() > 1 {
						return cli.Exit("too many arguments", 1)
					}

					args0 := args.First()
					if args0 == "" || args0 == "all" {
						return y.getBudgets(ctx)
					} else {
						return y.getBudgetById(ctx, args0)
					}
				},
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "include-accounts",
						Aliases: []string{"a"},
						Value:   false,
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
}

func (y *YnabCmd) getBudgets(ctx *cli.Context) error {
	params := &ynab.GetBudgetsParams{}
	if ctx.Bool("include-accounts") {
		val := true
		params.IncludeAccounts = &val
	}

	resp, err := y.client.GetBudgetsWithResponse(context.Background(), params)
	if err != nil {
		return err
	}

	switch resp.StatusCode() {
	case 200:
		printJson(*resp.JSON200)
	case 404:
		printJson(*resp.JSON404)
	default:
		printJson(*resp.JSONDefault)
	}

	return nil
}

func (y *YnabCmd) getBudgetById(ctx *cli.Context, budgetId string) error {
	params := &ynab.GetBudgetByIdParams{}
	val := ctx.Int64("last-server-knowledge")
	if val > 0 {
		params.LastKnowledgeOfServer = &val
	}

	resp, err := y.client.GetBudgetByIdWithResponse(context.Background(), budgetId, params)
	if err != nil {
		return err
	}

	switch resp.StatusCode() {
	case 200:
		printJson(*resp.JSON200)
	case 404:
		printJson(*resp.JSON404)
	default:
		printJson(*resp.JSONDefault)
	}

	return nil
}
