package youtuberest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"spotify-playlist-share/datamodel"
	"spotify-playlist-share/env/env"
	"strings"
)

func GetYoutubeVideoId(track datamodel.Song) string {
	song := track.Name + " " + track.Artist
	split := strings.Split(song, " ")
	query := strings.Join(split, "%20")
	var ytr datamodel.YoutubeResponse
	// Each Search has a cost of 100 with daily free limit of 10000. We are only ever able to use this function 100 times every day.
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
		youtubeurl := ytr.Items[0].Id.VideoId
		return youtubeurl
	} else {
		return ""
	}
}
