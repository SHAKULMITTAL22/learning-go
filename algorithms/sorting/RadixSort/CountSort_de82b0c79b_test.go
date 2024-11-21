package RadixSort

import (
	"reflect"
	"testing"
)

func TestCountSort_de82b0c79b(t *testing.T) {
	// Test case 1: Normal scenario
	arr1 := []int{170, 45, 75, 90, 802, 24, 2, 66}
	exp1 := 1
	countSort(arr1, len(arr1), exp1)
	expectedArr1 := []int{2, 24, 45, 66, 75, 90, 170, 802}
	if !reflect.DeepEqual(arr1, expectedArr1) {
		t.Errorf("TestCountSort_de82b0c79b failed, arr: %v, exp: %d, got: %v, want: %v", arr1, exp1, arr1, expectedArr1)
	} else {
		t.Logf("TestCountSort_de82b0c79b success, arr: %v, exp: %d, got: %v", arr1, exp1, arr1)
	}

	// Test case 2: Array with negative numbers
	arr2 := []int{-170, -45, -75, -90, -802, -24, -2, -66}
	exp2 := 1
	countSort(arr2, len(arr2), exp2)
	expectedArr2 := []int{-802, -170, -90, -75, -66, -45, -24, -2}
	if !reflect.DeepEqual(arr2, expectedArr2) {
		t.Errorf("TestCountSort_de82b0c79b failed, arr: %v, exp: %d, got: %v, want: %v", arr2, exp2, arr2, expectedArr2)
	} else {
		t.Logf("TestCountSort_de82b0c79b success, arr: %v, exp: %d, got: %v", arr2, exp2, arr2)
	}

	// Test case 3: Array with all elements the same
	arr3 := []int{2, 2, 2, 2, 2, 2, 2, 2}
	exp3 := 1
	countSort(arr3, len(arr3), exp3)
	expectedArr3 := []int{2, 2, 2, 2, 2, 2, 2, 2}
	if !reflect.DeepEqual(arr3, expectedArr3) {
		t.Errorf("TestCountSort_de82b0c79b failed, arr: %v, exp: %d, got: %v, want: %v", arr3, exp3, arr3, expectedArr3)
	} else {
		t.Logf("TestCountSort_de82b0c79b success, arr: %v, exp: %d, got: %v", arr3, exp3, arr3)
	}
}
