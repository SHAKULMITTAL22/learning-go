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


func Testlinearsearch614(t *testing.T) {
	testCases := []struct {
		description	string
		array		[]int
		query		int
		expected	int
	}{{description: "Validate Empty Email String", array: []int{}, query: 50, expected: -1}, {description: "Validate Maximum Length Email", array: generateLargeArray(10000), query: 9999, expected: 9999}}
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			w.Close()
			out := &bytes.Buffer{}
			fmt.Fprintf(out, "Running test case: %s\n", tc.description)
			result := linearSearch(tc.array, tc.query)
			fmt.Fprintf(out, "Expected: %d, Got: %d\n", tc.expected, result)
			if result != tc.expected {
				t.Errorf("Test %s failed. Expected %d, got %d", tc.description, tc.expected, result)
			} else {
				t.Logf("Test %s succeeded.", tc.description)
			}
			output := out.String()
			if !strings.Contains(output, fmt.Sprintf("Expected: %d, Got: %d", tc.expected, result)) {
				t.Errorf("Output did not match expectations. Output: %s", output)
			}
		})
	}
	os.Stdout = old
}

func generateLargeArray(size int) []int {
	largeArray := make([]int, size)
	for i := range largeArray {
		largeArray[i] = i
	}
	return largeArray
}

