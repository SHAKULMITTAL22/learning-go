package RadixSort

import (
	"testing"
	"reflect"
)

func TestRadixsort_c800f73b57(t *testing.T) {
	// Test case 1: Normal scenario
	arr1 := []int{170, 45, 75, 90, 802, 24, 2, 66}
	n1 := len(arr1)
	expected1 := []int{2, 24, 45, 66, 75, 90, 170, 802}
	radixsort(arr1, n1)
	if !reflect.DeepEqual(arr1, expected1) {
		t.Errorf("Test case 1 failed; got %v, want %v", arr1, expected1)
	} else {
		t.Logf("Test case 1 succeeded")
	}

	// Test case 2: Edge case of having all elements as the same
	arr2 := []int{5, 5, 5, 5, 5}
	n2 := len(arr2)
	expected2 := []int{5, 5, 5, 5, 5}
	radixsort(arr2, n2)
	if !reflect.DeepEqual(arr2, expected2) {
		t.Errorf("Test case 2 failed; got %v, want %v", arr2, expected2)
	} else {
		t.Logf("Test case 2 succeeded")
	}

	// Test case 3: Edge case of having the array in reverse order
	arr3 := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	n3 := len(arr3)
	expected3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	radixsort(arr3, n3)
	if !reflect.DeepEqual(arr3, expected3) {
		t.Errorf("Test case 3 failed; got %v, want %v", arr3, expected3)
	} else {
		t.Logf("Test case 3 succeeded")
	}
}
