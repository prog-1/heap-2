package main

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	for _, tc := range []struct {
		a       int
		ages    []int
		indices []int
		want    []int
	}{
		{3, []int{2, 10, 6, 0, 5, 2}, []int{1, 2, 4, 0, 5, 3}, []int{3, 0, 2, 5, 2, 3}},
		{5, []int{2, 10, 6, 0, 5, 2}, []int{1, 2, 4, 0, 5, 3}, []int{-1}},
		{3, []int{10, 1, 1, 1}, []int{0, 1, 2, 3}, []int{-1}},
		{2, []int{15, 3, 1, 12, 12, 0, 10, 11, 14, 25}, []int{9, 0, 8, 3, 4, 7, 6, 1, 2, 5}, []int{10, 4, 4, 1, 1, 5, 9, 9, 10, 0}},
		{5, []int{15, 3, 1, 12, 12, 0, 10, 11, 14, 25}, []int{9, 0, 8, 3, 4, 7, 6, 1, 2, 5}, []int{-1}},
	} {
		got := isValid(tc.ages, tc.indices, tc.a)
		if !equal(got, tc.want) {
			t.Errorf("isValid(%v) = %v, want = %v", tc.ages, got, tc.want)
		}
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
