package main

import (
	"fmt"
	"net/http"
	"time"
)

func httpPing(url string) error {
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	// Make a GET request
	resp, err := client.Get(url)
	if err != nil {
		return err // Handle error
	}
	defer resp.Body.Close() // Close the response body when the function returns

	// Check if the status code is a success
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		fmt.Printf("Ping to %s succeeded with status code: %d\n", url, resp.StatusCode)
	} else {
		fmt.Printf("Ping to %s failed with status code: %d\n", url, resp.StatusCode)
	}
	return nil
}

func main() {
	url := "http://www.google.com"
	err := httpPing(url)
	if err != nil {
		fmt.Printf("Error pinging %s: %s\n", url, err)
	}
}
