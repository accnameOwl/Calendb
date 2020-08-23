package main

import (
	json "encoding/json"
	"fmt"
	"github.com/accnameowl/calendb/src/auth"
	"io/ioutil"
	"net/http"
	"os"
)

func init() {
}

func main() {

	var authTest = AuthTest{
		url:      "https://api.nasa.gov/planetary/apod",
		token:    "?api_key=QeHSqI7jLoVAvghkd0SF0ZJ03v1XT2XlfY4dpBLT",
		protocol: "tcp",
	}

	res, err := authTest.Dial()
	if err != nil {
		log.Error(err)
	}
	fmt.Println("Successfully connected!")
}
