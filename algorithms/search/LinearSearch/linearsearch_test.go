func Testlinearsearch512(t *testing.T) {
	type testScenario struct {
		name		string
		arr		[]int
		query		int
		expected	int
		stdout		string
		captureOs	bool
	}
	tests := []testScenario{{name: "Empty Array", arr: []int{}, query: 1, expected: -1, stdout: "Query 1 not found in empty array\n", captureOs: true}, {name: "Element Present", arr: []int{10, 20, 30, 40, 50}, query: 30, expected: 2, stdout: "Found element 30 at index 2\n", captureOs: false}, {name: "Element Absent", arr: []int{10, 20, 30, 40, 50}, query: 60, expected: -1, stdout: "Query 60 not found. Returning -1\n", captureOs: true}, {name: "Multiple Occurrences", arr: []int{1, 2, 3, 2, 1}, query: 2, expected: 1, stdout: "Found element 2 at first occurrence index 1\n", captureOs: false}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var output strings.Builder
			if tt.captureOs {
				fmt.Fprintf(&output, "Test: %s\n", tt.stdout)
			}
			result := linearSearch(tt.arr, tt.query)
			if result != tt.expected {
				t.Errorf("Test %s failed: expected %v, got %v", tt.name, tt.expected, result)
			} else {
				t.Logf("Test %s successful: Got %v as expected", tt.name, result)
			}
			if tt.captureOs && output.String() != tt.stdout {
				t.Errorf("Expected stdout: %v, Got: %v", tt.stdout, output.String())
			}
		})
	}
}

