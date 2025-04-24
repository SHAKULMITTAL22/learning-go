package LinkedList

import (
	"testing"
)

// TestDisplay_679751592e tests the Display function.
func TestDisplay_679751592e(t *testing.T) {
	// Create a new linked list.
	ll := NewLinkedList()

	// Add some nodes to the linked list.
	ll.Add(NewNode(10))
	ll.Add(NewNode(20))
	ll.Add(NewNode(30))

	// Display the linked list.
	ll.Display()

	// Check that the output of the Display function is correct.
	expectedOutput := "10 2030\n"
	if ll.Display() != expectedOutput {
		t.Errorf("Expected the output of the Display function to be %s, but it was %s", expectedOutput, ll.Display())
	}

	// Log a message indicating that the test was successful.
	t.Log("TestDisplay_679751592e was successful.")
}
