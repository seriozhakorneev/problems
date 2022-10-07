package main

import (
	"bufio"
	"fmt"
	"os"
)

func readSlice(reader *bufio.Reader, n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return a
}

func countSum(count int, prices []int) int {
	// product price : count of this product price
	hash := make(map[int]int)
	for i := 0; i < count; i++ {
		hash[prices[i]]++
	}

	var sum int
	for k, v := range hash {
		sum += (v - (v / 3)) * k
	}

	return sum
}

func main() {
	input := bufio.NewReader(os.Stdin)
	output := bufio.NewWriter(os.Stdout)
	defer output.Flush()

	var inputCount int
	fmt.Fscan(input, &inputCount)

	for i := 0; i < inputCount; i++ {
		var count int
		fmt.Fscan(input, &count)
		prices := readSlice(input, count)

		fmt.Fprintln(output, countSum(count, prices))
	}
}

//count := 12
//prices := []int{2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3} // 22

//count := 1
//prices := []int{10000} // 10000

//count := 9
//prices := []int{1, 2, 3, 1, 2, 3, 1, 2, 3} // 12
//
//
//6
//300 100 200 300 200 300 // 1100
