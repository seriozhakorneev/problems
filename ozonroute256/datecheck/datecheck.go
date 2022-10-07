package main

import (
	"fmt"
	"os"
)

func main() {
	monthNDays := map[int]int{
		1: 31, 2: 28, 3: 31, 4: 30,
		5: 31, 6: 30, 7: 31, 8: 31,
		9: 30, 10: 31, 11: 30, 12: 31,
	}

	var t int
	fmt.Fscan(os.Stdin, &t)

	for i := 0; i < t; i++ {
		monthNDays[2] = 28
		var d, m, y int
		fmt.Fscan(os.Stdin, &d, &m, &y)

		if m == 2 && ((y%4 == 0 && y%100 != 0) || y%400 == 0) {
			monthNDays[2] = 29
		}

		if d > monthNDays[m] {
			fmt.Fprintln(os.Stdout, "NO")
			continue
		}
		fmt.Fprintln(os.Stdout, "YES")
	}
}
