package FastPowering

import (
	"testing"
)

func TestFastPowering_be2e642e07(t *testing.T) {
	// Test case 1: base = 2, power = 10
	expected := 1024.0
	actual := fastPowering(2, 10)
	if actual != expected {
		t.Errorf("Expected %f, got %f", expected, actual)
	}

	// Test case 2: base = -3, power = 3
	expected = -27.0
	actual = fastPowering(-3, 3)
	if actual != expected {
		t.Errorf("Expected %f, got %f", expected, actual)
	}

	// Test case 3: base = 0, power = 5
	expected = 0.0
	actual = fastPowering(0, 5)
	if actual != expected {
		t.Errorf("Expected %f, got %f", expected, actual)
	}

	// Test case 4: base = 1.5, power = -2
	expected = 0.4444444444444444
	actual = fastPowering(1.5, -2)
	if actual != expected {
		t.Errorf("Expected %f, got %f", expected, actual)
	}

	// Test case 5: base = 10, power = 0
	expected = 1.0
	actual = fastPowering(10, 0)
	if actual != expected {
		t.Errorf("Expected %f, got %f", expected, actual)
	}

	// Test case 6: base = 10, power = 1
	expected = 10.0
	actual = fastPowering(10, 1)
	if actual != expected {
		t.Errorf("Expected %f, got %f", expected, actual)
	}

	// Test case 7: base = 10, power = 2
	expected = 100.0
	actual = fastPowering(10, 2)
	if actual != expected {
		t.Errorf("Expected %f, got %f", expected, actual)
	}

	// Test case 8: base = 10, power = 3
	expected = 1000.0
	actual = fastPowering(10, 3)
	if actual != expected {
		t.Errorf("Expected %f, got %f", expected, actual)
	}

	// Test case 9: base = 10, power = 4
	expected = 10000.0
	actual = fastPowering(10, 4)
	if actual != expected {
		t.Errorf("Expected %f, got %f", expected, actual)
	}

	// Test case 10: base = 10, power = 5
	expected = 100000.0
	actual = fastPowering(10, 5)
	if actual != expected {
		t.Errorf("Expected %f, got %f", expected, actual)
	}
}
