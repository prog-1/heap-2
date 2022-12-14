package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, k int
	fmt.Fscan(r, &n, &k)
	ages := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &ages[i])
	}

	// k := 3
	// ages := []int{10, 1, 1, 1}

	// fmt.Fprint(os.Stdout, solve(k, ages))
	fmt.Println(solve(k, ages))
}

type sword struct {
	indx, age int
}

func solve(k int, ages []int) (parents []int) {
	swords := make([]sword, len(ages))
	for i := 0; i < len(ages); i++ {
		swords[i] = sword{indx: i, age: ages[i]}
	}
	sort.Slice(swords, func(i, j int) bool { return swords[i].age > swords[j].age })
	fmt.Println(swords)
	for i := len(ages) - 1; i > 0; i-- {

		if swords[(i-1)/2].age-swords[i].age < k {
			return []int{-1}
		}
	}
	if (0 < len(swords)-1) && (swords[0].age-swords[1].age < k || (2 < len(swords) && swords[0].age-swords[2].age < k)) {
		return []int{-1}
	}
	res := make([]int, len(ages))
	for i := 1; i < len(ages); i++ {
		res[swords[i].indx] = swords[((i-1)/2)].indx + 1
	}
	return res
}
