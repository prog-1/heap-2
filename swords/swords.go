package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {
	fmt.Println(solve(processInput(bufio.NewReader(os.Stdin))))
}

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

func solve(k int, ages []int) (parents []int) {
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
			return []int{-1}
		}
	}

	// How to output index of the parent of each sword?
	// Take index of current sword
	// Find it in sorted index array
	// Get parent of this element

	find := func(s []int, v int) int {
		for i, el := range s {
			if el == v {
				return i
			}
		}
		return -1
	}

	parents = make([]int, len(ages))
	for i := 0; i < len(parents); i++ {
		j := find(indices, i) // getting index of index of sword
		iip := (j - 1) / 2    // index of index of parent
		if iip == j {         // child is its own parent
			parents[i] = 0
		} else {
			ip := indices[iip]
			parents[i] = ip + 1
		}
	}

	return parents
}
