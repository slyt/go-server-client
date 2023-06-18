package main

import (
	"go-server/client"
	"go-server/server"
	"time"
)

func main() {

	go server.StartServer()
	// Wait for 1 second
	time.Sleep(1 * time.Second)
	client.StartClient()
	client.StartConcurrentClient()
}
