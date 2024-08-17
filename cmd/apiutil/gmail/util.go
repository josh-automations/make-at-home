package gmailcmd

import (
	"encoding/json"
	"fmt"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func (g *GmailCmd) getConfig(filePath string, scopes ...string) (*oauth2.Config, error) {
	secret, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return google.ConfigFromJSON(secret, scopes...)
}

func printJson(v any) error {
	res, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(res))
	return nil
}
