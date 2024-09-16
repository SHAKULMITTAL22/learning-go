package main

import (
    "fmt"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHelloRequest_aff8a82d01(t *testing.T) {
    // Create a request to the helloRequest function
    req, err := http.NewRequest("GET", "/hello", nil)
    if err != nil {
        t.Fatal(err)
    }

    // Create a ResponseRecorder to record the response
    rr := httptest.NewRecorder()

    // Call the helloRequest function with the request and ResponseRecorder
    helloRequest(rr, req)

    // Check the status code of the response
    if rr.Code != http.StatusOK {
        t.Errorf("Expected status code 200, got %d", rr.Code)
    }

    // Check the body of the response
    if rr.Body.String() != "Hello World!" {
        t.Errorf("Expected body 'Hello World!', got '%s'", rr.Body.String())
    }
}

func TestHelloRequest_Error(t *testing.T) {
    // Create a request to the helloRequest function with an invalid method
    req, err := http.NewRequest("POST", "/hello", nil)
    if err != nil {
        t.Fatal(err)
    }

    // Create a ResponseRecorder to record the response
    rr := httptest.NewRecorder()

    // Call the helloRequest function with the request and ResponseRecorder
    helloRequest(rr, req)

    // Check the status code of the response
    if rr.Code != http.StatusMethodNotAllowed {
        t.Errorf("Expected status code 405, got %d", rr.Code)
    }

    // Check the body of the response
    if rr.Body.String() != "Method Not Allowed" {
        t.Errorf("Expected body 'Method Not Allowed', got '%s'", rr.Body.String())
    }
}
