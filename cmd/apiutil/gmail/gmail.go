package gmailcmd

import (
	"context"
	"net/http"

	"github.com/urfave/cli/v2"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

type GmailCmd struct {
	command *cli.Command
	client  *http.Client
	service *gmail.Service
}

func NewGmailCmd() *GmailCmd {
	ret := &GmailCmd{}
	ret.initCommand()

	return ret
}

func (g *GmailCmd) GetCommand() *cli.Command {
	return g.command
}

func (g *GmailCmd) initCommand() {
	flags := g.getCmdFlags()

	g.command = &cli.Command{
		Name:  "gmail",
		Usage: "access gmail rest api client",
		Subcommands: []*cli.Command{
			g.getLabelsCommand(),
			g.getEmailsCommand(),
		},
		Flags: flags,
		Before: func(cliCtx *cli.Context) error {
			config, err := g.getConfig(cliCtx.Path("secret"), gmail.GmailReadonlyScope)
			if err != nil {
				return err
			}

			ctx := context.Background()

			if err = g.initClient(ctx, config); err != nil {
				return err
			}

			if err = g.initService(ctx); err != nil {
				return err
			}

			return nil
		},
	}
}

func (g *GmailCmd) getCmdFlags() []cli.Flag {
	return []cli.Flag{
		&cli.PathFlag{
			Name:    "secret",
			Usage:   "path to client secret file",
			Value:   "/etc/gmail/client_secret.json",
			Aliases: []string{"s"},
		},
	}
}

func (g *GmailCmd) initService(ctx context.Context) error {
	var err error
	g.service, err = gmail.NewService(ctx, option.WithHTTPClient(g.client))
	if err != nil {
		return err
	}

	return nil
}
