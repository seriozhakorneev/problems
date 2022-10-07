package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	var n int
	fmt.Fscan(os.Stdin, &n)

	ages := make([]int, n)
	fmt.Fscan(os.Stdin, packAddrsF(ages)...)

	sort.Ints(ages)

	sum := 0
	for left, right, i := 0, 0, 0; i < n; i++ {

		for left < n && float64(ages[left]) <= (float64(ages[i])*0.5)+7 {
			left++
		}
		for right < n && ages[right] <= ages[i] {
			right++
		}

		if left < right {
			sum += right - left - 1
		}
	}
	fmt.Fprintln(os.Stdout, sum)
}

func packAddrsF(n []int) []interface{} {
	p := make([]interface{}, len(n))
	for i := range n {
		p[i] = &n[i]
	}
	return p
}
