package main

import (
	"fmt"
	"sort"
)

type sword struct {
	index int
	age   int
}

func main() {
	var n, k int // n = sword amount, k = min age to copy
	fmt.Scanf("%v%v", &n, &k)
	a := make([]int, n) // a = sword ages
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	//a := []int{2, 10, 6, 0, 5, 2}
	//k := 3

	fmt.Println(swords(a, k))
}

func swords(a []int, k int) []int {
	//Creating sword slice
	s := make([]sword, len(a)) // s = sword slice
	for i := 0; i < len(a); i++ {
		s[i] = sword{index: i, age: a[i]}
	}

	//Sorting swords in decending order
	sort.Slice(s, func(i, j int) bool { return s[i].age > s[j].age })

	//Checking if each parent age is >= than k. If age is < k, return -1
	cp, lp := 0, (len(s)-2)/2 // cp - current parent, lp - last parent
	for ; cp < lp; cp++ {
		if s[cp].age-s[l(cp)].age < k || s[cp].age-s[r(cp)].age < k {
			return []int{-1}
		}
	}
	//handling last parent
	if cp < len(s)-1 {
		if s[cp].age-s[l(cp)].age < k || r(cp) < len(s) && s[cp].age-s[r(cp)].age < k {
			return []int{-1}
		}
	}

	//Finding parent of each sword
	parents := make([]int, len(a))
	parents[s[0].index] = 0 //setting 0 to location of oldest sword
	for i := 1; i < len(a); i++ {
		parents[s[i].index] = s[p(i)].index + 1 // setting the index of the parent in place of each sword
	}

	return parents
}

// heap left child
func l(i int) int {
	return 2*i + 1
}

// heap right child
func r(i int) int {
	return 2*i + 2
}

// heap parent
func p(i int) int {
	return (i - 1) / 2
}
