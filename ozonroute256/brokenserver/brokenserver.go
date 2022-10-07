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

		a := make([]int, n)
		fmt.Fscan(os.Stdin, packAddrs(a)...)

		fmt.Fprintln(os.Stdout, maxRequests(n, a))
	}
}

func maxRequests(n int, a []int) (max int) {
	for i := 0; i < n; i++ {
		tmp := 0
		for j := i; j < n+1; j++ {

			if max < len(a[i:j]) {
				max = len(a[i:j])
			}
			if j == n {
				break
			}
			if tmp == 0 && a[j] != a[i] {
				tmp = a[j]
				continue
			}
			if a[j] != a[i] && a[j] != tmp {
				break
			}
		}
	}
	return max
}

func packAddrs(n []int) []interface{} {
	p := make([]interface{}, len(n))
	for i := range n {
		p[i] = &n[i]
	}
	return p
}
