package main

import (
    "testing"
)

func TestRectangle_1b33bd8f37(t *testing.T) {
    tests := []struct {
        name     string
        input    int
        expected int
    }{
        {"positive", 10, 20},
        {"negative", -10, -20},
        {"zero", 0, 0},
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            actual := rectangle(test.input)
            if actual != test.expected {
                t.Errorf("expected %d, got %d", test.expected, actual)
            }
        })
    }
}
