func Testbinarysearch768(t *testing.T) {
	type testCase struct {
		inputArray	[]int
		query		int
		expectedIndex	int
		description	string
	}
	testCases := []testCase{{[]int{}, 10, -1, "Scenario 1: Validate Empty Email String - Empty array"}, {make([]int, 256), 0, -1, "Scenario 2: Validate Maximum Length Email - Large array with no matches"}}
	var buffer bytes.Buffer
	writer := bufio.NewWriter(&buffer)
	saveStdout := os.Stdout
	r, w, _ := os.Pipe()
	defer r.Close()
	defer w.Close()
	for _, tc := range testCases {
		t.Log(tc.description)
		os.Stdout = w
		go func() {
			result := binarySearch(tc.inputArray, tc.query)
			if result != tc.expectedIndex {
				fmt.Fprintf(writer, "Failed: for input %v and query %d, expected %d, got %d\n", tc.inputArray, tc.query, tc.expectedIndex, result)
				t.Errorf("Expected %d but got %d", tc.expectedIndex, result)
			} else {
				fmt.Fprintf(writer, "Passed: %s\n", tc.description)
			}
			writer.Flush()
		}()
		_, err := fmt.Fscanf(r, "%s\n", &buffer)
		if err != nil {
			t.Errorf("Error capturing output: %s", err)
		}
		fmt.Print(buffer.String())
	}
	os.Stdout = saveStdout
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

