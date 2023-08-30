package database

import (
	"database/sql"
	"fmt"
	"spotify-playlist-share/datamodel"
	"spotify-playlist-share/env/env"
	"spotify-playlist-share/youtuberest"
	"time"
)

func AddPlaylist(db *sql.DB, pl datamodel.Playlist) error {
	insertSql := fmt.Sprintf("INSERT INTO %s.playlist(owner, name, spodifyplaylistid, insertdatetime, updatedatetime, SpotifyOwnerId) VALUES(?,?,?,?,?,?)", env.Env.Schema)
	tm := time.Now()
	stmlins, err := db.Prepare(insertSql)
	if err != nil {
		// panic(err)
		return err
	}
	defer stmlins.Close()

	_, insertErr := stmlins.Exec(pl.PlaylistOwnerId, pl.Name, pl.SpotifyPlaylistId, tm, tm, pl.PlaylistOwnerId)
	if insertErr != nil {
		// panic(err)
		return err
	}
	return nil
}

func BulkAddPlaylists(db *sql.DB, pls []datamodel.Playlist, dict map[string]bool) {
	sqlStr := fmt.Sprintf("INSERT INTO %s.playlist(owner, name, spodifyplaylistid, insertdatetime, updatedatetime, SpotifyOwnerId, Public, SnapshotId) VALUES ", env.Env.Schema)
	var vals []interface{}
	tm := time.Now()
	counter := 0

	for _, pl := range pls {
		// SpotifyPlaylistId currently does not exist in db, Add values of playlist to query
		if dict[pl.SpotifyPlaylistId] != true {
			sqlStr += "(?, ?, ?, ?, ?, ?,?, ?),"
			vals = append(vals, pl.PlaylistOwnerDisplayName, pl.Name, pl.SpotifyPlaylistId, tm, tm, pl.PlaylistOwnerId, pl.Public, pl.SnapshotId)
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

	fmt.Println("Playlists added: ", counter)
}

func AddUser(db *sql.DB, userid string) {
	// q := fmt.Sprintf("select * from %s.spotify_users where SpotifyUserId='%s'", env.Env.Schema, userid)
	q := fmt.Sprintf("INSERT INTO %s.spotify_users(SpotifyUserId, InsertDateTime) VALUES(?, ?)", env.Env.Schema)
	tm := time.Now()
	//prepare the statement
	stmt, _ := db.Prepare(q)

	defer stmt.Close()

	stmt.Exec(userid, tm)
	fmt.Println("User: " + userid + " added")
}

func AddSongs(db *sql.DB, songs *[]datamodel.Song, dict map[string]bool) bool {
	sqlStr := fmt.Sprintf("INSERT INTO %s.song(title, artist, youtube_video_id, spotify_id) VALUES ", env.Env.Schema)
	var vals []interface{}
	counter := 0

	for _, song := range *songs {
		// SpotifyPlaylistId currently does not exist in db, Add values of playlist to query
		if dict[song.SpotifyId] != true {
			// check youtube for video url
			song.YoutubeId = youtuberest.GetYoutubeVideoId(song)
			sqlStr += "(?, ?, ?, ?),"
			vals = append(vals, song.Name, song.Artist, song.YoutubeId, song.SpotifyId)
			counter++
		}
	}
	if string(sqlStr[len(sqlStr)-1]) == " " {
		fmt.Println("No entries to add")
		return false
	}
	sqlStr = sqlStr[:len(sqlStr)-1]
	//prepare the statement
	stmt, _ := db.Prepare(sqlStr)

	defer stmt.Close()

	stmt.Exec(vals...)

	fmt.Println("Songs added: ", counter)
	return true
}
