/*
Runs the gochat server.
*/
package main

import (
	"github.com/brenns10/gochat"
)

func main() {
	srv := gochat.NewChatServer()
	srv.Run(":1234")
}
