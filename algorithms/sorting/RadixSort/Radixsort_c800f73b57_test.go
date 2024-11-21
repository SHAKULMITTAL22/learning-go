package RadixSort

import (
	"testing"
)

func TestRadixsort_c800f73b57(t *testing.T) {
	// Test case 1: Normal case
	arr1 := []int{170, 45, 75, 90, 802, 24, 2, 66}
	n1 := len(arr1)
	expected1 := []int{2, 24, 45, 66, 75, 90, 170, 802}
	radixsort(arr1, n1)
	for i, v := range arr1 {
		if v != expected1[i] {
			t.Errorf("Test case 1 failed: got %v, want %v", arr1, expected1)
			break
		}
	}
	t.Log("Test case 1 passed")

	// Test case 2: Edge case (array with one element)
	arr3 := []int{5}
	n3 := len(arr3)
	expected3 := []int{5}
	radixsort(arr3, n3)
	for i, v := range arr3 {
		if v != expected3[i] {
			t.Errorf("Test case 3 failed: got %v, want %v", arr3, expected3)
			break
		}
	}
	t.Log("Test case 3 passed")

	// Test case 4: Edge case (array with negative numbers)
	arr4 := []int{-5, -1, -7, -3}
	n4 := len(arr4)
	expected4 := []int{-7, -5, -3, -1}
	radixsort(arr4, n4)
	for i, v := range arr4 {
		if v != expected4[i] {
			t.Errorf("Test case 4 failed: got %v, want %v", arr4, expected4)
			break
		}
	}
	t.Log("Test case 4 passed")
}
