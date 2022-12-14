package main

import (
	"fmt"
	"sort"
)

func swords(k int, ages []int, sink func(v int)) {
	if len(ages) == 0 {
		panic("empty input")
	}
	indices := make([]int, len(ages))
	for i := range indices {
		indices[i] = i
	}
	sort.Slice(indices, func(i, j int) bool { return ages[indices[i]] > ages[indices[j]] })
	for i := len(indices) - 1; i > 0; i-- {
		if ages[indices[(i-1)/2]]-ages[indices[i]] < k {
			sink(-1)
			return
		}
	}
	ages[indices[0]] = 0
	for i := 1; i < len(ages); i++ {
		ages[indices[i]] = indices[(i-1)/2] + 1
	}
	for _, v := range ages {
		sink(v)
	}
}

func main() {
	var n, k int
	fmt.Scan(&n)
	fmt.Scan(&k)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	swords(k, a, func(v int) { fmt.Print(v, " ") })
}
