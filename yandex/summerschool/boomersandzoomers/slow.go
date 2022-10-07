package main

import (
	"fmt"
	"os"
	"sort"
)

/*
№	Вердикт	Ресурсы	Баллы
	1	ok	2ms / 516.00Kb	-
	2	ok	2ms / 512.00Kb	-
	3	ok	2ms / 516.00Kb	-
	4	ok	2ms / 516.00Kb	-
	5	ok	2ms / 520.00Kb	-
	6	ok	3ms / 512.00Kb	-
	7	ok	3ms / 512.00Kb	-
	8	ok	3ms / 520.00Kb	-
	9	ok	3ms / 512.00Kb	-
	10	ok	4ms / 516.00Kb	-
	11	ok	5ms / 516.00Kb	-
	12	ok	5ms / 520.00Kb	-
	13	ok	6ms / 512.00Kb	-
	14	ok	7ms / 512.00Kb	-
	15	ok	6ms / 512.00Kb	-
	16	ok	107ms / 728.00Kb	-
	17	ok	108ms / 728.00Kb	-
	18	ok	107ms / 728.00Kb	-
	19	ok	107ms / 728.00Kb	-
	20	ok	107ms / 724.00Kb	-
	21	time-limit-exceeded	1.09s / 3.02Mb	-
*/

func main() {
	var n int
	fmt.Fscan(os.Stdin, &n)
	ages := make([]int, n)
	fmt.Fscan(os.Stdin, packAddrs(ages)...)

	sort.Ints(ages)

	sum := 0
	for x := n - 1; x >= 0; x-- {

		h := (float64(ages[x]) * 0.5) + 7

		for y := x - 1; y >= 0; y-- {

			if ages[y] <= ages[x] && float64(ages[y]) > h {
				if ages[x] == ages[y] {
					sum += 2
				} else {
					sum++
				}
			}
		}
	}

	fmt.Fprintln(os.Stdout, sum)
}

func packAddrs(n []int) []interface{} {
	p := make([]interface{}, len(n))
	for i := range n {
		p[i] = &n[i]
	}
	return p
}
