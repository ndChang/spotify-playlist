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
		return true
	} else {
		fmt.Println(s == "0")
	}

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

func CheckPlaylistDB(db *sql.DB, pls []datamodel.Playlist) (map[string]bool, error) {
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
			&it, &ut, &playlist.UpdateCreatorId, &playlist.SpotifyOwnerId, &playlist.SnapshotId, &playlist.Public); err != nil {
			fmt.Println(err)
			return avail, err
		}
		avail[playlist.SpotifyPlaylistId] = true
	}
	return avail, nil

}

func CheckSpotifyUserDB(db *sql.DB, userid string) bool {
	q := fmt.Sprintf("select * from %s.spotify_users where SpotifyUserId='%s'", env.Env.Schema, userid)
	var user tables.SpotifyUser
	var it []uint8
	row := db.QueryRow(q)
	switch err := row.Scan(&user.Id, &user.SpotifyUserId, &it); err {
	case sql.ErrNoRows:
		return false
	case nil:
		fmt.Println("Found: ", user.SpotifyUserId)
		return true
	default:
		panic(err)
	}
	return false
}

func CheckSongDB(db *sql.DB, songs []datamodel.Song) map[string]bool {
	list := ""
	for _, pl := range songs {
		list += fmt.Sprintf("'%s', ", pl.SpotifyId)
	}
	if len(list) > 0 {
		list = list[:len(list)-2]
	}
	avail := make(map[string]bool)
	q := fmt.Sprintf("SELECT * FROM %s.song WHERE spotify_id in (%s)", env.Env.Schema, list)
	rows, err := db.Query(q)

	if err != nil {
		// return nil, err
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var song tables.Song
		if err := rows.Scan(&song.Id, &song.Title, &song.Artist, &song.Youtube_video_id,
			&song.Spotify_id); err != nil {
			fmt.Println(err)
			// return avail, err
		}
		avail[song.Spotify_id] = true
	}
	return avail
}
