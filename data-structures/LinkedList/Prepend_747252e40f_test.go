package LinkedList

import (
	"testing"
)

// TestRemoveBefore_747252e40x tests the RemoveBefore function.
func TestRemoveBefore_747252e40x(t *testing.T) {
	// Create a new linked list.
	ll := NewLinkedList()

	// Create a new node with a value of 10.
	n := NewNode(10)

	// Add the new node to the linked list.
	ll.Add(n)

	// Create a new node with a value of 20.
	n2 := NewNode(20)

	// Add the new node to the linked list.
	ll.Add(n2)

	// Remove the node before the second node.
	ll.RemoveBefore(n2)

	// Check that the first node is no longer in the linked list.
	if ll.Contains(n) {
		t.Errorf("Expected the first node to be removed from the linked list, but it was not")
	}

	// Check that the size of the linked list is now 1.
	if ll.Size() != 1 {
		t.Errorf("Expected the size of the linked list to be1, but it was not")
	}

	// Log a message indicating that the test was successful.
	t.Log("TestRemoveBefore_747252e40x was successful.")
}
