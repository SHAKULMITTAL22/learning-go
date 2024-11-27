package BinarySearch

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

import (
	"testing"
)


func Testinterpolationsearch658(t *testing.T) {
	type testCase struct {
		arr		[]int
		query		int
		expected	int
		description	string
	}
	testCases := []testCase{{arr: []int{}, query: 5, expected: -1, description: "Empty array should return -1 for any query"}, {arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, query: 10, expected: 9, description: "Query at the maximum length of a non-empty sorted array"}, {arr: []int{10, 20, 30, 40, 50}, query: 35, expected: -1, description: "Query for a non-existing value within the array limits"}, {arr: []int{5, 10, 15, 20, 25}, query: 5, expected: 0, description: "Query for the minimum item in the array"}, {arr: []int{5, 10, 15, 20, 25}, query: 25, expected: 4, description: "Query for the maximum item in the array"}, {arr: []int{10, 20, 30}, query: 40, expected: -1, description: "Query is greater than the maximum item in array"}, {arr: []int{10, 20, 30}, query: 5, expected: -1, description: "Query is less than the minimum item in array"}}
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := interpolationSearch(tc.arr, tc.query)
			if actual != tc.expected {
				t.Errorf("Failed: %s. Expected %d, Got %d", tc.description, tc.expected, actual)
			} else {
				t.Logf("Passed: %s", tc.description)
			}
		})
	}
}

