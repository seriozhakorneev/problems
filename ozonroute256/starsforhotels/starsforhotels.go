package main

import (
	"fmt"
	"os"
	"sort"
)

// TODO: partially solved
func main() {

	var t int
	fmt.Fscan(os.Stdin, &t)
	for ; t > 0; t-- {

		var n int
		fmt.Fscan(os.Stdin, &n)

		votes := make([]int, n)
		fmt.Fscan(os.Stdin, packAddrs(votes)...)

		votesCopy := make([]int, n)
		copy(votesCopy, votes)
		sort.Ints(votesCopy)

		// votes: stars
		stars := make(map[int]int)
		set := make(map[int]bool)

		fives, star := ((n-15)/5)+1, 5
		for i, count := n-1, fives; i >= 0; i-- {
			stars[votesCopy[i]] = star

			if !set[votesCopy[i]] {
				set[votesCopy[i]] = true
			}

			if star != 1 {
				count--
			}
			if count == 0 {
				star--
				fives++
				count = fives
			}
		}

		println()
		if len(set) < 5 {
			for range votes {
				fmt.Fprint(os.Stdout, -1, " ")
			}
		} else {
			for _, vote := range votes {
				fmt.Fprint(os.Stdout, stars[vote], " ")
			}
		}
	}
}

func packAddrs(n []int) []interface{} {
	p := make([]interface{}, len(n))
	for i := range n {
		p[i] = &n[i]
	}
	return p
}
