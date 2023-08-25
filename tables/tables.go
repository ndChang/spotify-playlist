package tables

import (
	"database/sql"
	"time"
)

type Playlist struct {
	Id                int
	Owner             string
	Name              string
	SpotifyPlaylistId string
	SpotifyOwnerId    string
	SnapshotId        string
	Public            bool
	Insertdatetime    time.Time
	UpdateDateTime    time.Time
	UpdateCreatorId   sql.NullString
}

type SpotifyUser struct {
	Id             int
	SpotifyUserId  string
	Insertdatetime time.Time
}

type Song struct {
	Id               int
	Title            string
	Artist           string
	Youtube_video_id string
	Spotify_id       string
}
