package main

import (
	"reflect"
	"testing"
)

func TestSwords(t *testing.T) {
	for _, tc := range []struct {
		name string
		k    int
		a    []int
		want []int
	}{
		{"case-1", 3, []int{2, 10, 6, 0, 5, 2}, []int{3, 0, 2, 5, 2, 3}},
		{"case-2", 3, []int{10, 1, 1, 1}, []int{-1}},
		{"case-3", 1, []int{2, 2, 2, 3, 4, 2, 2, 4, 5, 3, 3, 3, 2, 2, 2}, []int{11, 4, 4, 5, 9, 10, 10, 9, 0, 5, 8, 8, 11, 12, 12}},
		{"case-4", 9, []int{1}, []int{0}},
		{"case-5", 0, []int{1, 1, 1, 1, 1, 1}, []int{0, 1, 1, 2, 2, 3}},
		{"case-6", 0, []int{1, 1, 1, 2, 1, 1}, []int{4, 4, 1, 0, 1, 2}},
		{"case-7", 1, []int{0, 1, 2, 3, 4, 5}, []int{4, 5, 5, 6, 6, 0}},
		{"case-8", 2, []int{9, 8, 7, 11, 5, 2}, []int{4, 4, 1, 0, 1, 2}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var got []int
			swords(tc.k, tc.a, func(v int) { got = append(got, v) })
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
