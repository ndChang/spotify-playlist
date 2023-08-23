package filewrite

import "os"

func CreateDummyJson(filename string, payload []byte) {
	err := os.WriteFile("./env/svr/"+filename+".json", payload, 0644)
	check(err)
}
