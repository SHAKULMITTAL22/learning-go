package BubbleSort

import (
	"reflect"
	"testing"
)

func TestBubbleSort_dcbfe1ec71(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}

	for _, c := range cases {
		got := bubbleSort(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("bubbleSort(%v) == %v, want %v", c.in, got, c.want)
		} else {
			t.Logf("Success: bubbleSort(%v) == %v", c.in, got)
		}
	}
}
