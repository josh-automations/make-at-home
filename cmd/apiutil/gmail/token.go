package gmailcmd

import (
	"encoding/json"
	"fmt"
	"os"

	"golang.org/x/oauth2"
)

func getTokenFromFile(filePath string) (*oauth2.Token, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("filePath='%s': %w", filePath, err)
	}
	defer f.Close()

	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

func saveTokenToFile(filePath string, token *oauth2.Token) error {
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return fmt.Errorf("error opening file '%s': %w", filePath, err)
	}
	defer f.Close()

	err = json.NewEncoder(f).Encode(token)
	return err
}
