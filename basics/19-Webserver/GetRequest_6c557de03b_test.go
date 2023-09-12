package main

import (
    "fmt"
    "log"
    "net/http"
    "testing"
)

func TestGetRequest_6c557de03b(t *testing.T) {
    // Test case 1: Success case
    req, err := http.NewRequest("GET", "http://localhost:8080/index.html", nil)
    if err != nil {
        t.Error(err)
    }
    w := http.ResponseWriter{}
    getRequest(w, req)
    if w.StatusCode != 200 {
        t.Errorf("Expected status code 200, got %d", w.StatusCode)
    }

    // Test case 2: Failure case
    req, err = http.NewRequest("GET", "http://localhost:8080/non-existent-file.html", nil)
    if err != nil {
        t.Error(err)
    }
    w = http.ResponseWriter{}
    getRequest(w, req)
    if w.StatusCode != 404 {
        t.Errorf("Expected status code 404, got %d", w.StatusCode)
    }
}
