package main

func isHeap(s []int) bool {
	for i := 0; i < len(s); i++ {
		l := 2*i + 1
		r := 2*i + 2
		if l >= len(s) && r >= len(s) {
			for j := i; j < len(s); j++ {
				l := 2*j + 1
				r := 2*j + 2
				if l < len(s) || r < len(s) {
					return false
				}
			}
			return true
		} else if l >= 0 && l < len(s) && r >= 0 && r < len(s) && s[i] <= s[l] && s[i] <= s[r] {
			continue
		} else {
			return false
		}
	}
	return true
}

// func sowrds(n, k int, s []int) []int {
// 	var output []int
// 	s1 := s
// 	sort.Ints(s1)
// 	for i := range s1{

// 	}
// }
