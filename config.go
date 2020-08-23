package main

// Config ...
type Config struct {
	token string
	url   string
}

// Configuration ...
// Global configurations on initiation
var Configuration = Config{
	token: "QeHSqI7jLoVAvghkd0SF0ZJ03v1XT2XlfY4dpBLT",
	url:   "https://api.nasa.gov",
}
