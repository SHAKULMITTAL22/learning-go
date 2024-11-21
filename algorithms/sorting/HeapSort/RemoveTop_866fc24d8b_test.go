package HeapSort

import (
	"testing"
)

type Heap struct {
	Heapify func([]int, int, int)
}

func (h *Heap) RemoveTop(array []int, length int) {
	array[0], array[length-1] = array[length-1], array[0]
	h.Heapify(array, 0, length-2)
}

func TestRemoveTop_866fc24d8b(t *testing.T) {
	mockHeapify := func(array []int, root, end int) {
		for root*2+1 <= end {
			child := root*2 + 1

			if child+1 <= end && array[child] < array[child+1] {
				child++
			}

			if array[root] < array[child] {
				array[root], array[child] = array[child], array[root]
				root = child
			} else {
				return
			}
		}
	}

	{
		array := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
		expected := []int{9, 6, 8, 3, 5, 1, 4, 2, 7, 10}
		heap := Heap{Heapify: mockHeapify}

		heap.RemoveTop(array, len(array))

		for i, v := range array {
			if v != expected[i] {
				t.Errorf("Test case 1 failed: got %v, want %v", array, expected)
				break
			}
		}

		t.Log("Test case 1 passed")
	}

	{
		array := []int{1}
		expected := []int{1}
		heap := Heap{Heapify: mockHeapify}

		heap.RemoveTop(array, len(array))

		for i, v := range array {
			if v != expected[i] {
				t.Errorf("Test case 2 failed: got %v, want %v", array, expected)
				break
			}
		}

		t.Log("Test case 2 passed")
	}
}
