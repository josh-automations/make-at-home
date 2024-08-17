package gmailcmd

import (
	"github.com/urfave/cli/v2"
)

func (y *GmailCmd) getLabelsCommand() *cli.Command {
	return &cli.Command{
		Name:  "labels",
		Usage: "access labels resource",
		Subcommands: []*cli.Command{
			{
				Name:  "get",
				Usage: "get labels(s)",
				Action: func(ctx *cli.Context) error {
					args := ctx.Args()
					if args.Len() > 1 {
						return cli.Exit("too many arguments", 1)
					}

					args0 := args.First()
					if args0 == "" || args0 == "all" {
						return y.getLabels()
					} else {
						return y.getLabelById(args0)
					}
				},
			},
		},
	}
}

func (y *GmailCmd) getLabels() error {
	resp, err := y.service.Users.Labels.List("me").Do()
	if err != nil {
		return err
	}

	printJson(resp.Labels)

	return nil
}

func (y *GmailCmd) getLabelById(labelId string) error {
	resp, err := y.service.Users.Labels.Get("me", labelId).Do()
	if err != nil {
		return err
	}

	printJson(resp)

	return nil
}
