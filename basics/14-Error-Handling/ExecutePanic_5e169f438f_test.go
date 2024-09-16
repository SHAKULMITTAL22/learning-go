package main

import (
    "fmt"
    "testing"
)

func executePanic() {
    // Defer function call
    defer recoveryFunction()
    // Throw panic
    panic("Panic")
    fmt.Println("End: executePanic")
}

func recoveryFunction() {
    if r := recover(); r != nil {
        fmt.Println("Recovered from panic:", r)
    }
}

func TestExecutePanic_5e169f438f(t *testing.T) {
    // Test case 1: Panic is recovered
    func() {
        defer func() {
            if r := recover(); r == nil {
                t.Error("Panic not recovered")
            }
        }()
        executePanic()
    }()

    // Test case 2: Panic is not recovered
    func() {
        defer func() {
            if r := recover(); r != nil {
                t.Error("Panic recovered")
            }
        }()
        executePanic()
    }()
}
