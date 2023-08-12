package main

import (
	"fmt"
	"log"
	"spotify-playlist-share/auth"
	"spotify-playlist-share/env/env"
	"spotify-playlist-share/playlist"
	input "spotify-playlist-share/userInput"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
)

var accessToken *oauth2.Token
var authError error
var client spotify.Client

func init() {
	env.LoadEnv()
	accessToken, authError = auth.LoadAuth()
	if authError != nil {
		log.Fatalf("error retrieve access token: %v", authError)
	}
	client = playlist.StartClient(accessToken)
}

func main() {
	UserId := input.UserIdForPlaylistCollection()
	fmt.Println(UserId)
}
