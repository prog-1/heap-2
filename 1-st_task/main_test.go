package main

import "testing"

func TestIsHeap(t *testing.T) {
	for _, tc := range []struct {
		input []int
		want  bool
	}{
		{[]int{1}, true},
		{[]int{1, 2}, true},
		{[]int{1, 0}, false},
		{[]int{1, 3, 3}, true},
		{[]int{1, 3, 2}, true},
		{[]int{1, 3, 3, 2}, false},
		{[]int{1, 3, 2, 3, 3, 2}, true},
		{[]int{1, 3, 2, 3, 2}, false},
	} {
		if got := isHeap(tc.input); got != tc.want {
			t.Errorf("isHeap(%v)=%v, want = %v", tc.input, got, tc.want)
		}
	}
}
