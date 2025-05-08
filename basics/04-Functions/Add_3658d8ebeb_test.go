package main

import "testing"

func TestAdd_3658d8ebeb(t *testing.T) {
	tests := []struct {
		x, y int
		want int
	}{
		{1, 2, 3},
		{3, 4, 7},
		{-1, -2, -3},
		{0, 0, 0},
		{-100, 100, 0},
		{1000, -1000, 0},
	}
	for _, tt := range tests {
		got := add(tt.x, tt.y)
		if got != tt.want {
			t.Errorf("add(%d, %d) = %d, want %d", tt.x, tt.y, got, tt.want)
		}
	}
}
