package main

import "testing"

func TestIsHeap(t *testing.T) {
	for _, tc := range []struct {
		name string
		s    []int
		want bool
	}{
		{"case-1", nil, true},
		{"case-2", []int{1}, true},
		{"case-3", []int{1, 2}, true},
		{"case-4", []int{1, 2, 3}, true},
		{"case-5", []int{1, 2, 3, 4, 5}, true},
		{"case-6", []int{1, 2, 3, 4, 5, 2}, false},
		{"case-7", []int{2, 1}, false},
		{"case-8", []int{1, 1, 1, 1, 1}, true},
		{"case-9", []int{2, 2, 1}, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := isHeap(tc.s); got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
