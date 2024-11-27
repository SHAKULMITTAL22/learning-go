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
	"fmt"
	"os"
	"strings"
	"testing"
)

import (
	"strings"
	"testing"
)


func Testjumpsearch472(t *testing.T) {
	type testCase struct {
		name		string
		arr		[]int
		query		int
		expected	int
	}
	testCases := []testCase{{name: "Test_Empty_Array", arr: []int{}, query: 1, expected: -1}, {name: "Test_Max_Length_Email", arr: generateLargeArray(), query: 9999, expected: 9999}}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Logf("Running test case: %s", tc.name)
			result := jumpSearch(tc.arr, tc.query)
			if result != tc.expected {
				t.Errorf("Expected %d, but got %d", tc.expected, result)
			} else {
				t.Logf("Success case for %q with result: %d", tc.name, result)
			}
		})
	}
}

func generateLargeArray() []int {
	maxLength := 10000
	largeArray := make([]int, maxLength)
	for i := 0; i < maxLength; i++ {
		largeArray[i] = i
	}
	return largeArray
}

