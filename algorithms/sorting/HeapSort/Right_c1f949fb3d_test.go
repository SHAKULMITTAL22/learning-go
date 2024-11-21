package HeapSort

import (
	"testing"
)

type TestHeap struct {}

func (h *TestHeap) Right(root int) int {
	return (root * 2) + 2
}

func TestRight_c1f949fb3d(t *testing.T) {
	h := &TestHeap{}

	// Test case 1: Testing with a positive number
	root := 5
	expected := 12
	got := h.Right(root)
	if got != expected {
		t.Errorf("Right(%v) = %v; want %v", root, got, expected)
	} else {
		t.Logf("Success: Right(%v) = %v", root, got)
	}

	// Test case 2: Testing with zero
	root = 0
	expected = 2
	got = h.Right(root)
	if got != expected {
		t.Errorf("Right(%v) = %v; want %v", root, got, expected)
	} else {
		t.Logf("Success: Right(%v) = %v", root, got)
	}

	// Test case 3: Testing with a negative number
	root = -3
	expected = -4
	got = h.Right(root)
	if got != expected {
		t.Errorf("Right(%v) = %v; want %v", root, got, expected)
	} else {
		t.Logf("Success: Right(%v) = %v", root, got)
	}
}
