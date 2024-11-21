package EuclideanAlgorithm

import (
    "testing"
)

func TestGCD_ab1c91475d(t *testing.T) {
    tests := []struct {
        a, b, expected int
    }{
        {12, 18, 6},
        {-12, -18, 6},
        {12, 0, 12},
        {0, 0, 0},
        {1234567890, 9876543210, 10},
        {1, 2, 1},
        {7, 11, 1},
        {12, 15, 3},
        {12, 18, 6},
        {1234567890, 9876543210, 10},
    }
    for _, test := range tests {
        actual := GCD(test.a, test.b)
        if actual != test.expected {
            t.Errorf("GCD(%d, %d) = %d; expected %d", test.a, test.b, actual, test.expected)
        }
    }
}
