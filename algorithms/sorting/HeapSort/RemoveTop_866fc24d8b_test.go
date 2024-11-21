package HeapSort

import (
	"testing"
)

type Heap struct {
	heapArray []int
}

func (h *Heap) RemoveTop(array []int, size int) {
	if size == 0 {
		return
	}
	h.heapArray = array
	h.heapArray[0] = h.heapArray[size-1]
	h.heapArray = h.heapArray[:size-1]
	h.heapify(0)
}

func (h *Heap) heapify(index int) {
	left := 2*index + 1
	right := 2*index + 2
	largest := index

	if left < len(h.heapArray) && h.heapArray[left] > h.heapArray[largest] {
		largest = left
	}

	if right < len(h.heapArray) && h.heapArray[right] > h.heapArray[largest] {
		largest = right
	}

	if largest != index {
		h.heapArray[index], h.heapArray[largest] = h.heapArray[largest], h.heapArray[index]
		h.heapify(largest)
	}
}

func TestRemoveTop_866fc24d8b(t *testing.T) {
	heap := &Heap{}

	array1 := []int{40, 20, 15, 10, 30}
	heap.RemoveTop(array1, len(array1))
	if array1[0] != 30 {
		t.Errorf("RemoveTop failed for array1, expected 30, got %v", array1[0])
	}

	array2 := []int{10}
	heap.RemoveTop(array2, len(array2))
	if len(array2) != 0 {
		t.Errorf("RemoveTop failed for array2, expected empty array, got %v", array2)
	}

	array3 := []int{}
	heap.RemoveTop(array3, len(array3))
	if len(array3) != 0 {
		t.Errorf("RemoveTop failed for array3, expected empty array, got %v", array3)
	}
}
