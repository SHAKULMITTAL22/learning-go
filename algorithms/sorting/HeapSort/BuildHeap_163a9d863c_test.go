package HeapSort

import (
	"testing"
	"reflect"
)

type HeapTest struct {
	heapArray []int
}

func (heap *HeapTest) BuildHeapTest(array []int) {
	for i := len(array) / 2; i >= 0; i-- {
		heap.HeapifyTest(array, i, len(array))
	}
}

func (heap *HeapTest) HeapifyTest(array []int, idx int, max int) {
	largest, left, right := idx, 2*idx+1, 2*idx+2
	if left < max && array[left] > array[largest] {
		largest = left
	}
	if right < max && array[right] > array[largest] {
		largest = right
	}
	if largest != idx {
		array[idx], array[largest] = array[largest], array[idx]
		heap.HeapifyTest(array, largest, max)
	}
	heap.heapArray = array
}

func TestBuildHeapTest_163a9d863c(t *testing.T) {
	tests := []struct {
		name     string
		array    []int
		expected []int
	}{
		{
			name:     "Test Case 1: Valid array of integers",
			array:    []int{9, 5, 6, 2, 3},
			expected: []int{9, 5, 6, 2, 3},
		},
		{
			name:     "Test Case 2: Array of integers with negative numbers",
			array:    []int{-1, -2, -3, -4, -5},
			expected: []int{-1, -2, -3, -4, -5},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			heap := &HeapTest{}
			heap.BuildHeapTest(test.array)

			if !reflect.DeepEqual(heap.heapArray, test.expected) {
				t.Errorf("Method failed for %v. Expected %v but got %v", test.array, test.expected, heap.heapArray)
			} else {
				t.Logf("Method passed for %v. Expected %v and got %v", test.array, test.expected, heap.heapArray)
			}
		})
	}
}
