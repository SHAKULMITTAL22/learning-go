package LinkedList

import (
	"testing"
)

func TestRemoveAtBeg_e45f33f4b9(t *testing.T) {
	ll := NewLinkedList()

	// Test case 1: Remove from an empty list
	val := ll.RemoveAtBeg()
	if val != -1 {
		t.Errorf("Expected -1, got %d", val)
	}

	// Test case 2: Remove from a list with one element
	ll.AddAtBeg(1)
	val = ll.RemoveAtBeg()
	if val != 1 {
		t.Errorf("Expected 1, got %d", val)
	}

	// Test case 3: Remove from a list with multiple elements
	ll.AddAtBeg(2)
	ll.AddAtBeg(3)
	val = ll.RemoveAtBeg()
	if val != 3 {
		t.Errorf("Expected 3, got %d", val)
	}
	val = ll.RemoveAtBeg()
	if val != 2 {
		t.Errorf("Expected 2, got %d", val)
	}
}
