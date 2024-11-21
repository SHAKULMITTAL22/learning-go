package main

import (
    "testing"
)

type Car struct{}

func (c Car) Alert() string {
    return "Hup! Hup!"
}

func TestAlert(t *testing.T) {
    // Test case 1: Alert returns the expected string
    expected := "Hup! Hup!"
    actual := Car{}.Alert()
    if actual != expected {
        t.Errorf("Expected %q, got %q", expected, actual)
    }

    // Test case 2: Alert returns a different string
    expected = "Beep! Beep!"
    actual = Car{}.Alert()
    if actual == expected {
        t.Errorf("Expected %q, got %q", expected, actual)
    }
}
