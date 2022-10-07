package main

import (
	"fmt"
	"os"
)

func main() {
	var t int
	fmt.Fscan(os.Stdin, &t)
	for ; t > 0; t-- {

		var n int
		fmt.Fscan(os.Stdin, &n)

		list := make([]int, n)
		fmt.Fscan(os.Stdin, &list[n-2], &list[n-3], &list[n-1])

		triplets := make(map[int][]int)
		for i := 0; i < n-1; i++ {
			var e, a, b int
			fmt.Fscan(os.Stdin, &e, &a, &b)
			triplets[e] = []int{a, b}
		}

		for ; n-3 > 0; n-- {
			if list[n-2] == triplets[list[n-3]][0] {
				list[n-4] = triplets[list[n-3]][1]
			} else if list[n-2] == triplets[list[n-3]][1] {
				list[n-4] = triplets[list[n-3]][0]
			}
		}

		for f, h := 0, len(list)/2; f < len(list)/2; f, h = f+1, h+1 {
			fmt.Fprintln(os.Stdout, list[f], list[h])
		}
	}
}
