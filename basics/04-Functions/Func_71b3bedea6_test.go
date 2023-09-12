package main

import (
    "testing"
)

func TestFunc_71b3bedea6(t *testing.T) {
    tests := []struct {
        l, b int
        want int
    }{
        {1, 2, 2},
        {3, 4, 12},
        {5, 6, 30},
        {7, 8, 56},
        {9, 10, 90},
    }

    for _, tt := range tests {
        got := func(l int, b int) int {
            return l * b
        }(tt.l, tt.b)

        if got != tt.want {
            t.Errorf("Func(%d, %d) = %d, want %d", tt.l, tt.b, got, tt.want)
        }
    }
}
