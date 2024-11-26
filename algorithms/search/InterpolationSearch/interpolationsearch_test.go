func Testinterpolationsearch759(t *testing.T) {
	type test struct {
		name		string
		arr		[]int
		query		int
		expected	int
	}
	tests := []test{{name: "Scenario 1: Validate Empty Email String (Empty Array)", arr: []int{}, query: 5, expected: -1}, {name: "Scenario 2: Validate Maximum Length Email (Maximum Length Array)", arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, query: 5, expected: 4}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w
			result := interpolationSearch(tt.arr, tt.query)
			if result != tt.expected {
				t.Errorf("Test Failed: %s, Expected: %d, Got: %d", tt.name, tt.expected, result)
			} else {
				fmt.Fprintf(w, "Test Passed: %s, Expected: %d, Got: %d", tt.name, tt.expected, result)
			}
			w.Close()
			var buf bytes.Buffer
			fmt.Fscanf(r, "%s", &buf)
			if buf.String() != "" {
				t.Log(buf.String())
			}
			os.Stdout = oldStdout
		})
	}
}

