package PrimalityTest

import (
	"testing"
)

func TestMod_10abfc4edf(t *testing.T) {
	testCases := []struct {
		a, b, expected int
	}{
		{10, 3, 1},
		{-10, 3, -1},
		{10, -3, 1},
		{-10, -3, -1},
		{10, 1, 0},
		{-10, 1, 0},
		{10, 0, 0}, // error case
		{-10, 0, 0}, // error case
	}

	for _, tc := range testCases {
		actual := mod(tc.a, tc.b)
		if actual != tc.expected {
			t.Errorf("Expected %d, got %d", tc.expected, actual)
		}
	}
}

func mod(a, b int) int {
	if b == 0 {
		return 0
	}
	return a % b
}
