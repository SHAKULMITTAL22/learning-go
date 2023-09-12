package IsPowerOfTwo

import (
	"testing"
)

func TestIsPowerOfTwoBitwise_e8104160a5(t *testing.T) {
	// Test case 1: num is a power of 2
	num := 16
	expected := false
	actual := isPowerOfTwoBitwise(num)
	if actual != expected {
		t.Errorf("Expected %t, got %t", expected, actual)
	}

	// Test case 2: num is not a power of 2
	num = 15
	expected = false
	actual = isPowerOfTwoBitwise(num)
	if actual != expected {
		t.Errorf("Expected %t, got %t", expected, actual)
	}

	// Test case 3: num is 0
	num = 0
	expected = false
	actual = isPowerOfTwoBitwise(num)
	if actual != expected {
		t.Errorf("Expected %t, got %t", expected, actual)
	}

	// Test case 4: num is negative
	num = -16
	expected = false
	actual = isPowerOfTwoBitwise(num)
	if actual != expected {
		t.Errorf("Expected %t, got %t", expected, actual)
	}
}
