// Test generated by RoostGPT for test math-go using AI Type Open AI and AI Model gpt-4

package SelectionSort

import (
	"reflect"
	"testing"
)

func TestSelectionSort_bdf5262e27(t *testing.T) {
	// Test case 1: Normal case
	arr1 := []int{5, 2, 6, 7, 2, 1, 0, 3}
	want1 := []int{0, 1, 2, 2, 3, 5, 6, 7}
	got1 := selectionSort(arr1)

	if !reflect.DeepEqual(got1, want1) {
		t.Errorf("selectionSort(%v) = %v; want %v", arr1, got1, want1)
	} else {
		t.Logf("selectionSort(%v) = %v; test case passed", arr1, got1)
	}

	// Test case 2: Edge case with an empty array
	arr2 := []int{}
	want2 := []int{}
	got2 := selectionSort(arr2)

	if !reflect.DeepEqual(got2, want2) {
		t.Errorf("selectionSort(%v) = %v; want %v", arr2, got2, want2)
	} else {
		t.Logf("selectionSort(%v) = %v; test case passed", arr2, got2)
	}

	// Test case 3: Case with negative numbers
	arr3 := []int{-5, -2, -6, -7, -2, -1, 0, 3}
	want3 := []int{-7, -6, -5, -2, -2, -1, 0, 3}
	got3 := selectionSort(arr3)

	if !reflect.DeepEqual(got3, want3) {
		t.Errorf("selectionSort(%v) = %v; want %v", arr3, got3, want3)
	} else {
		t.Logf("selectionSort(%v) = %v; test case passed", arr3, got3)
	}
}
