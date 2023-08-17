package playlist

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"spotify-playlist-share/datamodel"
	"spotify-playlist-share/env/env"
	"strings"

	"github.com/zmb3/spotify"
)

type dummyResp struct {
	Body []datamodel.YoutubeResponse
}

func GrabDummySongs(client spotify.Client, playlistId string) []string {
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
		resp := SimulateResp("https://www.googleapis.com/youtube/v3/search?part=snippet&maxResults=1&q=" + query + "&type=video&key=" + env.Env.YoutubeApi)
		if err := json.Unmarshal(resp, &ytr); err != nil {
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

func SimulateResp(url string) []byte {
	file, err := os.Open("./env/svr/response.json")
	if err != nil {
		fmt.Println("Error opening file")
	}

	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)
	return byteValue
}
