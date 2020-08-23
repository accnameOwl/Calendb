package main

import (
	json "encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Config ...
//
// created from config.json and initialized on init()
var Config map[string]interface{}

func init() {
	// * parse JSON
	cf, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	defer cf.Close()

	data := make([]byte, 100)
	count, err := cf.Read(data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("config.json - Bytes %i ", count)
	fmt.Printf("Bytes: %i", count)

	err = json.Unmarshal([]byte(data), &Config)

}

func main() {
	fmt.Printf("Contents of URL: %s ", Config["url"]+Config["task"]+Config["token"])
	resp, err := http.Get(Config["url"] + Config["task"]["apiKey"] + Config["token"])
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", html)
}
