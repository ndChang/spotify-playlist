package filewrite

import (
	"log"
	"os"
	"spotify-playlist-share/datamodel"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func createDirectory(dirName string) {
	_, err := os.Stat("./spodify_list/" + dirName)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll("./spodify_list/"+dirName, 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	} else if err != nil {
		log.Fatal("error creating dir")
	}
}

func createTextFile(dirName string, songs string) {
	d1 := []byte(songs)
	err := os.WriteFile("./spodify_list/"+dirName+"/"+dirName+".txt", d1, 0644)
	check(err)

	//not sure if below code is needed
	f, err := os.Create("/tmp/dat2")
	check(err)

	defer f.Close()

}

func WriteSongs(dirName string, songs []datamodel.Song) {
	createDirectory(dirName)
	song := songParser(songs)
	createTextFile(dirName, song)
}

func songParser(songs []datamodel.Song) string {
	songlist := ""

	for _, song := range songs {
		songlist += song.Name + " " + song.Artist + " " + song.YoutubeId + "\n"
	}
	return songlist

}
