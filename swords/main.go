package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func isValid(ages, indices []int, maxAge int) []int {
	for i := len(indices) - 1; i > 0; i-- {
		p := (i - 1) / 2
		if ages[indices[p]]-ages[indices[i]] >= maxAge {
			ages[indices[i]] = indices[p] + 1
		} else {
			return []int{-1}
		}
	}
	ages[indices[0]] = 0
	return ages
}

func main() {
	swordsin := bufio.NewReader(os.Stdin)
	var n, a int
	var ages, indices []int
	fmt.Scan(&n, &a)
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(swordsin, &x)
		ages = append(ages, x)
		indices = append(indices, i)
	}
	sort.Slice(indices, func(i, j int) bool { return ages[indices[i]] > ages[indices[j]] })
	fmt.Println(isValid(ages, indices, a))
}
