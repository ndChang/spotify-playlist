package playlist

import (
	"fmt"
	"log"
	"spotify-playlist-share/datamodel"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
)

func StartClient(accessToken *oauth2.Token) spotify.Client {
	client := spotify.Authenticator{}.NewClient(accessToken)
	return client

}

func GrabSongs(client spotify.Client, playlistId string) []string {
	playlistSpotifyID := spotify.ID(playlistId)
	playlist, err := client.GetPlaylist(playlistSpotifyID)
	if err != nil {
		log.Fatalf("error retrieve playlist data: %v", err)
	}

	var list []string
	for _, value := range playlist.Tracks.Tracks {
		// fmt.Println(value.Track.Album.Artists[0].Name)
		list = append(list, value.Track.SimpleTrack.Name+" by "+value.Track.Album.Artists[0].Name)
	}
	return list
}

func GrabAllUsers(client spotify.Client, playlistId string) []datamodel.Playlist {
	var UserPlaylistsCollection []datamodel.Playlist
	userPlaylists, getPlaylistErr := client.GetPlaylistsForUser("31ttjryp6mvbrrgsd64j2arbskda")
	if getPlaylistErr != nil {
		log.Fatalf("error retrieve playlist data: %v", getPlaylistErr)
	}
	for _, Playlist := range userPlaylists.Playlists {
		tempPlaylist := datamodel.Playlist{SpotifyPlaylistId: Playlist.ID.String(), Name: Playlist.Name, PlaylistOwnerDisplayName: Playlist.Owner.DisplayName, PlaylistOwnerId: Playlist.Owner.ID}
		UserPlaylistsCollection = append(UserPlaylistsCollection, tempPlaylist)
	}

	fmt.Println(len(UserPlaylistsCollection))
	return UserPlaylistsCollection
}
