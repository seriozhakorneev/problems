package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//n := 5 // 2
	//n := 8 // 3
	//n := 2147483647 // 65535

	input := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(input, &n)
	fmt.Fprintln(os.Stdout, constructionOfStairs(n))
}

func constructionOfStairs(blocks int) int {
	count := 0
	for {
		if blocks < 1 || blocks/(count+1) == 0 && blocks-(count+1) != 0 {
			break
		}
		blocks -= count + 1
		count++
	}
	return count
}
