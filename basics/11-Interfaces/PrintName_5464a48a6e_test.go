package main

import (
    "fmt"
    "testing"
)

type Usr int

func (usr Usr) PrintName(name string) {
    fmt.Println("User Id:\t", usr)
    fmt.Println("User Name:\t", name)
}

func TestPrintName_5464a48a6e(t *testing.T) {
    // Test case 1: Valid user ID and name
    usr := Usr(123)
    name := "John Doe"
    usr.PrintName(name)
    expectedOutput := "User Id:\t 123\nUser Name:\t John Doe\n"
    if gotOutput := fmt.Sprintf("%v", usr); gotOutput != expectedOutput {
        t.Errorf("Expected output:\n%s\nGot output:\n%s", expectedOutput, gotOutput)
    }

    // Test case 2: Invalid user ID
    usr = Usr(-1)
    name = "Jane Doe"
    usr.PrintName(name)
    expectedOutput = "User Id:\t -1\nUser Name:\t Jane Doe\n"
    if gotOutput := fmt.Sprintf("%v", usr); gotOutput != expectedOutput {
        t.Errorf("Expected output:\n%s\nGot output:\n%s", expectedOutput, gotOutput)
    }

    // Test case 3: Empty user name
    usr = Usr(123)
    name = ""
    usr.PrintName(name)
    expectedOutput = "User Id:\t 123\nUser Name:\t \n"
    if gotOutput := fmt.Sprintf("%v", usr); gotOutput != expectedOutput {
        t.Errorf("Expected output:\n%s\nGot output:\n%s", expectedOutput, gotOutput)
    }
}
