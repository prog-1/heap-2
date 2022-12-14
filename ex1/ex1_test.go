package main

import "testing"

func TestIsHeap(t *testing.T) {
	for _, tc := range []struct {
		name string
		s    []int
		want bool
	}{
		{"1", nil, true},
		{"2", []int{1}, true},
		{"3", []int{1, 2}, true},
		{"4", []int{1, 2, 3}, true},
		{"5", []int{1, 2, 3, 4, 5}, true},
		{"6", []int{1, 2, 3, 4, 5, 2}, false},
		{"7", []int{2, 1}, false},
		{"8", []int{1, 1, 1, 1, 1}, true},
		{"9", []int{2, 2, 1}, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := isHeap(tc.s); got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
