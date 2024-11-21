package LinkedList

import (
	"testing"
)

// TestDisplayReverse_52a96579ee tests the DisplayReverse function.
func TestDisplayReverse_52a96579ee(t *testing.T) {
	// Create a new linked list.
	ll := NewLinkedList()

	// Add some nodes to the linked list.
	ll.Add(NewNode(10))
	ll.Add(NewNode(20))
	ll.Add(NewNode(30))

	// Display the linked list in reverse.
	ll.DisplayReverse()

	// Check that the output of the DisplayReverse function is correct.
	expectedOutput := "30 2010\n"
	if ll.DisplayReverse() != expectedOutput {
		t.Errorf("Expected the output of the DisplayReverse function to be %s, but it was %s", expectedOutput, ll.DisplayReverse())
	}

	// Log a message indicating that the test was successful.
	t.Log("TestDisplayReverse_52a96579ee was successful.")
}
