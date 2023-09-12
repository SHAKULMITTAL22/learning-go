package SquareRoot

import (
	"math"
	"testing"
)

func TestNewton_d937a20fc3(t *testing.T) {
	// Test case 1: Positive number
	z := 10.0
	x := 25.0
	expected := 5.0
	actual := newton(z, x)
	if actual != expected {
		t.Errorf("Expected %f, got %f", expected, actual)
	}

	// Test case 2: Negative number
	z = -10.0
	x = 100.0
	expected := 10.0
	actual = newton(z, x)
	if actual != expected {
		t.Errorf("Expected %f, got %f", expected, actual)
	}

	// Test case 3: Zero
	z = 0.0
	x = 0.0
	expected := 0.0
	actual = newton(z, x)
	if actual != expected {
		t.Errorf("Expected %f, got %f", expected, actual)
	}

	// Test case 4: NaN
	z = math.NaN()
	x = math.NaN()
	expected := math.NaN()
	actual = newton(z, x)
	if actual != expected {
		t.Errorf("Expected %f, got %f", expected, actual)
	}

	// Test case 5: Infinity
	z = math.Inf(1)
	x = math.Inf(1)
	expected := math.Inf(1)
	actual = newton(z, x)
	if actual != expected {
		t.Errorf("Expected %f, got %f", expected, actual)
	}

	// Test case 6: Negative infinity
	z = math.Inf(-1)
	x = math.Inf(-1)
	expected := math.Inf(-1)
	actual = newton(z, x)
	if actual != expected {
		t.Errorf("Expected %f, got %f", expected, actual)
	}
}
