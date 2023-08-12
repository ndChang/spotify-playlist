package auth

import (
	"context"
	"log"
	"spotify-playlist-share/env/env"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

func LoadAuth() (*oauth2.Token, error) {
	authConfig := &clientcredentials.Config{
		ClientID:     env.Env.ClientID,
		ClientSecret: env.Env.ClientSecret,
		TokenURL:     env.Env.TokenURL,
	}

	accessToken, err := authConfig.Token(context.Background())
	if err != nil {
		log.Fatalf("error retrieve access token: %v", err)
	}

	return accessToken, err
}
