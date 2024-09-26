package RadixSort

import (
	"testing"
	"reflect"
)

// TestRadixsort_c800f73b57 is a test function for radixsort
func TestRadixsort_c800f73b57(t *testing.T) {
	// Test case 1: Normal scenario
	arr1 := []int{170, 45, 75, 90, 802, 24, 2, 66}
	radixsort(arr1, len(arr1))
	expected1 := []int{2, 24, 45, 66, 75, 90, 170, 802}
	if !reflect.DeepEqual(arr1, expected1) {
		t.Errorf("Test case 1 failed: got %v, want %v", arr1, expected1)
	} else {
		t.Logf("Test case 1 succeeded")
	}

	// Test case 2: Edge case with single element
	arr2 := []int{5}
	radixsort(arr2, len(arr2))
	expected2 := []int{5}
	if !reflect.DeepEqual(arr2, expected2) {
		t.Errorf("Test case 2 failed: got %v, want %v", arr2, expected2)
	} else {
		t.Logf("Test case 2 succeeded")
	}

	// Test case 3: Edge case with negative numbers
	arr3 := []int{-5, -10, -3, -1}
	radixsort(arr3, len(arr3))
	expected3 := []int{-10, -5, -3, -1}
	if !reflect.DeepEqual(arr3, expected3) {
		t.Errorf("Test case 3 failed: got %v, want %v", arr3, expected3)
	} else {
		t.Logf("Test case 3 succeeded")
	}

	// Test case 4: Edge case with all zeros
	arr4 := []int{0, 0, 0, 0}
	radixsort(arr4, len(arr4))
	expected4 := []int{0, 0, 0, 0}
	if !reflect.DeepEqual(arr4, expected4) {
		t.Errorf("Test case 4 failed: got %v, want %v", arr4, expected4)
	} else {
		t.Logf("Test case 4 succeeded")
	}
}
