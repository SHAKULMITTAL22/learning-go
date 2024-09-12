package RadixSort

import (
	"testing"
	"reflect"
)

func radixsort(arr []int, n int) {
	max := getMax(arr, n)
	for exp := 1; max/exp > 0; exp *= 10 {
		countSort(arr, n, exp)
	}
}

func getMax(arr []int, n int) int {
	max := arr[0]
	for i := 1; i < n; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func countSort(arr []int, n int, exp int) {
	output := make([]int, n)
	count := make([]int, 10)

	for i := 0; i < n; i++ {
		count[(arr[i]/exp)%10]++
	}

	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	for i := n - 1; i >= 0; i-- {
		output[count[(arr[i]/exp)%10]-1] = arr[i]
		count[(arr[i]/exp)%10]--
	}

	for i := 0; i < n; i++ {
		arr[i] = output[i]
	}
}

func TestRadixsort_c800f73b57(t *testing.T) {
	testCases := []struct {
		input []int
		n     int
		want  []int
	}{
		{
			input: []int{170, 45, 75, 90, 802, 24, 2, 66},
			n:     8,
			want:  []int{2, 24, 45, 66, 75, 90, 170, 802},
		},
		{
			input: []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			n:     9,
			want:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, tc := range testCases {
		radixsort(tc.input, tc.n)
		if !reflect.DeepEqual(tc.input, tc.want) {
			t.Errorf("radixsort(%v, %v): expected %v, got %v", tc.input, tc.n, tc.want, tc.input)
		} else {
			t.Logf("Test passed for: radixsort(%v, %v)", tc.input, tc.n)
		}
	}
}
