package main

import (
	"fmt"
	"io"
	"sort"
)

func processInput(r io.Reader) (k int, ages []int) {
	// n - sword count
	// k - min age for cloning
	var n int
	fmt.Fscan(r, &n, &k)
	ages = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &ages[i])
	}
	return k, ages
}

func solve(k int, ages []int) (parents []int, ok bool) {
	indices := make([]int, len(ages))
	for i := range indices {
		indices[i] = i
	}

	sort.Slice(indices, func(i, j int) bool {
		return ages[indices[i]] > ages[indices[j]]
	})

	for childIndex := len(indices) - 1; childIndex > 0; childIndex-- {
		child := ages[indices[childIndex]]
		parent := ages[indices[(childIndex-1)/2]]
		if parent-child < k {
			return nil, false
		}
	}

	parents = make([]int, len(ages))
	parents[indices[0]] = 0
	for i := 1; i < len(indices); i++ {
		iip := (i - 1) / 2 // index of index of parent
		ip := indices[iip]
		parents[indices[i]] = ip + 1
	}

	return parents, true
}
