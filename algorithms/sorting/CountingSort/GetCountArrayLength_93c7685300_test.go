package CountingSort

import (
	"testing"
)

func TestGetCountArrayLength_93c7685300(t *testing.T) {
	// Test case 1: Empty array
	arr1 := []int{}
	expected := 0
	actual := getCountArrayLength(arr1)
	if actual != expected {
		t.Errorf("TestGetCountArrayLength_93c7685300 failed. Expected: %d, Actual: %d", expected, actual)
	} else {
		t.Log("TestGetCountArrayLength_93c7685300 passed.")
	}

	// Test case 2: Array with positive values
	arr2 := []int{1, 2, 3, 4, 5}
	expected = 6
	actual = getCountArrayLength(arr2)
	if actual != expected {
		t.Errorf("TestGetCountArrayLength_93c7685300 failed. Expected: %d, Actual: %d", expected, actual)
	} else {
		t.Log("TestGetCountArrayLength_93c7685300 passed.")
	}

	// Test case 3: Array with negative values
	arr3 := []int{-1, -2, -3, -4, -5}
	expected = 6
	actual = getCountArrayLength(arr3)
	if actual != expected {
		t.Errorf("TestGetCountArrayLength_93c7685300 failed. Expected: %d, Actual: %d", expected, actual)
	} else {
		t.Log("TestGetCountArrayLength_93c7685300 passed.")
	}

	// Test case 4: Array with mixed values
	arr4 := []int{1, -2, 3, -4, 5}
	expected = 7
	actual = getCountArrayLength(arr4)
	if actual != expected {
		t.Errorf("TestGetCountArrayLength_93c7685300 failed. Expected: %d, Actual: %d", expected, actual)
	} else {
		t.Log("TestGetCountArrayLength_93c7685300 passed.")
	}

	// Test case 5: Array with duplicate values
	arr5 := []int{1, 1, 2, 2, 3, 3}
	expected = 4
	actual = getCountArrayLength(arr5)
	if actual != expected {
		t.Errorf("TestGetCountArrayLength_93c7685300 failed. Expected: %d, Actual: %d", expected, actual)
	} else {
		t.Log("TestGetCountArrayLength_93c7685300 passed.")
	}
}
