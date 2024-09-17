package LinkedList

import (
	"fmt"
	"testing"
)

// TestSwap_60e84018892 tests the Swap function.
func TestSwap_60e84018892(t *testing.T) {
	// Create a new linked list.
	l := NewLinkedList()

	// Create a new node with a value of 10.
	n1 := NewNode(10)

	// Add the new node to the linked list.
	l.Add(n1)

	// Create a new node with a value of 20.
	n2 := NewNode(20)

	// Add the new node to the linked list.
	l.Add(n2)

	// Swap the two nodes.
	l.Swap(n1, n2)

	// Check that the values of the two nodes have been swapped.
	if n1.val != 20 || n2.val != 10 {
		t.Errorf("Expected the values of the two nodes to be swapped, but they were not")
	}

	// Check that the next pointers of the two nodes have been swapped.
	if n1.next != n2 || n2.next != nil {
		t.Errorf("Expected the next pointers of the two nodes to be swapped, but they were not")
	}

	// Check that the prev pointers of the two nodes have been swapped.
	if n1.prev != nil || n2.prev != n1 {
		t.Errorf("Expected the prev pointers of the two nodes to be swapped, but they were not")
	}

	// Log a message indicating that the test was successful.
	fmt.Println("TestSwap_60e84018892 was successful.")
}
