package database

import (
	"database/sql"
	"fmt"
	"spotify-playlist-share/datamodel"
	"spotify-playlist-share/env/env"
	"strconv"
)

func CheckPlaylistEntry(db *sql.DB, pl datamodel.Playlist) bool {
	insertSql := fmt.Sprintf("select * from %s.playlist where SpodifyPlaylistId='%s'", env.Env.Schema, pl.SpotifyPlaylistId)
	stmlins, err := db.Prepare(insertSql)
	if err != nil {
		// panic(err)
		fmt.Println("ISSUE HERE ", err)
	}
	defer stmlins.Close()
	res, _ := stmlins.Exec()
	r, _ := res.RowsAffected()
	s := strconv.FormatInt(r, 10)
	if s == "0" {
		return false
	}
	// row := db.QueryRow("select * from "+env.Env.PlaylistTable+" where SpodifyPlaylistId =?", pl.SpotifyPlaylistId)

	// var title string

	// err := row.Scan(&title)

	// fmt.Println(title)
	// if err != nil {
	// 	// panic(err)
	// 	fmt.Println("ISSUE HERE READ ", err)
	// 	return false
	// }

	return true
}
