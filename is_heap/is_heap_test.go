package main

import "testing"

func TestIsHeap(t *testing.T) {
	for _, tc := range []struct {
		name string
		s    []int
		want bool
	}{
		{"nil", nil, true},
		{"1", []int{1, 3, 9, 7, 5}, true},
		{"2", []int{1, 10, 9, 7, 5}, false},
	} {
		if got := isHeap(tc.s); got != tc.want {
			t.Errorf("got = %v, want = %v", got, tc.want)
		}
	}
}
