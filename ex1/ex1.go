package main

import "fmt"

func isHeap(h []int) bool {
	for i := len(h) - 1; i >= 0; i-- {
		if h[i] < h[(i-1)/2] {
			return false
		}
	}
	return true
}

func main() {
	h := []int{1, 2, 3, 4}
	fmt.Println(isHeap(h))
}
