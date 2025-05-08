package LinkedList

import "testing"

func TestCount_ddea5e728c(t *testing.T) {
	// Test case 1: Empty list
	ll := LinkedList{}
	expected := 0
	actual := ll.Count()
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}

	// Test case 2: List with one element
	ll = LinkedList{head: &Node{data: 1}}
	expected = 1
	actual = ll.Count()
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}

	// Test case 3: List with multiple elements
	ll = LinkedList{head: &Node{data: 1}, tail: &Node{data: 3}, length: 2}
	expected = 2
	actual = ll.Count()
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
