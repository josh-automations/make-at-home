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
	"github.com/josh-automations/make-at-home/pkg/ynab"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

type YnabCmd struct {
	apiToken string
	client   *ynab.ClientWithResponses
	command  *cli.Command
}

func NewYnabCmd() *YnabCmd {
	ret := &YnabCmd{}
	ret.initCommand()
	return ret
}

func (y *YnabCmd) GetCommand() *cli.Command {
	return y.command
}

func (y *YnabCmd) getCmdFlags() []cli.Flag {
	return []cli.Flag{
		&cli.PathFlag{
			Name:    "config",
			Usage:   "path to config file",
			Value:   "/etc/ynab/config.yaml",
			Aliases: []string{"c"},
		},
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "ynab_api_token",
			Destination: &y.apiToken,
		}),
	}
}

func (y *YnabCmd) initCommand() {
	flags := y.getCmdFlags()
	y.command = &cli.Command{
		Name:  "ynab",
		Usage: "access ynab rest api client",
		Subcommands: []*cli.Command{
			y.getBudgetsCommand(),
			y.getAccountsCommand(),
			y.getCategoriesCommand(),
		},
		Flags: flags,
		Before: func(ctx *cli.Context) error {
			if err := altsrc.InitInputSourceWithContext(flags,
				altsrc.NewYamlSourceFromFlagFunc("config"))(ctx); err != nil {
				return err
			}

			if err := y.initYnabClient(); err != nil {
				return err
			}

			return nil
		},
	}
}

func (y *YnabCmd) initYnabClient() error {
	var err error
	y.client, err = ynab.NewYnabClient(ynab.DefaultServer, y.apiToken)
	if err != nil {
		return err
	}

	return nil
}
