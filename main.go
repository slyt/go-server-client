package main

import (
	"flag"
	"fmt"
	"go-server/client"
	"go-server/server"
	"time"
)

func main() {

	var concurrency_count int
	flag.IntVar(&concurrency_count, "c", 2, "Number of concurrent clients")
	flag.Parse()
	fmt.Println("Number of concurrent clients: ", concurrency_count)

	go server.StartServer()
	// Wait for 1 second
	time.Sleep(1 * time.Second)
	client.StartClient()

	// measure total thread execution time
	startTime := time.Now()
	client.StartConcurrentClient(concurrency_count)
	duration := time.Since(startTime)
	fmt.Printf("Total time taken for %d concurrent requests: %s\n", concurrency_count, duration)
}
