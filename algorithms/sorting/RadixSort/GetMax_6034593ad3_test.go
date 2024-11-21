// Test generated by RoostGPT for test math-go using AI Type Open AI and AI Model gpt-4

package RadixSort

import (
	"testing"
)

func TestGetMax_6034593ad3(t *testing.T) {
	// Test case 1: Normal scenario with positive numbers
	arr1 := []int{10, 200, 30, 40, 50}
	expectedMax1 := 200
	actualMax1 := findMax(arr1, len(arr1))
	if actualMax1 != expectedMax1 {
		t.Error("Test Case 1 Failed: Array: ", arr1, " Expected: ", expectedMax1, " Got: ", actualMax1)
	} else {
		t.Log("Test Case 1 Passed")
	}

	// Test case 2: Scenario with negative numbers
	arr2 := []int{-10, -200, -30, -40, -50}
	expectedMax2 := -10
	actualMax2 := findMax(arr2, len(arr2))
	if actualMax2 != expectedMax2 {
		t.Error("Test Case 2 Failed: Array: ", arr2, " Expected: ", expectedMax2, " Got: ", actualMax2)
	} else {
		t.Log("Test Case 2 Passed")
	}

	// Test case 3: Scenario with same numbers
	arr3 := []int{5, 5, 5, 5, 5}
	expectedMax3 := 5
	actualMax3 := findMax(arr3, len(arr3))
	if actualMax3 != expectedMax3 {
		t.Error("Test Case 3 Failed: Array: ", arr3, " Expected: ", expectedMax3, " Got: ", actualMax3)
	} else {
		t.Log("Test Case 3 Passed")
	}

	// Test case 4: Scenario with empty array
	arr4 := []int{}
	expectedMax4 := 0
	actualMax4 := findMax(arr4, len(arr4))
	if actualMax4 != expectedMax4 {
		t.Error("Test Case 4 Failed: Array: ", arr4, " Expected: ", expectedMax4, " Got: ", actualMax4)
	} else {
		t.Log("Test Case 4 Passed")
	}
}

func findMax(arr []int, n int) int {
	if n == 0 {
		return 0
	}

	max := arr[0]
	for i := 1; i < n; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	return max
}
