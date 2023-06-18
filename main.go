package main

import (
	"go-server/client"
	"go-server/server"
)

func main() {
	server.StartServer()
	client.StartClient()
}
