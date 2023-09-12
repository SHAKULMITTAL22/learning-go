package LinkedList

import (
	"fmt"
	"testing"
)

func TestDisplay_679751592e(t *testing.T) {
	// Test case 1: Empty list
	ll := &LinkedList{}
	ll.Display()
	expected := ""
	if ll.String() != expected {
		t.Errorf("Expected: %s, got: %s", expected, ll.String())
	}

	// Test case 2: List with one element
	ll = &LinkedList{}
	ll.Insert(1)
	ll.Display()
	expected = "1 "
	if ll.String() != expected {
		t.Errorf("Expected: %s, got: %s", expected, ll.String())
	}

	// Test case 3: List with multiple elements
	ll = &LinkedList{}
	ll.Insert(1)
	ll.Insert(2)
	ll.Insert(3)
	ll.Display()
	expected = "1 2 3 "
	if ll.String() != expected {
		t.Errorf("Expected: %s, got: %s", expected, ll.String())
	}
}
