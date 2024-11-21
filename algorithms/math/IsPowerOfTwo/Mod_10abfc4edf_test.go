package IsPowerOfTwo

import (
	"testing"
)

func TestMod_10abfc4edf(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{10, 3, 1},
		{-10, 3, 2},
		{10, -3, 1},
		{-10, -3, 2},
		{0, 3, 0},
		{0, -3, 0},
		{-1, 3, 2},
		{-1, -3, 2},
	}
	for _, tt := range tests {
		got := mod(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("mod(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}
