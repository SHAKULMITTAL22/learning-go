package main

import (
	"errors"
	"fmt"
	"math"
	"testing"
)

func TestSqrt_acfd7a2e9b(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
		err      error
	}{
		{4, 2, nil},
		{9, 3, nil},
		{-1, 0, errors.New("Math: Not possible to calculate the square root of negative number")},
		{0, 0, nil},
	}

	for _, test := range tests {
		actual, err := Sqrt(test.input)
		if err != test.err {
			t.Errorf("Expected error %v, got %v", test.err, err)
		}
		if math.Abs(actual-test.expected) > 0.0001 {
			t.Errorf("Expected %v, got %v", test.expected, actual)
		}
	}
}
