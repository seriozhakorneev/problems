package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// TODO: unsolved ??
func main() {
	andreiAndLiquid()
}

func andreiAndLiquid() {
	count, n := stdIn()
	if count == 0 || n == nil {
		fmt.Println(-1)
		return
	}

	if len(n) != count {
		fmt.Println(-1)
		return
	}
	// проверяю возможность разлива кислоты в n
	// делаю множество set чтобы проверить n на наличие более двух одинаковых значений при которых разлив не возможен
	// переменная b проверяет находятся ли значения не нуждающиеся в разливе перед нуждающимися
	set := map[int]bool{}
	var b int
	for _, el := range n {
		if el >= b {
			b = el
		} else {
			fmt.Println(-1)
			return
		}
		set[el] = true
	}

	if len(set) > 2 {
		fmt.Println(-1)
		return
	}

	fmt.Println(b - n[0])
	return
}

func stdIn() (int, []int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	count, _ := strconv.Atoi(scanner.Text())

	if count < 1 || count > 100000 {
		return 0, nil
	}

	scanner.Scan()
	var n []int
	for _, el := range strings.Split(scanner.Text(), " ") {
		if string(el) != " " {

			k, _ := strconv.Atoi(el)
			if k < 1 || k > 1000000000 {
				return 0, nil
			}
			n = append(n, k)
		}
	}
	return count, n
}
