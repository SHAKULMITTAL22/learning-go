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


func Testinterpolationsearch466(t *testing.T) {
	tests := []struct {
		name		string
		arr		[]int
		query		int
		expected	int
	}{{name: "Validate Empty Email String", arr: []int{}, query: 10, expected: -1}, {name: "Validate Maximum Length Email", arr: createLargeArray(1000000), query: 1000000 - 1, expected: 1000000 - 1}}
	origStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Fprintf(os.Stdout, "Starting test: %s\n", tt.name)
			actual := interpolationSearch(tt.arr, tt.query)
			w.Close()
			var sb strings.Builder
			fmt.Fscanf(r, "%s", &sb)
			os.Stdout = origStdout
			t.Logf("Test %s ran, logging captured output:", tt.name)
			t.Log(sb.String())
			if actual != tt.expected {
				t.Errorf("Test %s failed: expected %d, got %d", tt.name, tt.expected, actual)
			} else {
				t.Logf("Test %s succeeded: expected %d, got %d", tt.name, tt.expected, actual)
			}
		})
	}
}

func createLargeArray(length int) []int {
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = i
	}
	return arr
}

