package LinkedList

import (
	"testing"
)

// TestRemoveAtBeg_e45f33f4b9 tests the RemoveAtBeg function.
func TestRemoveAtBeg_e45f33f4b9(t *testing.T) {
	// Create a new linked list.
	ll := NewLinkedList()

	// Add some nodes to the linked list.
	ll.Add(NewNode(10))
	ll.Add(NewNode(20))
	ll.Add(NewNode(30))

	// Remove the node at the beginning of the linked list.
	val := ll.RemoveAtBeg()

	// Check that the value of the removed node is 10.
	if val != 10 {
		t.Errorf("Expected the value of the removed node to be 10, but it was %d", val)
	}

	// Check that the size of the linked list is now 2.
	if ll.Size() != 2 {
		t.Errorf("Expected the size of the linked list to be 2, but it was %d", ll.Size())
	}

	// Check that the head of the linked list is now the second node.
	if ll.head.val != 20 {
		t.Errorf("Expected the head of the linked list to be the second node, but it was the first node")
	}

	// Check that the tail of the linked list is still the third node.
	if ll.tail.val != 30 {
		t.Errorf("Expected the tail of the linked list to be the third node, but it was the second node")
	}

	// Log a message indicating that the test was successful.
	t.Log("TestRemoveAtBeg_e45f33f4b9 was successful.")
}
