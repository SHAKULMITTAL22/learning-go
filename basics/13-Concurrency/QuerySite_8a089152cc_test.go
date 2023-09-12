package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"testing"
)

var wg sync.WaitGroup

func querySite(url string, requestBody chan int) {
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

	requestBody <- len(body)
}

func TestQuerySite_8a089152cc(t *testing.T) {
	// Test case 1: Valid URL
	requestBody := make(chan int)
	wg.Add(1)
	go querySite("https://www.google.com", requestBody)
	wg.Wait()
	bodyLength := <-requestBody
	if bodyLength == 0 {
		t.Error("Expected body length to be greater than 0")
	}

	// Test case 2: Invalid URL
	requestBody = make(chan int)
	wg.Add(1)
	go querySite("https://www.google.com/invalid", requestBody)
	wg.Wait()
	bodyLength = <-requestBody
	if bodyLength != 0 {
		t.Error("Expected body length to be 0")
	}
}
