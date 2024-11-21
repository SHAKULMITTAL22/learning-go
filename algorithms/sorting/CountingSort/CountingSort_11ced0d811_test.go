package CountingSort

import (
	"reflect"
	"testing"
)

func TestCountingSort_11ced0d811(t *testing.T) {
	getCountArrayLength := func(arr []int) int {
		max := 0
		for _, value := range arr {
			if max < value {
				max = value
			}
		}
		return max + 1
	}

	countingSort := func(arr []int) []int {
		k := getCountArrayLength(arr)
		// Create array with a length of k
		count := make([]int, k)

		// Store count of each character
		for i := 0; i < len(arr); i++ {
			count[arr[i]] += 1
		}

		for i, j := 0, 0; i < k; i++ {
			for {
				if count[i] > 0 {
					arr[j] = i
					j += 1
					count[i] -= 1
					continue
				}
				break
			}
		}

		return arr
	}

	// Test case 1: normal case
	arr1 := []int{9, 4, 3, 2, 5, 1, 7, 6, 8, 0}
	expected1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	result1 := countingSort(arr1)
	if !reflect.DeepEqual(result1, expected1) {
		t.Error("Test case 1 failed: input was", arr1, ", expected", expected1, ", but got", result1)
	} else {
		t.Log("Test case 1 passed")
	}

	// Test case 2: edge case with all elements being the same
	arr2 := []int{1, 1, 1, 1, 1}
	expected2 := []int{1, 1, 1, 1, 1}
	result2 := countingSort(arr2)
	if !reflect.DeepEqual(result2, expected2) {
		t.Error("Test case 2 failed: input was", arr2, ", expected", expected2, ", but got", result2)
	} else {
		t.Log("Test case 2 passed")
	}
}
