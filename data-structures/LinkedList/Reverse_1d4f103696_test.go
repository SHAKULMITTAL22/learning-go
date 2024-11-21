package LinkedList

import (
	"testing"
)

func TestReverse_1d4f103696(t *testing.T) {
	// Test case 1: Empty list
	ll := &LinkedList{}
	ll.Reverse()
	if ll.head != nil {
		t.Error("Expected head to be nil, but got", ll.head)
	}

	// Test case 2: List with one element
	ll = &LinkedList{head: &node{data: 1}}
	ll.Reverse()
	if ll.head.data != 1 {
		t.Error("Expected head.data to be 1, but got", ll.head.data)
	}

	// Test case 3: List with multiple elements
	ll = &LinkedList{head: &node{data: 1}, tail: &node{data: 3}, length: 2}
	ll.head.next = ll.tail
	ll.tail.prev = ll.head
	ll.Reverse()
	if ll.head.data != 3 || ll.tail.data != 1 {
		t.Error("Expected head.data to be 3 and tail.data to be 1, but got", ll.head.data, "and", ll.tail.data)
	}
}
