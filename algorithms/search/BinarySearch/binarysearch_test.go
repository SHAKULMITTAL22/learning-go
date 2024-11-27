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


func Testbinarysearch568(t *testing.T) {
	tests := []struct {
		name		string
		arr		[]int
		query		int
		expectedIndex	int
	}{{name: "EmptyEmailString", arr: []int{}, query: 5, expectedIndex: -1}, {name: "MaximumLengthEmail", arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, query: 10, expectedIndex: 9}}
	oldStdout := os.Stdout
	readPipe, writePipe, _ := os.Pipe()
	defer func() {
		os.Stdout = oldStdout
	}()
	os.Stdout = writePipe
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout := &bytes.Buffer{}
			os.Stdout = stdout
			result := binarySearch(tt.arr, tt.query)
			if result == tt.expectedIndex {
				t.Logf("SUCCESS: %s - Expected index %d, got %d", tt.name, tt.expectedIndex, result)
			} else {
				t.Errorf("FAILURE: %s - Expected index %d, got %d", tt.name, tt.expectedIndex, result)
			}
			writePipe.Close()
			os.Stdout = oldStdout
			var output string
			fmt.Fscanf(readPipe, "%s", &output)
			if output != "" {
				fmt.Fprintf(oldStdout, "Captured output: %s\n", output)
			}
		})
	}
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

