package main

func isHeap(s []int) bool {
	for i := 0; i < len(s); i++ {
		if left(i) < len(s) && s[left(i)] < s[i] {
			return false
		}
		if right(i) < len(s) && s[right(i)] < s[i] {
			return false
		}
	}
	return true
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func main() {

}
