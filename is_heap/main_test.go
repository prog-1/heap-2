package main

import "testing"

func TestIsHeap(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want bool
	}{
		{nil, true},
		{[]int{1}, true},
		{[]int{12, 12, 12, 12}, true},
		{[]int{1, 2, 3, 4, 5, 12, 7, 8, 9, 10, 11, 6, 13, 14, 15}, false},
		{[]int{1, 11, 3, 4, 13, 6, 12, 8, 9, 10, 2, 7, 5, 14, 15}, false},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, true},
		{[]int{1, 2, 5, 6, 4}, true},
		{[]int{1, 5, 2, 6, 4}, false},
	} {

		if got := isHeap(tc.s); got != tc.want {
			t.Errorf("Test(%v), got = %v, want = %v", tc.s, got, tc.want)
		}

	}
}
