package main

import (
	"fmt"
	"os"
)

func main() {
	n, k, nums := stdIn()
	fmt.Fprintln(os.Stdout, isRepeat(n, k, nums))
}

func isRepeat(n, distance int, nums []int) string {
	set := make(map[int]int)
	for i := 0; i < n; i++ {
		if index, ok := set[nums[i]]; ok && (i-index) <= distance {
			return "YES"
		}
		set[nums[i]] = i
	}

	return "NO"
}

func stdIn() (int, int, []int) {
	var n, k int
	fmt.Fscan(os.Stdin, &n, &k)

	nums := make([]int, n)
	fmt.Fscan(os.Stdin, packAddrs(nums)...)

	return n, k, nums
}

func packAddrs(n []int) []interface{} {
	p := make([]interface{}, len(n))
	for i := range n {
		p[i] = &n[i]
	}
	return p
}
