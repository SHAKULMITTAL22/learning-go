package RadixSort

import (
	"reflect"
	"testing"
)

func countSort(arr []int, n int, exp int) {
	// Function body here
}

// TestCountSort_de82b0c79b is a test function for countSort
func TestCountSort_de82b0c79b(t *testing.T) {
	// Test case 1: Normal scenario
	arr1 := []int{170, 45, 75, 90, 802, 24, 2, 66}
	exp1 := 1
	expected1 := []int{170, 90, 802, 2, 24, 45, 75, 66}
	countSort(arr1, len(arr1), exp1)
	if !reflect.DeepEqual(arr1, expected1) {
		t.Errorf("Test case 1 failed, expected: %v, got: %v", expected1, arr1)
	} else {
		t.Log("Test case 1 passed")
	}

	// Test case 2: Edge case of empty array
	arr2 := []int{}
	exp2 := 1
	expected2 := []int{}
	countSort(arr2, len(arr2), exp2)
	if !reflect.DeepEqual(arr2, expected2) {
		t.Errorf("Test case 2 failed, expected: %v, got: %v", expected2, arr2)
	} else {
		t.Log("Test case 2 passed")
	}

	// Test case 3: Array with negative numbers
	arr3 := []int{-1, -2, -3, -4, -5}
	exp3 := 1
	expected3 := []int{-1, -2, -3, -4, -5} // countSort does not sort negative numbers
	countSort(arr3, len(arr3), exp3)
	if !reflect.DeepEqual(arr3, expected3) {
		t.Errorf("Test case 3 failed, expected: %v, got: %v", expected3, arr3)
	} else {
		t.Log("Test case 3 passed")
	}
}
