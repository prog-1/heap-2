package main

// func main() {
// 	s := []int{1, 2, 4, 9, 3, 8, 7, 10, 16, 14}
// 	fmt.Println(isMinHeap(s))
// }

// function that verifies is given slice the minHeap
func isMinHeap(s []int) bool {
	n, lp := 0, (len(s)-2)/2 //last parent = parent of the last node
	for ; n < lp; n++ {
		if s[n] > s[l(n)] || s[n] > s[r(n)] {
			return false
		}
	}
	//for last parent
	if n < len(s)-1 {
		if s[n] > s[l(n)] || r(n) < len(s) && s[n] > s[r(n)] {
			return false
		}
	}
	return true
}

// heap left child
func l(i int) int {
	return 2*i + 1
}

// heap right child
func r(i int) int {
	return 2*i + 2
}
