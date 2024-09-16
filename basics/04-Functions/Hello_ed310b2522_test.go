package main

import (
    "fmt"
    "testing"
)

func hello(x string) {
    fmt.Printf("Hello %s\n", x)
}

func TestHello_ed310b2522(t *testing.T) {
    // Test case 1: Valid input
    hello("John")
    if output != "Hello John\n" {
        t.Error("Expected output: Hello John\n, but got:", output)
    }

    // Test case 2: Invalid input
    hello("")
    if output != "Hello \n" {
        t.Error("Expected output: Hello \n, but got:", output)
    }
}
