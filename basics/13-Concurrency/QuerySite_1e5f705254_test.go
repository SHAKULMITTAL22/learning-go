package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"
)

var wg sync.WaitGroup

func querySite(url string) {
	// Tell the WaitGroup that the goroutine is complete
	defer wg.Done()

	fmt.Println("HTTP Request: ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	fmt.Println("Read Body: ", url)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Body length: ", len(body))
}

func TestQuerySite_1e5f705254(t *testing.T) {
	// Test case 1: Valid URL
	url := "https://www.google.com"
	wg.Add(1)
	go querySite(url)
	wg.Wait()

	// Test case 2: Invalid URL
	url = "https://www.invalidurl.com"
	wg.Add(1)
	go querySite(url)
	wg.Wait()
	// TODO: Add more test cases as needed
}
