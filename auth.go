package main

import (
	"github.com/dghubble/oauth1"
	"golang.org/x/xerrors"
)

type auth struct {
	config        *oauth1.Config
	requestToken  string
	requestSecret string
}

func newAuth(consumerKey string, consumerSecret string) (*auth, error) {
	config := &oauth1.Config{
		ConsumerKey:    consumerKey,
		ConsumerSecret: consumerSecret,
		CallbackURL:    "oob",
		Endpoint: oauth1.Endpoint{
			RequestTokenURL: "https://api.twitter.com/oauth/request_token",
			AuthorizeURL:    "https://api.twitter.com/oauth/authorize",
			AccessTokenURL:  "https://api.twitter.com/oauth/access_token",
		},
	}

	requestToken, requestSecret, err := config.RequestToken()
	if err != nil {
		return nil, xerrors.Errorf("Fetch request token failed: %v", err)
	}

	return &auth{
		config:        config,
		requestToken:  requestToken,
		requestSecret: requestSecret,
	}, nil
}

func (client *auth) getAuthzURL() (string, error) {
	authorizationURL, err := client.config.AuthorizationURL(client.requestToken)
	if err != nil {
		return "", xerrors.Errorf("Fetch Authz URL failed: %v", err)
	}
	return authorizationURL.String(), nil
}

func (client *auth) getTokenAndSecret(pin string) (token, secret string, err error) {
	return client.config.AccessToken(client.requestToken, client.requestSecret, pin)
}
