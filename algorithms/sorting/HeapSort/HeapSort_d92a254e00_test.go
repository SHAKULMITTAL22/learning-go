package HeapSort

import (
	"reflect"
	"testing"
)

// TestHeapSort_d92a254e00 is a test function for the HeapSort method
func TestHeapSort_d92a254e00(t *testing.T) {
	heap := &Heap{}

	// Test case 1: Normal scenario
	input1 := []int{3, 2, 1, 5, 6, 4}
	expectedOutput1 := []int{1, 2, 3, 4, 5, 6}
	output1 := heap.HeapSort(input1)
	if !reflect.DeepEqual(output1, expectedOutput1) {
		t.Errorf("HeapSort test failed! Input: %v, Output: %v, ExpectedOutput: %v", input1, output1, expectedOutput1)
	} else {
		t.Logf("HeapSort test passed! Input: %v, Output: %v", input1, output1)
	}

	// Test case 2: Edge case with an empty array
	input2 := []int{}
	expectedOutput2 := []int{}
	output2 := heap.HeapSort(input2)
	if !reflect.DeepEqual(output2, expectedOutput2) {
		t.Errorf("HeapSort test failed! Input: %v, Output: %v, ExpectedOutput: %v", input2, output2, expectedOutput2)
	} else {
		t.Logf("HeapSort test passed! Input: %v, Output: %v", input2, output2)
	}
}
