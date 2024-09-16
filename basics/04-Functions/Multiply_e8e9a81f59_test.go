package main

import "testing"

func TestMultiply_e8e9a81f59(t *testing.T) {
	tests := []struct {
		x, y int
		want int
	}{
		{1, 2, 2},
		{3, 4, 12},
		{-1, -2, 2},
		{0, 1, 0},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := multiply(tt.x, tt.y)
			if got != tt.want {
				t.Errorf("multiply(%d, %d) = %d, want %d", tt.x, tt.y, got, tt.want)
			}
		})
	}
}
