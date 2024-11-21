package LevenshteinDistance

import (
	"testing"
)

func minTest(a, b uint16) uint16 {
	if a < b {
		return a
	}
	return b
}

func levenshteinDistanceTest(a, b string) int {
	if len(a) == 0 {
		return len(b)
	}

	if len(b) == 0 {
		return len(a)
	}

	if a == b {
		return 0
	}

	if len(a) > len(b) {
		a, b = b, a
	}
	lenS1 := len(a)
	lenS2 := len(b)

	x := make([]uint16, lenS1+1)

	for i := 1; i < len(x); i++ {
		x[i] = uint16(i)
	}

	for i := 1; i <= lenS2; i++ {
		prev := uint16(i)
		var current uint16
		for j := 1; j <= lenS1; j++ {
			if b[i-1] == a[j-1] {
				current = x[j-1]
			} else {
				current = minTest(minTest(x[j-1]+1, prev+1), x[j]+1)
			}
			x[j-1] = prev
			prev = current
		}
		x[lenS1] = prev
	}

	return int(x[lenS1])
}

func TestLevenshteinDistance_c58c51e350(t *testing.T) {
	tests := []struct {
		a        string
		b        string
		expected int
	}{
		{"kitten", "sitting", 3},
		{"", "test", 4},
		{"test", "", 4},
		{"test", "test", 0},
		{"longer", "lager", 2},
	}

	for _, test := range tests {
		result := levenshteinDistanceTest(test.a, test.b)
		if result != test.expected {
			t.Errorf("levenshteinDistanceTest(%q, %q) = %v; want %v", test.a, test.b, result, test.expected)
		} else {
			t.Logf("Success: levenshteinDistanceTest(%q, %q) = %v", test.a, test.b, result)
		}
	}
}
