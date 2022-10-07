package main

import (
	"fmt"
	"os"
)

/*
№	Вердикт	Ресурсы	Баллы
1	ok	3ms / 516.00Kb	-
2	ok	3ms / 508.00Kb	-
3	ok	6ms / 512.00Kb	-
4	ok	3ms / 512.00Kb	-
5	ok	3ms / 516.00Kb	-
6	ok	3ms / 512.00Kb	-
7	ok	4ms / 512.00Kb	-
8	ok	6ms / 520.00Kb	-
9	ok	15ms / 528.00Kb	-
10	ok	42ms / 896.00Kb	-
11	ok	133ms / 1.54Mb	-
12	ok	141ms / 1.59Mb	-
13	ok	464ms / 4.56Mb	-
14	ok	1.472s / 4.98Mb	-
15	ok	1.552s / 4.98Mb	-
*/

func main() {
	countOfPositiveFastest()
}

func countOfPositiveFastest() {

	var n int
	fmt.Fscan(os.Stdin, &n)
	nums := make([]int, n)
	fmt.Fscan(os.Stdin, packAddrsF(nums)...)

	posCount := make([]int, len(nums)+1)
	posCount[0] = 0
	for i, num := range nums {
		posCount[i+1] = posCount[i]
		if num >= 1 {
			posCount[i+1]++
		}
	}

	var q int
	fmt.Fscan(os.Stdin, &q)
	for i := 0; i < q; i++ {
		var first, second int
		fmt.Fscan(os.Stdin, &first, &second)
		fmt.Fprintln(os.Stdout, posCount[second]-posCount[first-1])
	}
}

func packAddrsF(n []int) []interface{} {
	p := make([]interface{}, len(n))
	for i := range n {
		p[i] = &n[i]
	}
	return p
}
