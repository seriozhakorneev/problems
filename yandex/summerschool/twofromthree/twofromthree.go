package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	nums := stdIn()
	fmt.Fprintln(os.Stdout, twoFromThree(nums))
}

func twoFromThree(nums [][]int) string {
	set := make(map[int][]int, 0)
	var result []int

	for arrCount, arr := range nums {
		for _, el := range arr {
			if v, ok := set[el]; ok {
				if v[0] < arrCount && v[1] != 1 {
					result = append(result, el)
					set[el][1]++
					set[el][0]++
				}
			} else {
				set[el] = []int{arrCount, 0}
			}
		}
	}

	sort.Ints(result)
	return toString(result)
}

func toString(a []int) (s string) {
	if len(a) < 1 {
		return
	}
	for _, el := range a {
		s += strconv.Itoa(el) + " "
	}
	return
}

func stdIn() [][]int {
	full := make([][]int, 0)
	for count := 3; count > 0; count-- {
		var n int
		fmt.Fscan(os.Stdin, &n)

		nums := make([]int, n)
		fmt.Fscan(os.Stdin, packAddrs(nums)...)
		full = append(full, nums)
	}

	return full
}

func packAddrs(n []int) []interface{} {
	p := make([]interface{}, len(n))
	for i := range n {
		p[i] = &n[i]
	}
	return p
}
