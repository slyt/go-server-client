package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type APIResponse struct {
	Data  json.RawMessage `json:"data"`
	Error error           `json:"error"`
}

func fetchData(url string, wg *sync.WaitGroup, results chan<- APIResponse) {
	defer wg.Done()

	startTime := time.Now() // start measuring request time of server

	resp, err := http.Get(url)
	if err != nil {
		results <- APIResponse{Error: err}
		return
	}

	// Measure the response time
	duration := time.Since(startTime)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		results <- APIResponse{Error: err}
		return
	}

	// Process the API response
	var data json.RawMessage
	// Modify the following block according to the structure of the API response.
	err = json.Unmarshal(body, &data)
	if err != nil {
		results <- APIResponse{Error: err}
		return
	}

	// print timing informaton for the request
	//fmt.Printf("%d Response time as reported by server via X-Request-time: %s\n", resp.Header.Get("X-Request-Time"))
	fmt.Printf("Response time measured by client: %s\n", duration)

	results <- APIResponse{Data: data}
}

func StartConcurrentClient(concurrency_count int) {
	apiURL := "http://localhost:8080"

	var wg sync.WaitGroup
	results := make(chan APIResponse)

	for i := 0; i < concurrency_count; i++ { // Spawn 10 go routines
		wg.Add(1)
		go fetchData(apiURL, &wg, results)
	}

	go func() { // Spawn a go routine to wait for all the requests to be processed so that we can continue processing of responses that are received
		wg.Wait()
		close(results)
	}()

	for result := range results {
		if result.Error != nil {
			fmt.Printf("Error occured in in concurrent_client.go: %s\n", result.Error)
		} // else {
		// 	dataStr := string(result.Data)
		// 	fmt.Printf("Data recieved in concurrent_client.go: %s\n", dataStr)
		// }

	}

}
