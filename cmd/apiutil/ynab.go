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
