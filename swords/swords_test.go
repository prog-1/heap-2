package main

import (
	"reflect"
	"testing"
)

func TestSwords(t *testing.T) {
	for _, tc := range []struct {
		name string
		a    []int
		k    int
		want []int
	}{
		{"example 1", []int{2, 10, 6, 0, 5, 2}, 3, []int{3, 0, 2, 5, 2, 3}},
		{"example 2", []int{10, 1, 1, 1}, 3, []int{-1}},
		{"same ages 1", []int{2, 2, 2, 3, 4, 2, 2, 4, 5, 3, 3, 3, 2, 2, 2}, 1, []int{12, 4, 11, 8, 9, 11, 12, 9, 0, 8, 5, 5, 4, 10, 10}},
		{"same ages 2", []int{2,2}, 1, []int{-1}},
		{"same ages 3", []int{2,2}, 0, []int{0,1}},
		{"1 element", []int{1}, 0, []int{0}},
		{"empty", []int{}, 0, []int{-1}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := swords(tc.a, tc.k)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
