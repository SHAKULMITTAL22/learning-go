package Factorial

import (
    "testing"
)

func TestFactorial_c8838e8b35(t *testing.T) {
    tests := []struct {
        name string
        num  int
        want int
    }{
        {"positive", 5, 120},
        {"negative", -5, 1},
        {"zero", 0, 1},
        {"one", 1, 1},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := Factorial(tt.num); got != tt.want {
                t.Errorf("Factorial(%d) = %d, want %d", tt.num, got, tt.want)
            }
        })
    }
}
