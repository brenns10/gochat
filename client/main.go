/*
Contains gochat client.
*/
package main

import (
	"github.com/brenns10/gochat"
)

func main() {
	client := gochat.NewChatClient("localhost:1234")
	client.Run()
}
