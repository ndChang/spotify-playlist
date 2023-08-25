package datamodel

type Playlist struct {
	SpotifyPlaylistId        string
	Name                     string
	PlaylistOwnerDisplayName string
	PlaylistOwnerId          string
	SnapshotId               string
	Public                   bool
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

type Song struct {
	Name      string
	Artist    string
	SpotifyId string
	YoutubeId string
}
