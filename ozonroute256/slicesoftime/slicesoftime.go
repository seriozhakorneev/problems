package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// TODO: partially solved
func main() {
	input := bufio.NewReader(os.Stdin)
	output := bufio.NewWriter(os.Stdout)
	defer output.Flush()

	var inputCount int
	fmt.Fscan(input, &inputCount)
	for ; inputCount > 0; inputCount-- {

		var count int
		fmt.Fscan(input, &count)
		var intervals [][]int
		for ; count > 0; count-- {

			var timeSlice string
			fmt.Fscan(input, &timeSlice)

			left, right := medianDivide(timeSlice)

			if !isValidRe(left) || !isValidRe(right) {
				fmt.Fprintln(output, "NO")
				skip(count-1, input)
				break
			}

			lSec, rSec := recountInSec(left, right)

			if lSec > rSec {
				fmt.Fprintln(output, "NO")
				skip(count-1, input)
				break
			}

			intervals = append(intervals, []int{lSec, rSec})
		}

		if count == 0 && isIntersecting(intervals, len(intervals)) {
			fmt.Fprintln(output, "NO")
			continue
		} else if count == 0 {
			fmt.Fprintln(output, "YES")
		}
	}
}

func skip(remainder int, r *bufio.Reader) {
	var tmp string
	for ; remainder > 0; remainder-- {
		fmt.Fscan(r, &tmp)
	}
}

func medianDivide(s string) (left, right string) {
	median := len(s) / 2
	left = s[:median]
	right = s[median+1:]
	return
}

func isValidRe(s string) bool {
	re := regexp.MustCompile(`^(?:[01]\d|2[0-3]):[0-5]\d:[0-5]\d$`)
	if !re.MatchString(s) {
		return false
	}
	return true
}

func recountInSec(left, right string) (lSec, rSec int) {
	var l, r int
	for i, j := 0, 2; i < len(left); i, j = i+3, j+3 {

		l, _ = strconv.Atoi(left[i:j])
		r, _ = strconv.Atoi(right[i:j])

		switch i {
		case 0:
			lSec, rSec = l*3600, r*3600
		case 3:
			lSec += l * 60
			rSec += r * 60
		case 6:
			lSec += l
			rSec += r
		}
	}
	return lSec, rSec
}

func isIntersecting(intervals [][]int, n int) bool {
	if n < 2 {
		return false
	}
	l, r := intervals[0][0], intervals[0][1]

	for i := 1; i < n; i++ {
		if intervals[i][0] > r || intervals[i][1] < l {
			return false
		} else {
			l = max(l, intervals[i][0])
			r = min(r, intervals[i][1])
		}
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
