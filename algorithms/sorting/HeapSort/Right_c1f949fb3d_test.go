package HeapSort

import (
	"testing"
)

type Heap struct {
	heapArray []int
}

func (h *Heap) Right(i int) int {
	return 2*i + 2
}

// TestRight_c1f949fb3d is a test function for Right method of Heap
func TestRight_c1f949fb3d(t *testing.T) {
	heap := &Heap{}

	// Test case 1: Testing with root as 0
	// Expected output is 2
	result := heap.Right(0)
	if result != 2 {
		t.Errorf("TestRight_c1f949fb3d failed, expected: %v, got: %v", 2, result)
	} else {
		t.Logf("TestRight_c1f949fb3d passed, expected: %v, got: %v", 2, result)
	}

	// Test case 2: Testing with root as 3
	// Expected output is 8
	result = heap.Right(3)
	if result != 8 {
		t.Errorf("TestRight_c1f949fb3d failed, expected: %v, got: %v", 8, result)
	} else {
		t.Logf("TestRight_c1f949fb3d passed, expected: %v, got: %v", 8, result)
	}

	// Test case 3: Testing with root as -1
	// Expected output is 0
	result = heap.Right(-1)
	if result != 0 {
		t.Errorf("TestRight_c1f949fb3d failed, expected: %v, got: %v", 0, result)
	} else {
		t.Logf("TestRight_c1f949fb3d passed, expected: %v, got: %v", 0, result)
	}
}
