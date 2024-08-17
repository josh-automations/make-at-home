package gmailcmd

import (
	"github.com/urfave/cli/v2"
)

func (y *GmailCmd) getEmailsCommand() *cli.Command {
	return &cli.Command{
		Name:  "emails",
		Usage: "access emails resource",
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:    "label",
				Aliases: []string{"s"},
			},
		},
		Subcommands: []*cli.Command{
			{
				Name:  "get",
				Usage: "get email(s)",
				Action: func(ctx *cli.Context) error {
					args := ctx.Args()
					if args.Len() > 1 {
						return cli.Exit("too many arguments", 1)
					}

					args0 := args.First()
					if args0 == "" || args0 == "all" {
						return y.getEmails(ctx.StringSlice("label"))
					} else {
						return y.getEmailById(args0)
					}
				},
			},
		},
	}
}

func (y *GmailCmd) getEmails(labelIds []string) error {
	svcCall := y.service.Users.Messages.List("me")

	if len(labelIds) > 0 {
		svcCall.LabelIds(labelIds...)
	}

	resp, err := svcCall.Do()
	if err != nil {
		return err
	}

	printJson(resp.Messages)

	return nil
}

func (y *GmailCmd) getEmailById(emailId string) error {
	resp, err := y.service.Users.Messages.Get("me", emailId).Do()
	if err != nil {
		return err
	}

	printJson(resp)

	return nil
}
