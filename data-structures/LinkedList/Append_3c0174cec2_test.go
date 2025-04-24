package LinkedList

import (
	"testing"
)

// TestInsertAfter_747252e40t tests the InsertAfter function.
func TestInsertAfter_747252e40t(t *testing.T) {
	// Create a new linked list.
	ll := NewLinkedList()

	// Create a new node with a value of 10.
	n := NewNode(10)

	// Insert the new node after the head node.
	ll.InsertAfter(ll.head, n)

	// Check that the new node is the second node in the linked list.
	if ll.head.next != n {
		t.Errorf("Expected the new node to be the second node in the linked list, but it was not")
	}

	// Check that the new node's value is 10.
	if n.val != 10 {
		t.Errorf("Expected the new node's value to be 10, but it was not")
	}

	// Check that the new node's next pointer is pointing to the third node in the linked list.
	if n.next != ll.tail {
		t.Errorf("Expected the new node's next pointer to be pointing to the third node in the linked list, but it was not")
	}

	// Check that the new node's prev pointer is pointing to the first node in the linked list.
	if n.prev != ll.head {
		t.Errorf("Expected the new node's prev pointer to be pointing to the first node in the linked list, but it was not")
	}

	// Check that the size of the linked list is now 3.
	if ll.size != 3 {
		t.Errorf("Expected the size of the linked list to be 3, but it was not")
	}

	// Log a message indicating that the test was successful.
	t.Log("TestInsertAfter_747252e40t was successful.")
}
