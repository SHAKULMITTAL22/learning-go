package main

import (
	"testing"
	"fmt"
	"bytes"
	"os"
)

var out *os.File

func sayHelloTo(name string) {
	fmt.Fprintln(out, "Hello", name)
}

func TestSayHelloTo_4a6fd0c56c(t *testing.T) {
    // Test case for correct name
    t.Run("Correct Name", func(t *testing.T) {
        var buf bytes.Buffer
        out = os.NewFile(0, "/dev/null")
        out.SetOutput(&buf)
        sayHelloTo("John")

        got := buf.String()
        want := "Hello John\n"

        if got != want {
            t.Errorf("got %q, but wanted %q", got, want)
        } else {
			t.Log("Success: The correct greeting is generated.")
		}
    })

    // Test case for empty name
    t.Run("Empty Name", func(t *testing.T) {
        var buf bytes.Buffer
        out = os.NewFile(0, "/dev/null")
        out.SetOutput(&buf)
        sayHelloTo("")

        got := buf.String()
        want := "Hello \n"

        if got != want {
            t.Errorf("got %q, but wanted %q", got, want)
        } else {
			t.Log("Success: The correct greeting is generated for blank name.")
		}
    })
    
    out = os.Stdout
}
