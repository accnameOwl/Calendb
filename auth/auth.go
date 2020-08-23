package auth

import (
	"fmt"
	"net"
)

// AuthTest ...
type AuthTest struct {
	url      string
	token    string
	protocol string
}

// Dial ...
func (at *AuthTest) Dial() (bool, error) {
	conn, err := net.Dial(at.protocol, at.url+at.token)
	if err != nil {
		fmt.Printf("Could not successfully recieve connection from: %s\n\n", at.url)
		fmt.Printf("auth token: %s+nprotocol%s\n", at.token, at.protocol)
		return false, err
	}
	fmt.Printf("Successfully made connection to: %s\n", at.url)
	fmt.Printf("auth token: %s+nprotocol%s\n", at.token, at.protocol)
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	return true, nil
}
