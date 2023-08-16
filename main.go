package main

import (
	"fmt"
	"log"
	"spotify-playlist-share/auth"
	"spotify-playlist-share/env/env"
	"spotify-playlist-share/playlist"
	"spotify-playlist-share/youtube"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
)

var accessToken *oauth2.Token
var authError error
var client spotify.Client
var title []string

func init() {
	env.LoadEnv()
	accessToken, authError = auth.LoadAuth()
	if authError != nil {
		log.Fatalf("error retrieve access token: %v", authError)
	}
	client = playlist.StartClient(accessToken)
}

func main() {
	fmt.Println("Enter UserId: ")
	// var input string
	// fmt.Scanln(&input)

	// collection := playlist.GrabAllUsers(client, env.Env.Collection)

	// for _, list := range collection {
	// 	title = append(title, list.Name)
	// 	retrieve := playlist.GrabSongs(client, list.SpotifyPlaylistId)
	// 	go filewrite.WriteSongs(list.Name, retrieve)
	// }

	youtube.Stuff()

}
