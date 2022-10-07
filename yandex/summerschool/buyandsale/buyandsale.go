package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
№	Вердикт	Ресурсы	Баллы
1	ok	3ms / 520.00Kb	-
2	ok	2ms / 516.00Kb	-
3	ok	3ms / 516.00Kb	-
4	ok	2ms / 516.00Kb	-
5	ok	3ms / 512.00Kb	-
6	wrong-answer	3ms / 512.00Kb	-
*/

func main() {
	n, prices := stdIn()
	f, s := buyNSell(n, prices)
	fmt.Fprintln(os.Stdout, f, s)
}

// todo последнее видео объяснение цены

func buyNSell(n int, prices []int) (first, second int) {
	profit, minF, minS := 0, 0, 0
	for i := 0; i < n; i++ {

		if prices[first] > prices[i] {
			first, second = i, i
		}
		if prices[second] < prices[i] {
			second = i
		}

		if prices[second]-prices[first] > profit && prices[first] <= 1000 {
			profit, minF, minS = prices[second]-prices[first], first, second
		}
	}

	if minF < minS {
		return minF + 1, minS + 1
	}
	return 0, 0
}

func stdIn() (int, []int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	count, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	var n []int
	for _, el := range strings.Split(scanner.Text(), " ") {
		if string(el) != " " {
			k, _ := strconv.Atoi(el)
			n = append(n, k)
		}
	}
	return count, n
}

//f, s := buyNSell(6, []int{0,0,0,0,0,0}) // 0, 0
//f, s := buyNSell(1, []int{9}) // 0, 0
//f, s := buyNSell(6, []int{10, 3, 5, 3, 11, 9}) // 2,5
//f, s := buyNSell(4, []int{5, 5, 5, 5}) // 0,0
//f, s := buyNSell(6, []int{7,1,5,3,6,4}) // 2,5
//f, s := buyNSell(5, []int{7,6,4,3,1}) // 0,0
//f, s := buyNSell(7, []int{2, 4, 6, 3, 7, 1, 5, 0})// 1,5
//fmt.Fprintln(os.Stdout, f, s)

func test() {
	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= 12; i++ {
		// Generate a random array of length i
		a := rand.Perm(i)

		f, s := buyNSell(len(a), a)
		fmt.Println(a, f, s)
	}
}
