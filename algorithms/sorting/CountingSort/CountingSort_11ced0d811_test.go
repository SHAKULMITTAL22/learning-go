package CountingSort

import (
	"reflect"
	"testing"
)

func TestCountingSort_11ced0d811(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "empty array",
			arr:  []int{},
			want: []int{},
		},
		{
			name: "single element array",
			arr:  []int{1},
			want: []int{1},
		},
		{
			name: "multiple elements array",
			arr:  []int{3, 2, 5, 1, 0, 4},
			want: []int{0, 1, 2, 3, 4, 5},
		},
		{
			name: "negative elements array",
			arr:  []int{-3, -2, -5, -1, 0, -4},
			want: []int{-5, -4, -3, -2, -1, 0},
		},
		{
			name: "duplicate elements array",
			arr:  []int{3, 2, 5, 1, 0, 4, 3, 2},
			want: []int{0, 1, 2, 2, 3, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countingSort(tt.arr)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countingSort(%v) = %v, want %v", tt.arr, got, tt.want)
			}
		})
	}
}
