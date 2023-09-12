package main

import "testing"

func TestAdd_1444c4af66(t *testing.T) {
	tests := []struct {
		x, y int
		want int
	}{
		{1, 2, 3},
		{3, 4, 7},
		{-1, -2, -3},
		{0, 0, 0},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d+%d", tt.x, tt.y), func(t *testing.T) {
			got := add(tt.x, tt.y)
			if got != tt.want {
				t.Errorf("add(%d, %d) = %d, want %d", tt.x, tt.y, got, tt.want)
			}
		})
	}
}
