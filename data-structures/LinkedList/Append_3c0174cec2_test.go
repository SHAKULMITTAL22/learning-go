package LinkedList

import (
	"testing"
)

func TestAppend_3c0174cec2(t *testing.T) {
	// Test case 1: Append to an empty list
	ll := LinkedList{}
	ll.Append(1)
	if ll.head.val != 1 {
		t.Error("Append to an empty list failed")
	}

	// Test case 2: Append to a non-empty list
	ll = LinkedList{}
	ll.Append(1)
	ll.Append(2)
	if ll.head.next.val != 2 {
		t.Error("Append to a non-empty list failed")
	}

	// Test case 3: Append a nil value
	ll = LinkedList{}
	ll.Append(nil)
	if ll.head != nil {
		t.Error("Append a nil value failed")
	}
}
