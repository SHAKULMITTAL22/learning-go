package main

import (
    "fmt"
    "testing"
)

func TestMultiply_3477c70c83(t *testing.T) {
    tests := []struct {
        x, y int
        want int
    }{
        {1, 2, 2},
        {3, 4, 12},
        {5, 6, 30},
        {-1, -2, 2},
        {-3, -4, 12},
        {-5, -6, 30},
        {0, 1, 0},
        {1, 0, 0},
        {-1, 0, 0},
        {0, -1, 0},
    }
    for _, tt := range tests {
        t.Run(fmt.Sprintf("%d_%d", tt.x, tt.y), func(t *testing.T) {
            got := Multiply(tt.x, tt.y)
            if got != tt.want {
                t.Errorf("Multiply(%d, %d) = %d, want %d", tt.x, tt.y, got, tt.want)
            }
        })
    }
}
