package LinkedList

import (
	"testing"
)

func TestNewNode_60e8401887(t *testing.T) {
	// Test case 1: Create a new node with a valid value
	node := NewNode(10)
	if node.val != 10 {
		t.Errorf("Expected node value to be 10, got %d", node.val)
	}
	if node.next != nil {
		t.Errorf("Expected node next to be nil, got %v", node.next)
	}
	if node.prev != nil {
		t.Errorf("Expected node prev to be nil, got %v", node.prev)
	}

	// Test case 2: Create a new node with a negative value
	node = NewNode(-10)
	if node.val != -10 {
		t.Errorf("Expected node value to be -10, got %d", node.val)
	}
	if node.next != nil {
		t.Errorf("Expected node next to be nil, got %v", node.next)
	}
	if node.prev != nil {
		t.Errorf("Expected node prev to be nil, got %v", node.prev)
	}

	// Test case 3: Create a new node with a zero value
	node = NewNode(0)
	if node.val != 0 {
		t.Errorf("Expected node value to be 0, got %d", node.val)
	}
	if node.next != nil {
		t.Errorf("Expected node next to be nil, got %v", node.next)
	}
	if node.prev != nil {
		t.Errorf("Expected node prev to be nil, got %v", node.prev)
	}
}
