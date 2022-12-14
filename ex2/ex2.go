package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Heap []int

func isMaxHeap(h []int, k int) bool {
	fmt.Println(h)
	for i := len(h) - 1; i > 0; i-- {
		fmt.Println(h[(i-1)/2], h[i], h[(i-1)/2]-h[i] <= k)
		if h[(i-1)/2]-h[i] < k {
			return false
		}
	}
	return true
}

func s(swords, sindex []int) {
	sort.Slice(sindex, func(i, j int) bool { return swords[sindex[i]] > swords[sindex[j]] })
}

func find(copys, sindex []int) []int {
	for i := len(sindex) - 1; i > 0; i-- {
		copys[sindex[i]] = sindex[(i-1)/2] + 1
	}
	copys[sindex[0]] = 0
	return copys
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, k int
	var swords, copys, sindex []int
	fmt.Println("First two")
	fmt.Scan(&n, &k)
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(in, &x)
		swords = append(swords, x)
		copys = append(copys, x)
		sindex = append(sindex, i)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(copys)))
	if !isMaxHeap(copys, k) {
		fmt.Println(-1)
		return
	} else {
		s(swords, sindex)
		out := find(copys, sindex)
		fmt.Println(swords, sindex, out)
		for _, v := range out {
			fmt.Print(v, " ")
		}
	}
}
