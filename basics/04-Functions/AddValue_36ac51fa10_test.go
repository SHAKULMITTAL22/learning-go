package main

import (
    "testing"
)

func TestAddValue_36ac51fa10(t *testing.T) {
    x := 5
    y := "Hello"

    addValue(&x, &y)

    if x != 10 {
        t.Error("x is not 10")
    }

    if y != "Hello World!" {
        t.Error("y is not Hello World!")
    }
}

func TestAddValue_Error(t *testing.T) {
    x := 5
    y := "Hello"

    addValue(&x, nil)

    if x != 10 {
        t.Error("x is not 10")
    }

    if y != "Hello World!" {
        t.Error("y is not Hello World!")
    }
}
