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


func Testjumpsearch347(t *testing.T) {
	type testCase struct {
		description	string
		arr		[]int
		query		int
		expected	int
	}
	testCases := []testCase{{description: "Validate searching in an empty array", arr: []int{}, query: 5, expected: -1}, {description: "Validate finding a number in a maximum length array", arr: generateMaxLengthArray(1000), query: 999, expected: 999}, {description: "Validate query number is not in the array", arr: []int{1, 2, 3, 4, 5}, query: 10, expected: -1}, {description: "Validate query number at the start of the array", arr: []int{5, 6, 7, 8, 9}, query: 5, expected: 0}, {description: "Validate query number at the end of the array", arr: []int{1, 2, 3, 4, 9}, query: 9, expected: 4}}
	var buf bytes.Buffer
	stdout := os.Stdout
	os.Stdout = &buf
	defer func() {
		os.Stdout = stdout
	}()
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			t.Log(fmt.Sprintf("Running scenario: %s", tc.description))
			result := jumpSearch(tc.arr, tc.query)
			if result != tc.expected {
				t.Errorf("Failed %s: expected %d, got %d", tc.description, tc.expected, result)
			} else {
				t.Log(fmt.Sprintf("Passed %s: got expected result %d", tc.description, tc.expected))
			}
		})
	}
	if buf.Len() > 0 {
		t.Log("Captured output from stdout: ", buf.String())
	}
}

func generateMaxLengthArray(length int) []int {
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = i
	}
	return arr
}

