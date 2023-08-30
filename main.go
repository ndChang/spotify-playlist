package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"spotify-playlist-share/auth"
	"spotify-playlist-share/database"
	"spotify-playlist-share/datamodel"
	"spotify-playlist-share/env/env"
	"spotify-playlist-share/filewrite"
	"spotify-playlist-share/playlist"
	"spotify-playlist-share/youtubeapi"
	"sync"

	"github.com/go-sql-driver/mysql"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
	"google.golang.org/api/youtube/v3"
)

var accessToken *oauth2.Token
var authError error
var client spotify.Client
var youtubeClient youtube.Service
var title []string
var wg sync.WaitGroup
var mysqlerr error
var db *sql.DB
var retrieve []datamodel.Song
var userid string

func init() {
	env.LoadEnv()
	userid = "31ttjryp6mvbrrgsd64j2arbskda"
	accessToken, authError = auth.LoadAuth()
	if authError != nil {
		log.Fatalf("error retrieve access token: %v", authError)
	}
	client = playlist.StartClient(accessToken)
	youtubeClient = youtubeapi.StartClient()
}

func main() {
	cfg := mysql.Config{
		User:                 env.Env.User,
		Passwd:               env.Env.Passwd,
		Net:                  env.Env.Net,
		Addr:                 env.Env.Addr,
		DBName:               env.Env.DBName,
		AllowNativePasswords: env.Env.AllowNativePasswords,
	}
	db, mysqlerr = sql.Open("mysql", cfg.FormatDSN())
	if mysqlerr != nil {
		fmt.Println(mysqlerr)
	}
	defer db.Close()
	ping := db.Ping()
	if ping != nil {
		log.Fatal(ping)
	}
	fmt.Println("Connected!")

	// Check Db for User if doesnt exist add user
	inudb := database.CheckSpotifyUserDB(db, userid)
	if inudb != true {
		database.AddUser(db, userid)
	} else {
		filewrite.CleanPlaylistDirectory(userid)
	}

	// Grab all playlists from a user
	collection := playlist.GrabAllUsers(client, env.Env.Collection)
	// Generate a svr file to monitor changes
	// DummyGeneration("SnapshotId", collection) //SnapshotId - monitors changes to playlist // PlaylistId does not changes
	// Fheck Db for collection of playlists returning a map of playlists
	indb, err := database.CheckPlaylistDB(db, collection)
	if err != nil {
		fmt.Println("Error in check: ", err)
		return
	}
	// if a Playlist does not exist add playlist to database
	database.BulkAddPlaylists(db, collection, indb)

	for _, list := range collection {
		title = append(title, list.Name)
		if database.CheckPlaylistEntry(db, list) == false {
			plErr := database.AddPlaylist(db, list)
			if plErr != nil {
				log.Fatal("Playlist failed to add to db")
			}
			fmt.Println("Playlist Added to db")
		} else {
			fmt.Println("Playlist Exists in db")

		}
		// // For each playlist loop over tracks and contact youtube for video id
		retrieve := playlist.GrabSongs(client, list.SpotifyPlaylistId)
		// retrieve = playlist.GrabDummySongs(client, list.SpotifyPlaylistId)
		if len(retrieve) > 0 {
			retrieve, insdb := database.CheckSongDB(db, retrieve, true) // map[song.Spotify_id]bool check to add songs to db
			repull := database.AddSongs(db, &retrieve, insdb)
			if repull == true {
				fmt.Println("Checkdb")
				retrieve, insdb = database.CheckSongDB(db, retrieve, false) // secondary check to remap songs
			}
			//	// List of songs are sent to file writer to generate folder and list of songs
			filewrite.WriteSongs(userid, list.Name, retrieve)
		} else {
			fmt.Println("Empty Playlist")
		}
	}
	wg.Wait()

}

func DummyGeneration(file string, payload []datamodel.Playlist) {
	j, _ := json.MarshalIndent(payload, "", "  ")
	filewrite.CreateDummyJson(file, j)
}

func LoadDummyPlaylist() []datamodel.Playlist {
	data := []datamodel.Playlist{}
	file, _ := ioutil.ReadFile("./env/svr/GrabAllUsers.json")
	_ = json.Unmarshal([]byte(file), &data)
	return data
}

func CheckUser() {
	inudb := database.CheckSpotifyUserDB(db, "31ttjryp6mvbrrgsd64j2arbskda")
	if inudb != true {
		database.AddUser(db, "31ttjryp6mvbrrgsd64j2arbskda")
	}
}

func Console() {
	// fmt.Println("Enter UserId: ")
	// youtubeapi.FindVideo(&youtubeClient, "")
	// var input string
	// fmt.Scanln(&input)
}

func CheckPlaylist() {
	// Grab all playlists from a user
	collection := playlist.GrabAllUsers(client, env.Env.Collection)
	indb, err := database.CheckPlaylistDB(db, collection)
	if err != nil {
		fmt.Println("Error in check: ", err)
		return
	}
	database.BulkAddPlaylists(db, collection, indb)
}

func DummyUserCollectionGeneration() {
	// Create Json of Response
	collection := LoadDummyPlaylist()
	DummyGeneration("GrabAllUsers", collection)

}

func ProcessCollection() {
	// for _, list := range collection {
	// 	title = append(title, list.Name)
	// 	if database.CheckPlaylistEntry(db, list) == false {
	// 		plErr := database.AddPlaylist(db, list)
	// 		if plErr != nil {
	// 			log.Fatal("Playlist failed to add to db")
	// 		}
	// 		fmt.Println("Playlist Added to db")
	// 	} else {
	// 		fmt.Println("Playlist Exists in db")

	// 	}
	// 	retrieve := playlist.GrabSongs(client, list.SpotifyPlaylistId)
	// 	// retrieve := playlist.GrabDummySongs(client, list.SpotifyPlaylistId)
	// 	filewrite.WriteSongs(list.Name, retrieve)
	// }
}
