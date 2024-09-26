package LinkedList

import (
	"fmt"
	"testing"
)

func TestDisplayReverse_52a96579ee(t *testing.T) {
	// Test case 1: Empty list
	ll := &LinkedList{}
	ll.DisplayReverse()
	if ll.head != nil {
		t.Error("Expected head to be nil, but got", ll.head)
	}

	// Test case 2: List with one element
	ll = &LinkedList{}
	ll.InsertAtHead(1)
	ll.DisplayReverse()
	if ll.head.val != 1 {
		t.Error("Expected head.val to be 1, but got", ll.head.val)
	}

	// Test case 3: List with multiple elements
	ll = &LinkedList{}
	ll.InsertAtHead(1)
	ll.InsertAtHead(2)
	ll.InsertAtHead(3)
	ll.DisplayReverse()
	if ll.head.val != 3 || ll.head.next.val != 2 || ll.head.next.next.val != 1 {
		t.Error("Expected head.val to be 3, head.next.val to be 2, and head.next.next.val to be 1, but got", ll.head.val, ll.head.next.val, ll.head.next.next.val)
	}
}
