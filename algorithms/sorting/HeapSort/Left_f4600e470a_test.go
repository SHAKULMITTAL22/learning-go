package HeapSort

import (
	"testing"
)

type Heap struct {
	heapArray []int
}

func (h *Heap) Left(i int) int {
	return 2*i + 1
}

func TestLeft_f4600e470a(t *testing.T) {
	heap := Heap{}

	// Test Case 1: Testing with root value 0
	root := 0
	expected := 1
	got := heap.Left(root)
	if got != expected {
		t.Errorf("Test Case 1 Failed: Got %v but Expected %v", got, expected)
	} else {
		t.Logf("Test Case 1 Passed")
	}

	// Test Case 2: Testing with root value 2
	root = 2
	expected = 5
	got = heap.Left(root)
	if got != expected {
		t.Errorf("Test Case 2 Failed: Got %v but Expected %v", got, expected)
	} else {
		t.Logf("Test Case 2 Passed")
	}

	// Test Case 3: Testing with negative root value
	root = -2
	expected = -3
	got = heap.Left(root)
	if got != expected {
		t.Errorf("Test Case 3 Failed: Got %v but Expected %v", got, expected)
	} else {
		t.Logf("Test Case 3 Passed")
	}
}
