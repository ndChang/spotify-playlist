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
	Insertdatetime    time.Time
	UpdateDateTime    time.Time
	UpdateCreatorId   sql.NullString
}
