func Testjumpsearch856(t *testing.T) {
	testCases := []struct {
		name		string
		array		[]int
		query		int
		expected	int
	}{{name: "EmptyArray", array: []int{}, query: 5, expected: -1}, {name: "SingleElementNotFound", array: []int{3}, query: 5, expected: -1}, {name: "SingleElementFound", array: []int{5}, query: 5, expected: 0}, {name: "SmallArrayQueryFound", array: []int{1, 2, 3, 4, 5}, query: 3, expected: 2}, {name: "SmallArrayQueryNotFound", array: []int{1, 2, 4, 5}, query: 3, expected: -1}, {name: "MaximumLengthEmail", array: generateArray(100), query: 99, expected: 99}, {name: "LargeArray", array: generateArray(1000), query: 995, expected: 995}}
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				os.Stdout = oldStdout
			}()
			result := jumpSearch(tt.array, tt.query)
			if result != tt.expected {
				t.Errorf("jumpSearch(%v, %d) = %d; expected %d", tt.array, tt.query, result, tt.expected)
			}
			t.Logf("Test %s: success", tt.name)
		})
	}
	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		fmt.Fscanf(r, "%s", &buf)
		outC <- buf.String()
	}()
	w.Close()
	out := <-outC
	if !strings.Contains(out, "success") {
		t.Errorf("Unexpected std output: %s", out)
	}
}

func generateArray(n int) []int {
	array := make([]int, n)
	for i := 0; i < n; i++ {
		array[i] = i
	}
	return array
}

