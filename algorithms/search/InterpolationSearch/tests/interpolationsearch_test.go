package InterpolationSearch

import "testing"

/*
ROOST_METHOD_HASH=interpolationSearch_d0bdb80f3c
ROOST_METHOD_SIG_HASH=interpolationSearch_23d9bb25dd
*/
func Testinterpolationsearch424(t *testing.T) {
	testCases := []struct {
		name        string
		arr         []int
		query       int
		expected    int
		description string
	}{{name: "Present Value", arr: []int{1, 3, 5, 7, 9}, query: 5, expected: 2, description: "Search value exists in the array at index 2."}, {name: "Absent Value", arr: []int{2, 4, 6, 8, 10}, query: 5, expected: -1, description: "Search value does not exist in the array."}, {name: "Empty Array", arr: []int{}, query: 1, expected: -1, description: "Search in an empty array should return -1."}, {name: "Boundary Start", arr: []int{10, 20, 30, 40, 50}, query: 10, expected: 0, description: "Search value is at the start of the array."}, {name: "Boundary End", arr: []int{10, 20, 30, 40, 50}, query: 50, expected: 4, description: "Search value is at the end of the array."}, {name: "Identical Elements", arr: []int{7, 7, 7, 7, 7}, query: 7, expected: 0, description: "Array with identical elements, any index between 0-4 is valid."}, {name: "Large Array", arr: createLargeArray(10000), query: 9876, expected: 9875, description: "Large array search for performance."}, {name: "Min Integer", arr: []int{-2147483648, 0, 2147483647}, query: -2147483648, expected: 0, description: "Search for minimum integer value."}, {name: "Max Integer", arr: []int{-2147483648, 0, 2147483647}, query: 2147483647, expected: 2, description: "Search for maximum integer value."}}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.description)
			index := interpolationSearch(tc.arr, tc.query)
			if index != tc.expected {
				t.Errorf("failed %s: expected index %d but got %d", tc.name, tc.expected, index)
			} else {
				t.Logf("passed %s: successfully found index %d", tc.name, index)
			}
		})
	}
}

func createLargeArray(size int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = i + 1
	}
	return arr
}
