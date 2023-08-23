package database

import (
	"database/sql"
	"fmt"
	"spotify-playlist-share/datamodel"
	"spotify-playlist-share/env/env"
	"time"
)

func AddPlaylist(db *sql.DB, pl datamodel.Playlist) error {
	insertSql := fmt.Sprintf("INSERT INTO %s.playlist(owner, name, spodifyplaylistid, insertdatetime, updatedatetime) VALUES(?,?,?,?,?)", env.Env.Schema)
	tm := time.Now()
	stmlins, err := db.Prepare(insertSql)
	if err != nil {
		// panic(err)
		return err
	}
	defer stmlins.Close()

	_, insertErr := stmlins.Exec(pl.PlaylistOwnerId, pl.Name, pl.SpotifyPlaylistId, tm, tm)
	if insertErr != nil {
		// panic(err)
		return err
	}
	return nil
}
