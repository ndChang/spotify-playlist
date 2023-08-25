package playlist

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"spotify-playlist-share/datamodel"
	"spotify-playlist-share/env/env"
	"strings"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
)

func StartClient(accessToken *oauth2.Token) spotify.Client {
	client := spotify.Authenticator{}.NewClient(accessToken)
	fmt.Println("Connected to Spotify")
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
		song := value.Track.SimpleTrack.Name + " by " + value.Track.Album.Artists[0].Name
		split := strings.Split(song, " ")
		query := strings.Join(split, "%20")

		var ytr datamodel.YoutubeResponse
		resp, err := http.Get("https://www.googleapis.com/youtube/v3/search?part=snippet&maxResults=1&q=" + query + "&type=video&key=" + env.Env.YoutubeApi)
		if err != nil || resp.StatusCode != 200 {
			panic("Issue with api call")
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err := json.Unmarshal(body, &ytr); err != nil {
			fmt.Println("error", err)
		}

		if len(ytr.Items) > 0 {
			entry := song + " " + "https://www.youtube.com/watch?v=" + ytr.Items[0].Id.VideoId
			list = append(list, entry)
		}

		// fmt.Println(ytr.Items[0].Id.VideoId)

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
		tempPlaylist := datamodel.Playlist{SpotifyPlaylistId: Playlist.ID.String(), Name: Playlist.Name, PlaylistOwnerDisplayName: Playlist.Owner.DisplayName, PlaylistOwnerId: Playlist.Owner.ID, SnapshotId: Playlist.SnapshotID, Public: Playlist.IsPublic}
		UserPlaylistsCollection = append(UserPlaylistsCollection, tempPlaylist)
	}

	return UserPlaylistsCollection
}
