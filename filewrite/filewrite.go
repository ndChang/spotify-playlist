package filewrite

import (
	"fmt"
	"log"
	"os"
	"spotify-playlist-share/datamodel"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func createDirectory(id string, dirName string) {
	_, err := os.Stat("./spotify_list_" + id + "/" + dirName)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll("./spotify_list_"+id+"/"+dirName, 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	} else if err != nil {
		log.Fatal("error creating dir")
	}
}

func createTextFile(id string, dirName string, songs string) {
	d1 := []byte(songs)
	err := os.WriteFile("./spotify_list_"+id+"/"+dirName+"/"+dirName+".txt", d1, 0644)
	check(err)

	//not sure if below code is needed
	f, err := os.Create("/tmp/dat2")
	check(err)

	defer f.Close()

}

func WriteSongs(id string, dirName string, songs []datamodel.Song) {
	createDirectory(id, dirName)
	song := songParser(songs)
	createTextFile(id, dirName, song)
}

func songParser(songs []datamodel.Song) string {
	songlist := ""

	for _, song := range songs {
		songlist += song.Name + " " + song.Artist + " https://www.youtube.com/watch?v=" + song.YoutubeId + "\n"
	}
	return songlist
}

func CleanPlaylistDirectory(id string) {
	_, err := os.Stat("./spotify_list_" + id)
	if os.IsNotExist(err) {
		fmt.Println("Directory spotify_list does not exist")
	} else {
		fmt.Println("Removing directory")
	}
	err = os.RemoveAll("./spotify_list" + id)
	check(err)

	fmt.Println("Directory removed")
}
