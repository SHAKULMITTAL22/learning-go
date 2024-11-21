package LinkedList

import (
	"testing"
)

// TestCount_ddea5e728c tests the Count function.
func TestCount_ddea5e728c(t *testing.T) {
	// Create a new linked list.
	ll := NewLinkedList()

	// Add some nodes to the linked list.
	ll.Add(NewNode(10))
	ll.Add(NewNode(20))
	ll.Add(NewNode(30))

	// Check that the size of the linked list is 3.
	if ll.Count() != 3 {
		t.Errorf("Expected the size of the linked list to be3, but it was %d", ll.Count())
	}

	// Remove a node from the linked list.
	ll.Remove(ll.head)

	// Check that the size of the linked list is now2.
	if ll.Count() !=2 {
		t.Errorf("Expected the size of the linked list to be2, but it was %d", ll.Count())
	}

	// Add a node back to the linked list.
	ll.Add(NewNode(40))

	// Check that the size of the linked list is now3.
	if ll.Count() !=3 {
		t.Errorf("Expected the size of the linked list to be3, but it was %d", ll.Count())
	}

	// Log a message indicating that the test was successful.
	t.Log("TestCount_ddea5e728c was successful.")
}
