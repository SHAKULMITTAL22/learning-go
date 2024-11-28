package InterpolationSearch

import "testing"

/*
ROOST_METHOD_HASH=interpolationSearch_d0bdb80f3c
ROOST_METHOD_SIG_HASH=interpolationSearch_23d9bb25dd


 */
func Testinterpolationsearch592(t *testing.T) {
	tests := []struct {
		name   string
		arr    []int
		query  int
		expect int
	}{

		{"ExistingElement", []int{10, 20, 30, 40, 50}, 30, 2},

		{"NonExistingElement", []int{10, 20, 30, 40, 50}, 35, -1},

		{"EmptyArray", []int{}, 30, -1},

		{"IdenticalElementsPresent", []int{10, 10, 10, 10, 10}, 10, 0},

		{"SingleElementMatching", []int{30}, 30, 0},

		{"QueryLargerThanMax", []int{10, 20, 30, 40}, 50, -1},

		{"QuerySmallerThanMin", []int{10, 20, 30, 40}, 5, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log("Starting test for:", tt.name)
			result := interpolationSearch(tt.arr, tt.query)
			if result != tt.expect {
				t.Errorf("Failed %s: expected %d, got %d", tt.name, tt.expect, result)
			} else {
				t.Logf("Success %s: expected %d, got %d", tt.name, tt.expect, result)
			}
		})
	}
}

