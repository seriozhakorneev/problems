package main

/*

import (
	"fmt"
	"os"
)

//№	Вердикт	Ресурсы	Баллы
//1	ok	2ms / 516.00Kb	-
//2	ok	2ms / 516.00Kb	-
//3	ok	2ms / 516.00Kb	-
//4	ok	2ms / 512.00Kb	-
//5	ok	6ms / 512.00Kb	-
//6	ok	3ms / 516.00Kb	-
//7	ok	3ms / 512.00Kb	-
//8	ok	6ms / 512.00Kb	-
//9	ok	14ms / 540.00Kb	-
//10	ok	44ms / 888.00Kb	-
//11	ok	161ms / 2.08Mb	-
//12	ok	175ms / 2.10Mb	-
//13	ok	0.834s / 4.72Mb	-
//14	time-limit-exceeded	2.015s / 8.87Mb	-


func main() {
	nums, indexes := stdIn()
	countOfPositiveSlow(nums, indexes)
}

func countOfPositiveSlow(nums []int, indexes [][]int) {
	for _, indexRange := range indexes {
		positive := 0
		for _, num := range nums[indexRange[0]-1 : indexRange[1]] {
			if num > 0 {
				positive++
			}
		}
		fmt.Fprintln(os.Stdout, positive)
	}
}

func stdIn() (nums []int, indexes [][]int) {
	var n int
	fmt.Fscan(os.Stdin, &n)
	nums = make([]int, n)
	fmt.Fscan(os.Stdin, packAddrs(nums)...)

	var q int
	fmt.Fscan(os.Stdin, &q)
	indexes = make([][]int, q)
	for i := 0; i < q; i++ {
		slice := make([]int, 2)
		fmt.Fscan(os.Stdin, packAddrs(slice)...)
		indexes[i] = slice
	}

	return nums, indexes
}

func packAddrs(n []int) []interface{} {
	p := make([]interface{}, len(n))
	for i := range n {
		p[i] = &n[i]
	}
	return p
}
*/
