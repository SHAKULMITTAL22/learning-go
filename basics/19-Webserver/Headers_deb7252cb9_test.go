package main

import (
    "fmt"
    "log"
    "net/http"
    "testing"
)

func headers(w http.ResponseWriter, req *http.Request) {
    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func TestHeaders(t *testing.T) {
    // Test case 1: Success case
    req, err := http.NewRequest("GET", "http://example.com", nil)
    if err != nil {
        t.Error(err)
    }
    req.Header.Set("Content-Type", "application/json")
    w := http.ResponseWriter{}
    headers(w, req)
    expected := "Content-Type: application/json\n"
    if w.Header().Get("Content-Type") != expected {
        t.Errorf("Expected %s, got %s", expected, w.Header().Get("Content-Type"))
    }

    // Test case 2: Failure case
    req, err = http.NewRequest("GET", "http://example.com", nil)
    if err != nil {
        t.Error(err)
    }
    req.Header.Set("Content-Type", "application/xml")
    w = http.ResponseWriter{}
    headers(w, req)
    expected = "Content-Type: application/xml\n"
    if w.Header().Get("Content-Type") != expected {
        t.Errorf("Expected %s, got %s", expected, w.Header().Get("Content-Type"))
    }
}
