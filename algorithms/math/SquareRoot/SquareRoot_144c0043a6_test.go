package SquareRoot

import (
    "math"
    "testing"
)

func TestSquareRoot_144c0043a6(t *testing.T) {
    tests := []struct {
        input    float64
        expected float64
    }{
        {4, 2},
        {9, 3},
        {16, 4},
        {25, 5},
        {36, 6},
        {49, 7},
        {64, 8},
        {81, 9},
        {100, 10},
        {0, 0},
        {-4, -1},
    }

    for _, test := range tests {
        actual := squareRoot(test.input)
        if actual != test.expected {
            t.Errorf("Expected %f, got %f", test.expected, actual)
        }
    }
}

func TestSquareRoot_EdgeCases(t *testing.T) {
    tests := []struct {
        input    float64
        expected float64
    }{
        {math.Inf(1), math.Inf(1)},
        {math.Inf(-1), math.Inf(-1)},
        {math.NaN(), math.NaN()},
    }

    for _, test := range tests {
        actual := squareRoot(test.input)
        if actual != test.expected {
            t.Errorf("Expected %f, got %f", test.expected, actual)
        }
    }
}
