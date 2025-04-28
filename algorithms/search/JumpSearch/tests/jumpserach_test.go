package JumpSearch

import "testing"

/*
ROOST_METHOD_HASH=jumpSearch_cf44201f2c
ROOST_METHOD_SIG_HASH=jumpSearch_45970a9b95
*/
func Testjumpsearch434(t *testing.T) {
	type testCase struct {
		name     string
		arr      []int
		query    int
		expected int
	}
	tests := []testCase{{name: "Scenario 1: Search for an Element in a Small Array", arr: []int{1, 3, 5, 7}, query: 5, expected: 2}, {name: "Scenario 2: Search for an Element at the Start of the Array", arr: []int{2, 4, 6, 8}, query: 2, expected: 0}, {name: "Scenario 3: Element Not Present in the Array", arr: []int{1, 4, 6, 9}, query: 3, expected: -1}, {name: "Scenario 4: Search for an Element in a Large Array", arr: generateLargeArray(), query: 5023, expected: 5023}, {name: "Scenario 5: Search in an Array with Duplicate Elements", arr: []int{1, 3, 5, 5, 7}, query: 5, expected: 2}, {name: "Scenario 6: Empty Array Check", arr: []int{}, query: 10, expected: -1}, {name: "Scenario 7: Element at the End of the Array", arr: []int{1, 3, 5, 9, 11}, query: 11, expected: 4}}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := jumpSearch(test.arr, test.query)
			if actual != test.expected {
				t.Errorf("Test %s failed: Expected index %d, but got %d", test.name, test.expected, actual)
			} else {
				t.Logf("Test %s succeeded: Correct index %d was returned", test.name, test.expected)
			}
		})
	}
}

func generateLargeArray() []int {
	arr := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		arr[i] = i
	}
	return arr
}
