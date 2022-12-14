package main

import "testing"

func TestIsHeap(t *testing.T) {
	for _, tc := range []struct {
		name string
		s    []int
		want bool
	}{
		{"nil", nil, true},
		{"1", []int{1, 2, 9, 90, 5}, true},
		{"2", []int{1, 123, 9, 4, 5}, false},
		{"3", []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, true},
		{"4", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, false},
	} {
		if got := isHeap(tc.s); got != tc.want {
			t.Errorf("got = %v, want = %v", got, tc.want)
		}
	}
}
