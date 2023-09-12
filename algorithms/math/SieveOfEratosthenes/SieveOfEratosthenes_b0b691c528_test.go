package SieveOfEratosthenes

import "testing"

func TestSieveOfEratosthenes_b0b691c528(t *testing.T) {
	// Test case 1: maxNumber is 10
	expected := []int{2, 3, 5, 7}
	result := sieveOfEratosthenes(10)
	if expected != result {
		t.Errorf("Expected: %v, got: %v", expected, result)
	}

	// Test case 2: maxNumber is 20
	expected = []int{2, 3, 5, 7, 11, 13, 17, 19}
	result = sieveOfEratosthenes(20)
	if expected != result {
		t.Errorf("Expected: %v, got: %v", expected, result)
	}

	// Test case 3: maxNumber is 1
	expected = []int{}
	result = sieveOfEratosthenes(1)
	if expected != result {
		t.Errorf("Expected: %v, got: %v", expected, result)
	}

	// Test case 4: maxNumber is 0
	expected = []int{}
	result = sieveOfEratosthenes(0)
	if expected != result {
		t.Errorf("Expected: %v, got: %v", expected, result)
	}

	// Test case 5: maxNumber is -1
	expected = []int{}
	result = sieveOfEratosthenes(-1)
	if expected != result {
		t.Errorf("Expected: %v, got: %v", expected, result)
	}
}
