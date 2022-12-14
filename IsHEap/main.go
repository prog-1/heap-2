package main

func isHeap(s []int) bool {
	if len(s) == 1 {
		return true
	}
	for i := len(s) - 1; i > 0; i-- {
		if s[i] < s[(i-1)/2] {
			return false
		}
	}
	return true
}
