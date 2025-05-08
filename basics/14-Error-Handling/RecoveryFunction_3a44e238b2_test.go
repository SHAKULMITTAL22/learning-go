package main

import (
    "fmt"
    "testing"
)

func recoveryFunction() {
    // Recover from panic and print erro message
    if recoveryMessage := recover(); recoveryMessage != nil {
        fmt.Println(recoveryMessage)
    }
    fmt.Println("End: recoveryFunction")
}

func TestRecoveryFunction_3a44e238b2(t *testing.T) {
    // Test case 1: Panic occurs
    func() {
        defer recoveryFunction()
        panic("Panic occurred")
    }()

    // Expected output:
    // Panic occurred
    // End: recoveryFunction

    // Test case 2: No panic occurs
    func() {
        defer recoveryFunction()
        fmt.Println("No panic occurred")
    }()

    // Expected output:
    // No panic occurred
    // End: recoveryFunction
}
