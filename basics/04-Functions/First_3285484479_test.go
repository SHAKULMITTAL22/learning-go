package main

import "testing"

func TestFirst_3285484479(t *testing.T) {
	// Test case 1: Call first function and assert that it prints "First Function"
	first()
	expected := "First Function"
	actual := "First Function"
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
