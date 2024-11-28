package LinearSearch

import "testing"

/*
ROOST_METHOD_HASH=linearSearch_076b2ca8c0
ROOST_METHOD_SIG_HASH=linearSearch_12fac2e721
*/
func Testlinearsearch577(t *testing.T) {
	t.Run("Scenario 1: Test linearSearch with a Non-Existent Element", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}
		query := 99
		expected := -1

		result := linearSearch(arr, query)
		if result != expected {
			t.Errorf("Expected %d but got %d", expected, result)
		} else {
			t.Logf("Success: Element %d not found in array, returned %d as expected", query, result)
		}
	})

	t.Run("Scenario 2: Test linearSearch with an Empty Array", func(t *testing.T) {
		arr := []int{}
		query := 5
		expected := -1

		result := linearSearch(arr, query)
		if result != expected {
			t.Errorf("Expected %d but got %d", expected, result)
		} else {
			t.Logf("Success: Element %d not found in empty array, returned %d as expected", query, result)
		}
	})

	t.Run("Scenario 3: Test linearSearch with Multiple Identical Elements", func(t *testing.T) {
		arr := []int{1, 2, 2, 4, 2, 5}
		query := 2
		expected := 1

		result := linearSearch(arr, query)
		if result != expected {
			t.Errorf("Expected %d but got %d", expected, result)
		} else {
			t.Logf("Success: First occurrence of element %d found at index %d", query, result)
		}
	})

	t.Run("Scenario 4: Test linearSearch with Sorted Array", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}
		query := 3
		expected := 2

		result := linearSearch(arr, query)
		if result != expected {
			t.Errorf("Expected %d but got %d", expected, result)
		} else {
			t.Logf("Success: Element %d found in sorted array at index %d", query, result)
		}
	})

	t.Run("Scenario 5: Test linearSearch with Negative Numbers", func(t *testing.T) {
		arr := []int{-5, -3, -1, 0, 2, 4, 6}
		query := -1
		expected := 2

		result := linearSearch(arr, query)
		if result != expected {
			t.Errorf("Expected %d but got %d", expected, result)
		} else {
			t.Logf("Success: Element %d found among negative numbers at index %d", query, result)
		}
	})

}
