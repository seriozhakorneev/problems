package main

import (
	"fmt"
	"os"
)

/*
№	Вердикт	Ресурсы	Баллы
1	ok	3ms / 516.00Kb	-
2	ok	3ms / 520.00Kb	-
3	ok	3ms / 516.00Kb	-
4	ok	2ms / 512.00Kb	-
5	ok	3ms / 512.00Kb	-
6	ok	4ms / 516.00Kb	-
7	ok	4ms / 516.00Kb	-
8	ok	6ms / 516.00Kb	-
9	ok	14ms / 576.00Kb	-
10	ok	40ms / 968.00Kb	-
11	ok	123ms / 2.16Mb	-
12	ok	129ms / 2.16Mb	-
13	ok	454ms / 4.66Mb	-
14	ok	1.387s / 10.14Mb -
15	ok	1.454s / 9.63Mb	-
*/

func main() {
	nums, indexes := stdIn()
	countOfPositiveFast(nums, indexes)
}

func countOfPositiveFast(nums []int, indexes [][]int) {
	posCount := make([]int, len(nums)+1)
	posCount[0] = 0
	for i, num := range nums {
		posCount[i+1] = posCount[i]
		if num >= 1 {
			posCount[i+1]++
		}
	}

	for _, index := range indexes {
		fmt.Fprintln(
			os.Stdout,
			posCount[index[1]]-posCount[index[0]-1],
		)
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
