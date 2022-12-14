package main

func isHeap(s []int) bool {
	if s == nil {
		return true
	}
	for i := len(s) - 1; i != 0; i-- {
		if s[(i-1)/2] > s[i] {
			return false
		}
	}
	return true
}
