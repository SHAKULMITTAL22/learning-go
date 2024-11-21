package LinkedList

import (
	"testing"
)

// TestRemoveAtEnd_5992bd4f5a tests the RemoveAtEnd function.
func TestRemoveAtEnd_5992bd4f5a(t *testing.T) {
	// Create a new linked list.
	ll := NewLinkedList()

	// Add some nodes to the linked list.
	ll.Add(NewNode(10))
	ll.Add(NewNode(20))
	ll.Add(NewNode(30))

	// Remove the node at the end of the linked list.
	val := ll.RemoveAtEnd()

	// Check that the value of the removed node is 30.
	if val != 30 {
		t.Errorf("Expected the value of the removed node to be 30, but it was %d", val)
	}

	// Check that the size of the linked list is now 2.
	if ll.Size() != 2 {
		t.Errorf("Expected the size of the linked list to be 2, but it was %d", ll.Size())
	}

	// Check that the head of the linked list is still the first node.
	if ll.head.val != 10 {
		t.Errorf("Expected the head of the linked list to be the first node, but it was the second node")
	}

	// Check that the tail of the linked list is now the second node.
	if ll.tail.val != 20 {
		t.Errorf("Expected the tail of the linked list to be the second node, but it was the third node")
	}

	// Log a message indicating that the test was successful.
	t.Log("TestRemoveAtEnd_5992bd4f5a was successful.")
}
