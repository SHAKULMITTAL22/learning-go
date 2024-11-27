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


func Testlinearsearch598(t *testing.T) {
	tests := []struct {
		description	string
		arr		[]int
		query		int
		expected	int
	}{{description: "Empty array", arr: []int{}, query: 5, expected: -1}, {description: "Single element array with match", arr: []int{5}, query: 5, expected: 0}, {description: "Single element array without match", arr: []int{7}, query: 5, expected: -1}, {description: "Array with multiple elements, query present", arr: []int{1, 2, 3, 5, 7}, query: 5, expected: 3}, {description: "Array with multiple elements, query absent", arr: []int{1, 2, 3, 4, 6, 7}, query: 5, expected: -1}, {description: "Query at the beginning", arr: []int{5, 2, 3}, query: 5, expected: 0}, {description: "Query at the end", arr: []int{1, 2, 5}, query: 5, expected: 2}, {description: "High volume data with match at middle", arr: make([]int, 1000), query: 500, expected: 500}, {description: "High volume data with no match", arr: make([]int, 1000), query: 5001, expected: -1}}
	var buffer bytes.Buffer
	old := os.Stdout
	os.Stdout = &buffer
	defer func() {
		os.Stdout = old
	}()
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			t.Log("Running test case:", tt.description)
			if len(tt.arr) > 0 {
				for i := range tt.arr {
					tt.arr[i] = i
				}
			}
			result := linearSearch(tt.arr, tt.query)
			fmt.Fprintf(os.Stdout, "Expected: %d, Got: %d\n", tt.expected, result)
			if result != tt.expected {
				t.Errorf("Failure in %s: expected %d, got %d", tt.description, tt.expected, result)
			} else {
				t.Logf("Success in %s: expected and got %d", tt.description, result)
			}
		})
	}
	fmt.Fscanf(&buffer, "Format output from buffer")
}

