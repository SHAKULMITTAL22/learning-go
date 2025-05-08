package HeapSort

import (
	"reflect"
	"testing"
)

type HeapTest struct{}

func (heap *HeapTest) HeapSort(array []int) []int {
	heap.BuildHeap(array)

	for length := len(array); length > 1; length-- {
		heap.RemoveTop(array, length)
	}

	return array
}

func (heap *HeapTest) BuildHeap(array []int) {
	// TODO: Implement this method
}

func (heap *HeapTest) RemoveTop(array []int, length int) {
	// TODO: Implement this method
}

func TestHeapSort_d92a254e00(t *testing.T) {
	heap := &HeapTest{}

	t.Run("Test with normal array", func(t *testing.T) {
		array := []int{5, 2, 3, 1, 4}
		expected := []int{1, 2, 3, 4, 5}
		result := heap.HeapSort(array)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("HeapSort(%v) = %v, want %v", array, result, expected)
		} else {
			t.Logf("Success: HeapSort(%v) = %v", array, result)
		}
	})

	t.Run("Test with empty array", func(t *testing.T) {
		array := []int{}
		expected := []int{}
		result := heap.HeapSort(array)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("HeapSort(%v) = %v, want %v", array, result, expected)
		} else {
			t.Logf("Success: HeapSort(%v) = %v", array, result)
		}
	})

	t.Run("Test with array of same elements", func(t *testing.T) {
		array := []int{1, 1, 1, 1, 1}
		expected := []int{1, 1, 1, 1, 1}
		result := heap.HeapSort(array)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("HeapSort(%v) = %v, want %v", array, result, expected)
		} else {
			t.Logf("Success: HeapSort(%v) = %v", array, result)
		}
	})
}
