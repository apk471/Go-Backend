package main

import (
	"context"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var oauthConfig = &oauth2.Config{
	ClientID:     "GOOGLE_CLIENT_ID",
	ClientSecret: "GOOGLE_CLIENT_SECRET",
	RedirectURL:  "http://localhost:8080/oauth/callback",
	Scopes:       []string{"email", "profile"},
	Endpoint:     google.Endpoint,
}

func exchangeCode(code string) (*oauth2.Token, error) {
	return oauthConfig.Exchange(context.Background(), code)
}
