package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func StartServer() {
	// Define a handler function to handle incoming HTTP requests
	handler := func(w http.ResponseWriter, r *http.Request) {
		response := map[string]string{"message": "Hello, World!"}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//fmt.Println("Server replied with Hello, World! in JSON format") // Send a response to the client
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	}

	// Register the handler function to handle all requests on the root path ("/")
	http.HandleFunc("/", handler)

	// Start the HTTP server on any free
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)

}
