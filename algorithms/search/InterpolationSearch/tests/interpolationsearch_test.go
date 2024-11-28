package InterpolationSearch

import "testing"

/*
ROOST_METHOD_HASH=interpolationSearch_d0bdb80f3c
ROOST_METHOD_SIG_HASH=interpolationSearch_23d9bb25dd


 */
func Testinterpolationsearch769(t *testing.T) {
	tests := []struct {
		name	string
		arr	[]int
		query	int
		expect	int
	}{{name: "Scenario 1: Search in a Sorted Array with Present Value", arr: []int{1, 3, 5, 7, 9}, query: 5, expect: 2}, {name: "Scenario 2: Search for a Value Not in the Array", arr: []int{2, 4, 6, 8, 10}, query: 5, expect: -1}, {name: "Scenario 3: Search in an Empty Array", arr: []int{}, query: 1, expect: -1}, {name: "Scenario 4: Search at the Array Boundaries - Start", arr: []int{10, 20, 30, 40, 50}, query: 10, expect: 0}, {name: "Scenario 4: Search at the Array Boundaries - End", arr: []int{10, 20, 30, 40, 50}, query: 50, expect: 4}, {name: "Scenario 5: Search with Identical Elements", arr: []int{7, 7, 7, 7, 7}, query: 7, expect: 0}, {name: "Scenario 6: Large Array Search Performance", arr: createLargeArray(1, 10000), query: 9876, expect: 9875}, {name: "Scenario 7: Search for Minimum and Maximum Integer Values - Min", arr: []int{-2147483648, 0, 2147483647}, query: -2147483648, expect: 0}, {name: "Scenario 7: Search for Minimum and Maximum Integer Values - Max", arr: []int{-2147483648, 0, 2147483647}, query: 2147483647, expect: 2}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := interpolationSearch(tt.arr, tt.query)
			if got != tt.expect {
				t.Errorf("%s failed: expected %d, got %d", tt.name, tt.expect, got)
			} else {
				t.Logf("%s succeeded: found %d at index %d", tt.name, tt.query, got)
			}
		})
	}
}

/*
ROOST_METHOD_HASH=interpolationSearch_d0bdb80f3c
ROOST_METHOD_SIG_HASH=interpolationSearch_23d9bb25dd


 */
func createLargeArray(start, end int) []int {
	arr := make([]int, end-start+1)
	for i := range arr {
		arr[i] = start + i
	}
	return arr
}

