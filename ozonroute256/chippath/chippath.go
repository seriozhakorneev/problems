package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	field   [][]string
	path    []uint8
	row     string
	t, n, m int
)

func main() {

	fmt.Fscan(os.Stdin, &t)
	for ; t > 0; t-- {

		fmt.Fscan(os.Stdin, &n, &m)
		field = make([][]string, n)
		for i := 0; i < n; i++ {

			fmt.Fscan(os.Stdin, &row)
			field[i] = strings.Split(row, "")
		}
		getPath(field)
	}
}

func getPath(field [][]string) {
	for a, array := range field {
		for i, v := range array {
			if v == "*" {
				if nA, nI, final := isFinal(a, i); final {
					toNext(nA, nI, a, i)
					fmt.Fprintln(os.Stdout, string(path))
					return
				}
			}
		}
	}
}

func toNext(a, i, pA, pI int) {
	// up
	if (a != 0 && field[a-1][i] == "*") && (a-1 != pA || i != pI) {
		nA, nI := a-1, i
		path = append(path, 'U')
		toNext(nA, nI, a, i)
		return
	}
	// left
	if (i != 0 && field[a][i-1] == "*") && (a != pA || i-1 != pI) {
		nA, nI := a, i-1
		path = append(path, 'L')
		toNext(nA, nI, a, i)
		return
	}
	// right
	if (i != m-1 && field[a][i+1] == "*") && (a != pA || i+1 != pI) {
		nA, nI := a, i+1
		path = append(path, 'R')
		toNext(nA, nI, a, i)
		return
	}
	// down
	if (a != n-1 && field[a+1][i] == "*") && (a+1 != pA || i != pI) {
		nA, nI := a+1, i
		path = append(path, 'D')
		toNext(nA, nI, a, i)
		return
	}
}

func isFinal(a, i int) (nA, nI int, final bool) {
	var next uint8
	count := 0
	// up
	if a != 0 && field[a-1][i] == "*" {
		count++
		next = 'U'
		nA, nI = a-1, i
	}
	// left
	if i != 0 && field[a][i-1] == "*" {
		count++
		next = 'L'
		nA, nI = a, i-1
	}
	// right
	if i != m-1 && field[a][i+1] == "*" {
		count++
		next = 'R'
		nA, nI = a, i+1
	}
	// down
	if a != n-1 && field[a+1][i] == "*" {
		count++
		next = 'D'
		nA, nI = a+1, i
	}

	if count == 2 {
		return 0, 0, false
	}
	path = []uint8{next}
	return nA, nI, true
}
