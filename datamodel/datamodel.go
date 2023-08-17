package datamodel

type Playlist struct {
	SpotifyPlaylistId        string
	Name                     string
	PlaylistOwnerDisplayName string
	PlaylistOwnerId          string
}

type YoutubeResponse struct {
	Kind  string   `json:"kind"`
	Items []YTItem `json:"items"`
}

type YTItem struct {
	Id YTId `json:"id"`
}

type YTId struct {
	VideoId string `json:"videoId"`
}
