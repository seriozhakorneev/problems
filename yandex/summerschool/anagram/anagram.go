package main

import (
	"fmt"
	"os"
)

func main() {
	var first, second string
	fmt.Fscan(os.Stdin, &first, &second)
	fmt.Fprintln(os.Stdout, isAnagram(first, second))
}

func isAnagram(first, second string) string {
	if len(first) != len(second) {
		return "NO"
	}

	var sum uint8
	for i := 0; i < len(first); i++ {
		sum += first[i]
		sum -= second[i]
	}

	if sum != 0 {
		return "NO"
	}
	return "YES"
}
