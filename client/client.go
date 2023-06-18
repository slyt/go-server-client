package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func StartClient() {
	// Send an HTTP GET request to a URL
	response, err := http.Get("http://localhost:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the response body as a string
	fmt.Println(string(body))
}
