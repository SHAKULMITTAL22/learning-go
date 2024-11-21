package LinkedList

import (
	"fmt"
	"testing"
)

func TestPrepend_747252e40f(t *testing.T) {
	// Test case 1: Prepend to an empty list
	ll := NewLinkedList()
	ll.Prepend(1)
	if ll.head.val != 1 {
		t.Error("Expected head value to be 1, got", ll.head.val)
	}

	// Test case 2: Prepend to a non-empty list
	ll = NewLinkedList()
	ll.Prepend(1)
	ll.Prepend(2)
	if ll.head.val != 2 {
		t.Error("Expected head value to be 2, got", ll.head.val)
	}
	if ll.head.next.val != 1 {
		t.Error("Expected next node value to be 1, got", ll.head.next.val)
	}

	fmt.Println("Test Prepend_747252e40f passed")
}
