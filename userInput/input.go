package input

import (
	"fmt"
	"os"
)

var input string
var UserId string

func UserIdForPlaylistCollection() (string, int) {
	fmt.Println("Enter User Id of Playlists you would like: ")
	fmt.Println("Go back to prior menu with c ")
	var UserId string
	// Taking UserId from user
	for {
		fmt.Scanln(&UserId)
		if UserId == "c" {
			return "", 1
		} else if UserId == "-help" {
			fmt.Println("Grab ID from url of a spotify playlist")
			fmt.Println("Example: 5CNtRR2z3mUMnVvGLnXj4Q would be the id of ")
			fmt.Println("https://open.spotify.com/playlist/5CNtRR2z3mUMnVvGLnXj4Q")
		} else if len(UserId) < 8 {
			fmt.Println("I doubt this Id exists")
		} else if UserId == "exit" {
			os.Exit(0)
		} else {
			return UserId, 0
		}
	}
}

func Cli() bool {
	fmt.Println("Welcome")
	fmt.Println("Type -help for more information ")
	OpeningLines()

	for {
		// Taking input from user
		fmt.Scanln(&input)

		if input == "exit" {
			os.Exit(0)
		} else if input == "-help" {
			fmt.Println("User 	    -u 		Use UserID to grab more information about a specific profile")
			fmt.Println("Playlist   -p 		Get Songs from a specific playlist")
			fmt.Println("Collection -p 		Get All Playlists from a specific User")
			fmt.Println("Generate   -g 		Create a file of all songs logged")
			fmt.Println("exit				Quit Program")
		} else if input == "User" {
			user, cancellation := UserIdForPlaylistCollection()
			if cancellation == 1 {
				OpeningLines()
				continue
			}
			fmt.Println(user)
		} else {
			fmt.Println("Invalid option: Type -help for more information")
		}
	}

}

func OpeningLines() {
	fmt.Println("What would you like to do?")
}
