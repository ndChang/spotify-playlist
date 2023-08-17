package main

import (
	"log"
	"spotify-playlist-share/auth"
	"spotify-playlist-share/env/env"
	"spotify-playlist-share/filewrite"
	"spotify-playlist-share/playlist"
	"spotify-playlist-share/youtubeapi"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
	"google.golang.org/api/youtube/v3"
)

var accessToken *oauth2.Token
var authError error
var client spotify.Client
var youtubeClient youtube.Service
var title []string

func init() {
	env.LoadEnv()
	accessToken, authError = auth.LoadAuth()
	if authError != nil {
		log.Fatalf("error retrieve access token: %v", authError)
	}
	client = playlist.StartClient(accessToken)
	youtubeClient = youtubeapi.StartClient()
}

func main() {
	// fmt.Println("Enter UserId: ")
	// youtubeapi.FindVideo(&youtubeClient, "")
	// var input string
	// fmt.Scanln(&input)

	collection := playlist.GrabAllUsers(client, env.Env.Collection)

	for _, list := range collection {
		title = append(title, list.Name)
		// playlist.GrabSongs(client, list.SpotifyPlaylistId)
		retrieve := playlist.GrabDummySongs(client, list.SpotifyPlaylistId)
		// resp, err := http.Get("https://www.googleapis.com/youtube/v3/search?part=snippet&maxResults=20&q=" + query + "&type=video&key=AIzaSyCAyXUobLITDElgedb3SbRIs67sBWlDAGQ")
		go filewrite.WriteSongs(list.Name, retrieve)
	}

}
