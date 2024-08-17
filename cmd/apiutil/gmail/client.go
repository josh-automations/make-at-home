package gmailcmd

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/josh-automations/make-at-home/internal/oauth2util"
	"golang.org/x/oauth2"
)

func getTokenFromWeb(ctx context.Context, config *oauth2.Config) (*oauth2.Token, error) {
	authUrl, _ := oauth2util.GetAuthCodeUrl(config)
	fmt.Printf("Go to the following link in your browser then type the authorization code: \n%v\n",
		authUrl)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		return nil, fmt.Errorf("error reading auth code: %w", err)
	}

	tok, err := oauth2util.GetTokenFromAuthCode(ctx, config, authCode, "")
	if err != nil {
		return nil, err
	}
	return tok, nil
}

func (g *GmailCmd) initClient(ctx context.Context, config *oauth2.Config) error {
	tokenPath := "/etc/gmail/token.json"
	token, err := oauth2util.GetTokenFromFile(tokenPath)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	} else if err != nil {
		token, err = getTokenFromWeb(ctx, config)
		if err != nil {
			return err
		}

		oauth2util.SaveTokenToFile(tokenPath, token)
	}

	g.client = config.Client(ctx, token)
	return nil
}
