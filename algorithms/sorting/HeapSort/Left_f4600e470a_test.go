package HeapSort

import (
	"testing"
)

type HeapTest struct {}

func (h *HeapTest) LeftTest(root int) int {
	return (root * 2) + 1
}

func TestLeft_f4600e470a(t *testing.T) {
	h := &HeapTest{}

	// Test case 1: Normal scenario
	root := 3
	expected := 7
	if result := h.LeftTest(root); result != expected {
		t.Errorf("Test case 1 failed. Expected %d but got %d", expected, result)
	} else {
		t.Logf("Test case 1 succeeded")
	}

	// Test case 2: Edge case with root as 0
	root = 0
	expected = 1
	if result := h.LeftTest(root); result != expected {
		t.Errorf("Test case 2 failed. Expected %d but got %d", expected, result)
	} else {
		t.Logf("Test case 2 succeeded")
	}

	// Test case 3: Negative root
	root = -2
	expected = -3
	if result := h.LeftTest(root); result != expected {
		t.Errorf("Test case 3 failed. Expected %d but got %d", expected, result)
	} else {
		t.Logf("Test case 3 succeeded")
	}
}
