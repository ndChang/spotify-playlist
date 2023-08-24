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

func BulkAddPlaylists(db *sql.DB, pls []datamodel.Playlist, dict map[string]bool) {
	sqlStr := fmt.Sprintf("INSERT INTO %s.playlist(owner, name, spodifyplaylistid, insertdatetime, updatedatetime) VALUES ", env.Env.Schema)
	var vals []interface{}
	tm := time.Now()
	counter := 0

	for _, pl := range pls {
		// SpotifyPlaylistId currently does not exist in db, Add values of playlist to query
		if dict[pl.SpotifyPlaylistId] != true {
			sqlStr += "(?, ?, ?, ?, ?),"
			vals = append(vals, pl.PlaylistOwnerDisplayName, pl.Name, pl.SpotifyPlaylistId, tm, tm)
			counter++
		}
	}
	if string(sqlStr[len(sqlStr)-1]) == " " {
		fmt.Println("No entries to add")
		return
	}
	sqlStr = sqlStr[:len(sqlStr)-1]
	//prepare the statement
	stmt, _ := db.Prepare(sqlStr)

	defer stmt.Close()

	stmt.Exec(vals...)

	fmt.Println("Rows added: ", counter)
}
