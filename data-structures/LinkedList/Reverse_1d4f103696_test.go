package LinkedList

import (
	"testing"
)

// TestReverse_1d4f103696 tests the Reverse function.
func TestReverse_1d4f103696(t *testing.T) {
	// Create a new linked list.
	ll := NewLinkedList()

	// Add some nodes to the linked list.
	ll.Add(NewNode(10))
	ll.Add(NewNode(20))
	ll.Add(NewNode(30))

	// Reverse the linked list.
	ll.Reverse()

	// Check that the linked list is now reversed.
	if ll.head.val != 30 || ll.tail.val !=10 {
		t.Errorf("Expected the linked list to be reversed, but it was not")
	}

	// Log a message indicating that the test was successful.
	t.Log("TestReverse_1d4f103696 was successful.")
}
