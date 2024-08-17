package gmailcmd

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
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

// Generates a random ascii string using human-readable characters
// but excludes SPACE
func getRandomString(len int) (string, error) {
	var minVal, maxVal int64
	minVal = 33  //!
	maxVal = 126 //~

	ret := make([]byte, len)
	for i := range len {
		if val, err := rand.Int(rand.Reader, big.NewInt(maxVal-minVal+1)); err != nil {
			return "", err
		} else {
			ret[i] = byte(val.Int64())
		}
	}

	return string(ret), nil
}
