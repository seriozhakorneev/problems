package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const (
	// 1≤ai≤100
	impossibleValue = 101
	// 100 - 1
	placeholder = impossibleValue + 99
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
		levels := readSlice(input, count)
		makePairs(count, levels, output)
		fmt.Fprintln(output)
	}
}

func readSlice(reader *bufio.Reader, n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return a
}

func makePairs(count int, levels []int, output *bufio.Writer) {
	for i := 0; i < count; i++ {
		if levels[i] == placeholder {
			continue
		}
		index := closestLevelIndex(levels[i+1:], levels[i]) + len(levels[:i])
		levels[i], levels[index] = placeholder, placeholder
		fmt.Fprintln(output, i+1, index+1)
	}
}

func closestLevelIndex(array []int, num int) int {
	diff := float64(impossibleValue)
	var ans int
	for i := 0; i < len(array); i++ {
		m := math.Abs(float64(num - array[i]))
		if m < diff {
			diff = m
			ans = i
		}
	}
	return ans + 1
}
