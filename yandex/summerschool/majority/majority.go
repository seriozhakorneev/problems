package main

import (
	"fmt"
	"os"
)

func main() {
	n, nums := stdIn()
	fmt.Fprintln(os.Stdout, majority(n, nums))
}

func majorityMap(n int, nums []int) (target int) {
	set := make(map[int]int)

	for i := 0; i < n; i++ {
		set[nums[i]]++
	}

	for k, v := range set {
		if v > (n / 2) {
			target = k
		}
	}
	return
}

func majority(n int, nums []int) int {
	target := nums[0]
	count := 1
	for i := 1; i < n; i++ {
		if nums[i] == target {
			count++
		} else {
			count--
			if count == 0 {
				target = nums[i]
				count = 1
			}
		}
	}
	return target
}

func stdIn() (int, []int) {
	var n int
	fmt.Fscan(os.Stdin, &n)

	nums := make([]int, n)
	fmt.Fscan(os.Stdin, packAddrs(nums)...)

	return n, nums
}

func packAddrs(n []int) []interface{} {
	p := make([]interface{}, len(n))
	for i := range n {
		p[i] = &n[i]
	}
	return p
}
