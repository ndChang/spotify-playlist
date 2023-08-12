package input

import "fmt"

func UserIdForPlaylistCollection() string {
	var input string
	fmt.Println("Enter User Id of Playlists you would like: ")
	fmt.Println("Type -help for more information ")
	for {
		// Taking input from user
		fmt.Scanln(&input)

		if input == "-" {
			fmt.Println("Invalid option: Type -help for more information")
		} else if input == "-help" {
			fmt.Println("Grab ID from url of a spotify playlist")
			fmt.Println("Example: 5CNtRR2z3mUMnVvGLnXj4Q would be the id of ")
			fmt.Println("https://open.spotify.com/playlist/5CNtRR2z3mUMnVvGLnXj4Q")
		} else if len(input) < 8 {
			fmt.Println("I doubt this Id exists")
		} else {
			return input
		}
	}
}
