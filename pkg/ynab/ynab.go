package ynab

import (
	"context"
	"net/http"
)

const DefaultServer string = "https://api.ynab.com/v1"

func bearerTokenSetter(apiToken string) func(ctx context.Context, req *http.Request) error {
	return func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", "Bearer "+apiToken)
		return nil
	}
}

func NewYnabClient(server string, apiToken string) (*ClientWithResponses, error) {
	return NewClientWithResponses(server, WithRequestEditorFn(bearerTokenSetter(apiToken)))
}
