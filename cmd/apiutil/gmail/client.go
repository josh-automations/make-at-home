package gmailcmd

import (
	"context"
	"errors"
	"fmt"
	"os"

	"golang.org/x/oauth2"
)

func getTokenFromWeb(ctx context.Context, config *oauth2.Config) (*oauth2.Token, error) {
	verifier, err := getRandomString(16)
	if err != nil {
		return nil, err
	}

	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline,
		oauth2.S256ChallengeOption(verifier))
	fmt.Printf("Go to the following link in your browser then type the authorization code: \n%v\n",
		authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		return nil, fmt.Errorf("error reading auth code: %w", err)
	}

	tok, err := config.Exchange(ctx, authCode)
	if err != nil {
		return nil, err
	}
	return tok, nil
}

func (g *GmailCmd) initClient(ctx context.Context, config *oauth2.Config) error {
	tokenPath := "/etc/gmail/token.json"
	token, err := getTokenFromFile(tokenPath)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	} else if err != nil {
		token, err = getTokenFromWeb(ctx, config)
		if err != nil {
			return err
		}

		saveTokenToFile(tokenPath, token)
	}

	g.client = config.Client(ctx, token)
	return nil
}
