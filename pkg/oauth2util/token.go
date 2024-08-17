package oauth2util

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"golang.org/x/oauth2"
)

// Returns a tuple (url, stateChecker)
//
//	url - the AuthCode URL.
//	stateChecker - a function used to verify the state. After user authenticates,
//	  `stateChecker(state)` returns true if the response state is valid, false otherwise.
//
// Parameters:
//
//	config - the oauth2 config
func GetAuthCodeUrl(config *oauth2.Config) (string, func(string) bool) {
	state := oauth2.GenerateVerifier()
	verifier := oauth2.GenerateVerifier()

	authUrl := config.AuthCodeURL(state, oauth2.AccessTypeOffline,
		oauth2.S256ChallengeOption(verifier))
	return authUrl, func(respState string) bool { return respState == state }
}

func GetTokenFromAuthCode(ctx context.Context, config *oauth2.Config, authCode string,
	verifier string) (*oauth2.Token, error) {
	opts := []oauth2.AuthCodeOption{}

	if verifier != "" {
		opts = append(opts, oauth2.VerifierOption(verifier))
	}
	return config.Exchange(ctx, authCode, opts...)
}

func GetTokenFromFile(filePath string) (*oauth2.Token, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("filePath='%s': %w", filePath, err)
	}
	defer f.Close()

	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

func SaveTokenToFile(filePath string, token *oauth2.Token) error {
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return fmt.Errorf("error opening file '%s': %w", filePath, err)
	}
	defer f.Close()

	err = json.NewEncoder(f).Encode(token)
	return err
}
