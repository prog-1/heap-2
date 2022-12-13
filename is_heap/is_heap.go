package main

func isHeap(s []int) bool {
	if len(s) < 2 {
		return true
	}
	for i, p := range s[:len(s)/2-1] {
		if s[2*i+1] < p || s[2*i+2] < p {
			return false
		}
	}
	if s[len(s)-1] < s[len(s)/2-1] {
		return false
	}
	if len(s)%2 != 0 && s[len(s)-2] < s[len(s)/2-1] {
		return false
	}

	return true
}
