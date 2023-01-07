package main

import (
	"sort"
	"testing"
)

func TestSwords(t *testing.T) {
	type Input struct {
		k    int
		ages []int
	}
	for _, tc := range []struct {
		input  Input
		wantOK bool
	}{
		{Input{3, []int{2, 10, 6, 0, 5, 2}}, true},
		{Input{1, []int{2, 2, 2, 3, 4, 2, 2, 4, 5, 3, 3, 3, 2, 2, 2}}, true},

		{Input{3, []int{2, 10, 6, 0, 5, 2}}, true},
		{Input{5, []int{2, 10, 6, 0, 5, 2}}, false},
		{Input{3, []int{10, 1, 1, 1}}, false},
		{Input{2, []int{15, 3, 1, 12, 12, 0, 10, 11, 14, 25}}, true},
		{Input{5, []int{15, 3, 1, 12, 12, 0, 10, 11, 14, 25}}, false},
		// (c) Lera
	} {
		// 1. Check whether function returns correct success indicator
		// 2. Check whether it is a valid max heap
		got, ok := solve(tc.input.k, tc.input.ages)
		if ok != tc.wantOK {
			t.Errorf("Incorrect success indicator: got = %v, want = %v", ok, tc.wantOK)
		}
		if ok {
			sort.Slice(got, func(i, j int) bool {
				return got[i] > got[j]
			})
			if !IsMaxHeap(got) {
				t.Errorf("Output slice does not form a valid max heap")
			}
		}
	}
}

func IsMaxHeap(s []int) bool {
	if len(s) < 2 {
		return true
	}
	for i := len(s)/2 - 2; i >= 0; i-- {
		p := s[i]
		l := s[i*2+1]
		r := s[i*2+2]
		if l > p || r > p {
			return false
		}
	}
	i := len(s)/2 - 1
	p := s[i]
	c1 := s[len(s)-1]
	if c1 > p {
		return false
	}
	if len(s)%2 != 0 {
		l := s[len(s)-2]
		if l > p {
			return false
		}
	}

	return true
}
