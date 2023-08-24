package database

import (
	"database/sql"
	"fmt"
	"spotify-playlist-share/datamodel"
	"spotify-playlist-share/env/env"
	"spotify-playlist-share/tables"
	"strconv"
)

func CheckPlaylistEntry(db *sql.DB, pl datamodel.Playlist) bool {
	insertSql := fmt.Sprintf("select * from %s.playlist where SpodifyPlaylistId='%s'", env.Env.Schema, pl.SpotifyPlaylistId)
	fmt.Println(insertSql)
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
		fmt.Println("HIT HERE")
		return true
	} else {
		fmt.Println(s == "0")
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

	return false
}

func CheckAllPlaylistEntries(db *sql.DB, pls []datamodel.Playlist) {
	list := ""
	for _, pl := range pls {
		list += fmt.Sprintf("'%s', ", pl.SpotifyPlaylistId)
	}
	list += fmt.Sprintf("select * from %s.playlist where SpodifyPlaylistId in (%s)", env.Env.Schema, list)
	fmt.Println(list)
}

func Check(db *sql.DB, pls []datamodel.Playlist) (map[string]bool, error) {
	list := ""
	for _, pl := range pls {
		list += fmt.Sprintf("'%s', ", pl.SpotifyPlaylistId)
	}
	if len(list) > 0 {
		list = list[:len(list)-2]
	}
	avail := make(map[string]bool)
	q := fmt.Sprintf("SELECT * FROM %s.playlist WHERE SpodifyPlaylistId in (%s)", env.Env.Schema, list)
	rows, err := db.Query(q)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var playlist tables.Playlist
		var it, ut []uint8
		if err := rows.Scan(&playlist.Id, &playlist.Name, &playlist.Owner, &playlist.SpotifyPlaylistId,
			&it, &ut, &playlist.UpdateCreatorId); err != nil {
			fmt.Println(err)
			return avail, err
		}
		avail[playlist.SpotifyPlaylistId] = true
	}
	return avail, nil

}
