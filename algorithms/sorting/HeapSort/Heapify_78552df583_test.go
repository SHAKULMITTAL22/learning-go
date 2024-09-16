package HeapSort

import (
	"testing"
)

// TestHeapify_78552df583 method to test Heapify method
func TestHeapify_78552df583(t *testing.T) {
	heap := &Heap{}

	// Test case 1
	{
		array := []int{1, 2, 3, 4, 5}
		expected := []int{5, 4, 3, 1, 2}
		heap.Heapify(array, 0, len(array))
		for i, v := range array {
			if v != expected[i] {
				t.Errorf("Test case 1 failed! Expected %v, got %v", expected, array)
				break
			}
		}
		t.Log("Test case 1 passed!")
	}

	// Test case 2
	{
		array := []int{5, 4, 3, 2, 1}
		expected := []int{5, 4, 3, 2, 1}
		heap.Heapify(array, 0, len(array))
		for i, v := range array {
			if v != expected[i] {
				t.Errorf("Test case 2 failed! Expected %v, got %v", expected, array)
				break
			}
		}
		t.Log("Test case 2 passed!")
	}
}
