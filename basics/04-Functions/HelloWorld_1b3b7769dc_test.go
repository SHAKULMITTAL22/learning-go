package main

import (
    "testing"
)

func TestHelloWorld_1b3b7769dc(t *testing.T) {
    // Test that the helloWorld function prints the correct message.
    expected := "Hello World!"
    actual := helloWorld()
    if actual != expected {
        t.Errorf("Expected %q, got %q", expected, actual)
    }
}
