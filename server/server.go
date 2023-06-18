package server

import (
	"fmt"
	"net/http"
)

func StartServer() {
	// Define a handler function to handle incoming HTTP requests
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!") // Send a response to the client
	}

	// Register the handler function to handle all requests on the root path ("/")
	http.HandleFunc("/", handler)

	// Start the HTTP server on port 8080
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
