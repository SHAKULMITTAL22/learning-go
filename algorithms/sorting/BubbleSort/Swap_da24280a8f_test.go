package BubbleSort

import (
	"reflect"
	"testing"
)

func TestSwap_da24280a8f(t *testing.T) {
	// Test case 1: Check if elements at two indices are swapped correctly
	{
		arr := []int{1, 2, 3, 4, 5}
		swapTest(arr, 1, 3)
		expectedArr := []int{1, 4, 3, 2, 5}
		if !reflect.DeepEqual(arr, expectedArr) {
			t.Errorf("Test case 1 failed. Expected %v, got %v", expectedArr, arr)
		} else {
			t.Log("Test case 1 passed.")
		}
	}

	// Test case 2: Check if elements are swapped correctly when both indices are same
	{
		arr := []int{1, 2, 3, 4, 5}
		swapTest(arr, 2, 2)
		expectedArr := []int{1, 2, 3, 4, 5}
		if !reflect.DeepEqual(arr, expectedArr) {
			t.Errorf("Test case 2 failed. Expected %v, got %v", expectedArr, arr)
		} else {
			t.Log("Test case 2 passed.")
		}
	}

	// Test case 3: Check if elements are swapped correctly for an array with duplicate elements
	{
		arr := []int{1, 2, 2, 1, 3}
		swapTest(arr, 1, 3)
		expectedArr := []int{1, 1, 2, 2, 3}
		if !reflect.DeepEqual(arr, expectedArr) {
			t.Errorf("Test case 3 failed. Expected %v, got %v", expectedArr, arr)
		} else {
			t.Log("Test case 3 passed.")
		}
	}
}

func swapTest(arr []int, i, j int) {
	tmp := arr[j]
	arr[j] = arr[i]
	arr[i] = tmp
}
