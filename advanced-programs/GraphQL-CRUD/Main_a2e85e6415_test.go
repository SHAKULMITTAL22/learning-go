// Test code generated by AssistantGPT for test azure-32k-go using AI Type Azure Open AI and AI Model assistant-gpt4-32k

package main

import (
	"testing"
)

func add(a int, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	total := add(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}

	total = add(0, 0)
	if total != 0 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 0)
	}
	
	total = add(-5, -5)
	if total != -10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, -10)
	}
}
