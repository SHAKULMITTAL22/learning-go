package router

import (
    "github.com/gorilla/mux"
    "github.com/tannergabriel/learning-go/beginner-programs/todo-list/backend/middleware"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestRouter_522df6b55a(t *testing.T) {
    // Create a new router
    router := Router()

    // Create a new request
    req, err := http.NewRequest("GET", "/todo", nil)
    if err != nil {
        t.Fatal(err)
    }

    // Create a new response recorder
    rr := httptest.NewRecorder()

    // Serve the request
    router.ServeHTTP(rr, req)

    // Check the status code
    if rr.Code != http.StatusOK {
        t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
    }

    // Check the response body
    if rr.Body.String() != "Hello, World!" {
        t.Errorf("Expected response body %s, got %s", "Hello, World!", rr.Body.String())
    }
}

func TestRouter_404(t *testing.T) {
    // Create a new router
    router := Router()

    // Create a new request
    req, err := http.NewRequest("GET", "/not-found", nil)
    if err != nil {
        t.Fatal(err)
    }

    // Create a new response recorder
    rr := httptest.NewRecorder()

    // Serve the request
    router.ServeHTTP(rr, req)

    // Check the status code
    if rr.Code != http.StatusNotFound {
        t.Errorf("Expected status code %d, got %d", http.StatusNotFound, rr.Code)
    }

    // Check the response body
    if rr.Body.String() != "404 page not found" {
        t.Errorf("Expected response body %s, got %s", "404 page not found", rr.Body.String())
    }
}
