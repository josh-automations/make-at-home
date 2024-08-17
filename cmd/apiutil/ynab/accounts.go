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

	"github.com/google/uuid"
	"github.com/josh-automations/make-at-home/internal/ynab"
	"github.com/urfave/cli/v2"
)

func (y *YnabCmd) getAccountsCommand() *cli.Command {
	return &cli.Command{
		Name:  "accounts",
		Usage: "access accounts resource",
		Subcommands: []*cli.Command{
			{
				Name:  "get",
				Usage: "get account(s)",
				Action: func(ctx *cli.Context) error {
					args := ctx.Args()
					if args.Len() > 1 {
						return cli.Exit("too many arguments", 1)
					}

					args0 := args.First()
					budgetId := ctx.String("budget")
					if args0 == "" || args0 == "all" {
						lastSvrKnowledge := ctx.Int64("last-server-knowledge")
						return y.getAccounts(budgetId, lastSvrKnowledge)
					} else {
						return y.getAccountById(budgetId, args0)
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
}

func (y *YnabCmd) getAccounts(budgetId string, lastSvrKnowledge int64) error {
	params := &ynab.GetAccountsParams{}
	if lastSvrKnowledge > 0 {
		params.LastKnowledgeOfServer = &lastSvrKnowledge
	}

	resp, err := y.client.GetAccountsWithResponse(context.Background(), budgetId, params)
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

func (y *YnabCmd) getAccountById(budgetId string, acctId string) error {
	acctUuid, err := uuid.Parse(acctId)
	if err != nil {
		return err
	}

	resp, err := y.client.GetAccountByIdWithResponse(context.Background(), budgetId, acctUuid)
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
