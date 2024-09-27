package IsPowerOfTwo

import (
	"testing"
)

func TestIsPowerOfTwo_a909b954a6(t *testing.T) {
	// Test case 1: num is a power of two
	num := 16
	expected := true
	actual := isPowerOfTwo(num)
	if actual != expected {
		t.Errorf("TestIsPowerOfTwo_a909b954a6 failed. num: %d, expected: %t, actual: %t", num, expected, actual)
	} else {
		t.Log("TestIsPowerOfTwo_a909b954a6 passed.")
	}

	// Test case 2: num is not a power of two
	num = 15
	expected = false
	actual = isPowerOfTwo(num)
	if actual != expected {
		t.Errorf("TestIsPowerOfTwo_a909b954a6 failed. num: %d, expected: %t, actual: %t", num, expected, actual)
	} else {
		t.Log("TestIsPowerOfTwo_a909b954a6 passed.")
	}

	// Test case 3: num is 0
	num = 0
	expected = false
	actual = isPowerOfTwo(num)
	if actual != expected {
		t.Errorf("TestIsPowerOfTwo_a909b954a6 failed. num: %d, expected: %t, actual: %t", num, expected, actual)
	} else {
		t.Log("TestIsPowerOfTwo_a909b954a6 passed.")
	}

	// Test case 4: num is 1
	num = 1
	expected = false
	actual = isPowerOfTwo(num)
	if actual != expected {
		t.Errorf("TestIsPowerOfTwo_a909b954a6 failed. num: %d, expected: %t, actual: %t", num, expected, actual)
	} else {
		t.Log("TestIsPowerOfTwo_a909b954a6 passed.")
	}

	// Test case 5: num is negative
	num = -16
	expected = false
	actual = isPowerOfTwo(num)
	if actual != expected {
		t.Errorf("TestIsPowerOfTwo_a909b954a6 failed. num: %d, expected: %t, actual: %t", num, expected, actual)
	} else {
		t.Log("TestIsPowerOfTwo_a909b954a6 passed.")
	}
}

func TestMod_10abfc4edf(t *testing.T) {
	// Test case 1: a is positive and b is positive
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

	// Test case 4: a is negative and b is negative
	a = -10
	b = -3
	expected = -1
	actual = mod(a, b)
	if actual != expected {
		t.Errorf("TestMod_10abfc4edf failed. a: %d, b: %d, expected: %d, actual: %d", a, b, expected, actual)
	} else {
		t.Log("TestMod_10abfc4edf passed.")
	}

	// Test case 5: a is 0
	a = 0
	b = 3
	expected = 0
	actual = mod(a, b)
	if actual != expected {
		t.Errorf("TestMod_10abfc4edf failed. a: %d, b: %d, expected: %d, actual: %d", a, b, expected, actual)
	} else {
		t.Log("TestMod_10abfc4edf passed.")
	}

	// Test case 6: b is 0
	a = 10
	b = 0
	// TODO: uncomment the following line to see the panic
	// expected = 1
	// actual = mod(a, b)
	// if actual != expected {
	// 	t.Errorf("TestMod_10abfc4edf failed. a: %d, b: %d, expected: %d, actual: %d", a, b, expected, actual)
	// } else {
	// 	t.Log("TestMod_10abfc4edf passed.")
	// }
}

func TestIsPowerOfTwoBitwise(t *testing.T) {
	// Test case 1: num is a power of two
	num := 16
	expected := true
	actual := isPowerOfTwoBitwise(num)
	if actual != expected {
		t.Errorf("TestIsPowerOfTwoBitwise failed. num: %d, expected: %t, actual: %t", num, expected, actual)
	} else {
		t.Log("TestIsPowerOfTwoBitwise passed.")
	}

	// Test case 2: num is not a power of two
	num = 15
	expected = false
	actual = isPowerOfTwoBitwise(num)
	if actual != expected {
		t.Errorf("TestIsPowerOfTwoBitwise failed. num: %d, expected: %t, actual: %t", num, expected, actual)
	} else {
		t.Log("TestIsPowerOfTwoBitwise passed.")
	}

	// Test case 3: num is 0
	num = 0
	expected = false
	actual = isPowerOfTwoBitwise(num)
	if actual != expected {
		t.Errorf("TestIsPowerOfTwoBitwise failed. num: %d, expected: %t, actual: %t", num, expected, actual)
	} else {
		t.Log("TestIsPowerOfTwoBitwise passed.")
	}

	// Test case 4: num is 1
	num = 1
	expected = true
	actual = isPowerOfTwoBitwise(num)
	if actual != expected {
		t.Errorf("TestIsPowerOfTwoBitwise failed. num: %d, expected: %t, actual: %t", num, expected, actual)
	} else {
		t.Log("TestIsPowerOfTwoBitwise passed.")
	}

	// Test case 5: num is negative
	num = -16
	expected = false
	actual = isPowerOfTwoBitwise(num)
	if actual != expected {
		t.Errorf("TestIsPowerOfTwoBitwise failed. num: %d, expected: %t, actual: %t", num, expected, actual)
	} else {
		t.Log("TestIsPowerOfTwoBitwise passed.")
	}
}
