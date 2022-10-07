package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewReader(os.Stdin)
	output := bufio.NewWriter(os.Stdout)
	defer output.Flush()

	var inputCount int
	fmt.Fscan(input, &inputCount)

	for i := 0; i < inputCount; i++ {
		var count int
		fmt.Fscan(input, &count)
		report := readSlice(input, count)
		meetsCriterion(count, report, output)
	}
}

func readSlice(reader *bufio.Reader, n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return a
}

func meetsCriterion(count int, report []int, output *bufio.Writer) {
	hash := make(map[int]bool)
	for i := 0; i < count; i++ {
		if hash[report[i]] && report[i-1] != report[i] {
			fmt.Fprintln(output, "NO")
			return
		}
		hash[report[i]] = true
	}
	fmt.Fprintln(output, "YES")
	return
}
