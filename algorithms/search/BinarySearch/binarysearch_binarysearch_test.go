package BinarySearch

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestbinarySearch(t *testing.T) {

	originalStdout := os.Stdout
	defer func() { os.Stdout = originalStdout }()
	r, w, _ := os.Pipe()
	os.Stdout = w

	tests := []struct {
		name     string
		arr      []int
		query    int
		expected int
		log      string
	}{
		{
			name:     "Empty array",
			arr:      []int{},
			query:    5,
			expected: -1,
			log:      "Scenario 1: Validate Empty Email String",
		},
		{
			name:     "Maximum integer search",
			arr:      []int{1, 2, 3, 4, 5},
			query:    2147483647,
			expected: -1,
			log:      "Scenario 2: Validate Maximum Length Email",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			t.Log(test.log)

			result := binarySearch(test.arr, test.query)

			fmt.Fprintf(w, "Test %s: Got %d, Expected %d\n", test.name, result, test.expected)

			if result != test.expected {
				t.Errorf("Failed %s: Expected %d, got %d", test.name, test.expected, result)
			}

			w.Close()
			var buf strings.Builder
			fmt.Fscan(r, &buf)
			expectedOutput := fmt.Sprintf("Test %s: Got %d, Expected %d\n", test.name, result, test.expected)
			if buf.String() != expectedOutput {
				t.Errorf("Unexpected output: got %q expected %q", buf.String(), expectedOutput)
			}

		})
	}

}
