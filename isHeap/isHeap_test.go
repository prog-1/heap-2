package main

import (
	"testing"
)

func TestIsHeap(t *testing.T) {
	for _, tc := range []struct {
		name string
		s    []int
		want bool
	}{
		{"1", []int{1, 2, 4, 9, 3, 8, 7, 10, 16, 14}, true},
		{"2", []int{1, 2, 4, 20, 3, 8, 7, 10, 16, 14}, false},
		{"same elements", []int{1, 1, 1, 2, 2, 2, 2}, true},
		{"1 element", []int{1}, true},
		{"empty", []int{}, true},
		{"nil", nil, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := isMinHeap(tc.s)
			if got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}

func main() {
	testing.Main(
		/* matchString */ func(a, b string) (bool, error) { return a == b, nil },
		/* tests */ []testing.InternalTest{
			{Name: "Test is Min Heap", F: TestIsHeap},
		},
		/* benchmarks */ nil /* examples */, nil)
}
