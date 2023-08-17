package env

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var Env *Config

type Config struct {
	ClientID      string `json:"ClientID"`
	ClientSecret  string `json:"ClientSecret"`
	TokenURL      string `json:"TokenUrl"`
	Collection    string `json:"collection"`
	YoutubeConfig string `json:"youtubeConfig"`
	YoutubeApi    string `json:"youtubeApiKey"`
}

func newConfig() *Config {

	conf := Config{}

	file, err := os.Open("./env/svr/config.json")
	if err != nil {
		fmt.Println("Error opening file")
	}

	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	err = json.Unmarshal(byteValue, &conf)
	if err != nil {
		fmt.Println("Error opening file")
	}
	return &conf
}

func LoadEnv() {
	fmt.Println("Loading Environment Variabls")

	Env = newConfig()
}
