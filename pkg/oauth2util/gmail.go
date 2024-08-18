package oauth2util

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GetGoogleConfig(filePath string, scopes ...string) (*oauth2.Config, error) {
	secret, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return google.ConfigFromJSON(secret, scopes...)
}
