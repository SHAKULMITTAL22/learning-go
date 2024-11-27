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


func Testbinarysearch(t *testing.T) {
	type testCase struct {
		name		string
		arr		[]int
		query		int
		expected	int
	}
	testCases := []testCase{{"Empty Email String", nil, 0, -1}, {"Maximum Length Email", generateMaxLengthArray(), 12345, 12345}}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			fmt.Fprintf(&buf, "Testing scenario: %s\n", tt.name)
			result := binarySearch(tt.arr, tt.query)
			if result != tt.expected {
				t.Errorf("Test %s failed: expected %d, got %d", tt.name, tt.expected, result)
			} else {
				t.Logf("Test %s succeeded: expected %d, got %d", tt.name, tt.expected, result)
			}
		})
	}
}

func generateMaxLengthArray() []int {
	maxLength := 100000
	arr := make([]int, maxLength)
	for i := 0; i < maxLength; i++ {
		arr[i] = i
	}
	return arr
}

func Testbinarysearch1(t *testing.T) {
	type testData struct {
		arr		[]int
		query		int
		expected	int
	}
	tests := []testData{{arr: []int{}, query: 0, expected: -1}, {arr: []int{1, 3, 5, 7, 9, 11, 13, 15}, query: 15, expected: 7}, {arr: []int{1, 3, 5, 7, 9}, query: 8, expected: -1}}
	for _, test := range tests {
		t.Run(fmt.Sprintf("arr: %v, query: %d", test.arr, test.query), func(t *testing.T) {
			var buffer bytes.Buffer
			old := os.Stdout
			os.Stdout = &buffer
			defer func() {
				os.Stdout = old
			}()
			result := binarySearch1(test.arr, test.query)
			t.Logf("Testing with input array: %v, query: %d", test.arr, test.query)
			if result != test.expected {
				t.Errorf("Expected %d, but got %d", test.expected, result)
			} else {
				t.Log("Test passed")
			}
			fmt.Fprintf(os.Stdout, "result: %d\n", result)
			if strings.TrimSpace(buffer.String()) != fmt.Sprintf("result: %d", result) {
				t.Errorf("Unexpected output to stdout")
			}
		})
	}
}

