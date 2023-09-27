package IsPowerOfTwo

import (
	"testing"
)

func TestMod_10abfc4edf(t *testing.T) {
	// Test case 1: a and b are both positive
	a := 10
	b := 3
	expected := 1
	actual := mod(a, b)
	if actual != expected {
		t.Errorf("TestMod_10abfc4edf failed. a: %d, b: %d, expected: %d, actual: %d", a, b, expected, actual)
	} else {
		t.Log("TestMod_10abfc4edf passed.")
	}

	// Test case 2: a is negative and b is positive
	a = -10
	b = 3
	expected = 2
	actual = mod(a, b)
	if actual != expected {
		t.Errorf("TestMod_10abfc4edf failed. a: %d, b: %d, expected: %d, actual: %d", a, b, expected, actual)
	} else {
		t.Log("TestMod_10abfc4edf passed.")
	}

	// Test case 3: a is positive and b is negative
	a = 10
	b = -3
	expected = -2
	actual = mod(a, b)
	if actual != expected {
		t.Errorf("TestMod_10abfc4edf failed. a: %d, b: %d, expected: %d, actual: %d", a, b, expected, actual)
	} else {
		t.Log("TestMod_10abfc4edf passed.")
	}

	// Test case 4: a and b are both negative
	a = -10
	b = -3
	expected = -1
	actual = mod(a, b)
	if actual != expected {
		t.Errorf("TestMod_10abfc4edf failed. a: %d, b: %d, expected: %d, actual: %d", a, b, expected, actual)
	} else {
		t.Log("TestMod_10abfc4edf passed.")
	}

	// Test case 5: b is 0
	a = 10
	b = 0
	expected = 0
	actual = mod(a, b)
	if actual != expected {
		t.Errorf("TestMod_10abfc4edf failed. a: %d, b: %d, expected: %d, actual: %d", a, b, expected, actual)
	} else {
		t.Log("TestMod_10abfc4edf passed.")
	}

	// Test case 6: a is 0
	a = 0
	b = 3
	expected = 0
	actual = mod(a, b)
	if actual != expected {
		t.Errorf("TestMod_10abfc4edf failed. a: %d, b: %d, expected: %d, actual: %d", a, b, expected, actual)
	} else {
		t.Log("TestMod_10abfc4edf passed.")
	}
}
