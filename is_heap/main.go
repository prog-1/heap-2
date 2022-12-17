package main

import (
	"fmt"
)

func isHeap(s []int) bool {
	var res bool
	if len(s) == 1 || s == nil {
		return true
	}
	for i := len(s) - 1; i > 0; i-- {
		if s[(i-1)/2] <= s[i] {
			res = true
		} else {
			return false
		}
	}

	return res
}

func main() {
	// s := []int{1, 2, 5, 6, 4}
	// s1 := []int{1, 5, 2, 6, 4}
	// s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	s := []int{1, 2, 3, 4, 5, 12, 7, 8, 9, 10, 11, 6, 13, 14, 15}
	s1 := []int{1, 11, 3, 4, 13, 6, 12, 8, 9, 10, 2, 7, 5, 14, 15}
	fmt.Println(isHeap(s))
	fmt.Println(isHeap(s1))
}
