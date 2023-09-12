// Test generated by RoostGPT for test math-go using AI Type Vertex AI and AI Model code-bison

package Radian

import (
	"math"
	"testing"
)

func TestRadianToDegree_9d9fa92dd5(t *testing.T) {
	// Test case 1: Valid radian value
	expected := 180.0
	actual := radianToDegree(math.Pi)
	if actual != expected {
		t.Errorf("Expected %f, got %f", expected, actual)
	}

	// Test case 2: Negative radian value
	expected = -180.0
	actual = radianToDegree(-math.Pi)
	if actual != expected {
		t.Errorf("Expected %f, got %f", expected, actual)
	}

	// Test case 3: Zero radian value
	expected = 0.0
	actual = radianToDegree(0)
	if actual != expected {
		t.Errorf("Expected %f, got %f", expected, actual)
	}

	// Test case 4: Radian value greater than 2*Pi
	expected = 360.0
	actual = radianToDegree(2 * math.Pi)
	if actual != expected {
		t.Errorf("Expected %f, got %f", expected, actual)
	}

	// Test case 5: Radian value less than -2*Pi
	expected = -360.0
	actual = radianToDegree(-2 * math.Pi)
	if actual != expected {
		t.Errorf("Expected %f, got %f", expected, actual)
	}
}
