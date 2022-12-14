package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Heap []element

type element struct {
	index int
	age   int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, k int
	fmt.Fscan(reader, &n, &k)
	heap := make([]element, n)
	for i := 0; i < n; i++ {
		var tmp int
		fmt.Fscan(reader, &tmp)
		heap[i] = element{index: i, age: tmp}
	}
	sort.Slice(heap, func(i, j int) bool { return heap[i].age > heap[j].age })
	s := make([]int, n)
	solved := solve(heap, s, k)
	if !solved {
		fmt.Println(-1)
	} else {
		for _, v := range s {
			fmt.Print(v, " ")
		}
	}
}

func solve(heap Heap, s []int, dif int) bool {
	for i := 0; i < len(heap); i++ {
		left := 2*i + 1
		right := left + 1
		if left < len(heap) && heap[left].age > heap[i].age-dif {
			return false
		} else {
			if left < len(heap) {
				s[heap[left].index] = heap[i].index + 1
			}
		}
		if right < len(heap) && heap[right].age > heap[i].age-dif {
			return false
		} else {
			if right < len(heap) {
				s[heap[right].index] = heap[i].index + 1
			}
		}
	}
	return true
}
