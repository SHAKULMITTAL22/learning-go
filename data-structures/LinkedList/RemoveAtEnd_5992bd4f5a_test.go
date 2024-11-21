package LinkedList

import "testing"

func TestRemoveAtEnd_5992bd4f5a(t *testing.T) {
	// Test case 1: Remove from empty list
	ll := LinkedList{}
	if ll.RemoveAtEnd() != -1 {
		t.Error("RemoveAtEnd() should return -1 for empty list")
	}

	// Test case 2: Remove from list with one element
	ll = LinkedList{head: &Node{val: 1}}
	if ll.RemoveAtEnd() != 1 {
		t.Error("RemoveAtEnd() should return 1 for list with one element")
	}

	// Test case 3: Remove from list with multiple elements
	ll = LinkedList{head: &Node{val: 1}, tail: &Node{val: 3}}
	ll.tail.prev = ll.head
	ll.head.next = ll.tail
	if ll.RemoveAtEnd() != 3 {
		t.Error("RemoveAtEnd() should return 3 for list with multiple elements")
	}
}
