package main

import (
	"fmt"
	"os"
)

/*
Группа small: 15.0 балл(ов)
    Тест 1: Тест пройден
    Тест 2: Тест пройден
    Тест 3: Тест пройден
    Тест 4: Тест пройден
    Тест 5: Тест пройден
    Тест 6: Тест пройден
    Тест 7: Тест пройден
    Тест 8: Тест пройден
    Тест 9: Тест пройден
    Тест 10: Тест пройден
    Тест 11: Тест пройден
    Тест 12: Тест пройден
    Тест 13: Тест пройден
Группа large: 0.0 балл(ов)
    Тест 14: Тест пройден
    Тест 15: Тест пройден
    Тест 16: Тест пройден
    Тест 17: Превышено ограничение времени
*/

func main() {
	var t int
	fmt.Fscan(os.Stdin, &t)
	for ; t > 0; t-- {

		var n int
		fmt.Fscan(os.Stdin, &n)

		list := make([]int, 3)
		fmt.Fscan(os.Stdin, &list[1], &list[0], &list[2])

		triplets := make(map[int][3]int)
		for i := 0; i < n-1; i++ {
			var triplet [3]int
			fmt.Fscan(
				os.Stdin,
				&triplet[1],
				&triplet[0],
				&triplet[2],
			)
			triplets[i] = triplet
		}

		l, r := 0, len(list)-1
	loop:
		for {
			for k, triplet := range triplets {
				if list[l] == triplet[1] {
					if list[l+1] == triplet[0] {
						list = append([]int{triplet[2]}, list...)
						delete(triplets, k)

					} else if list[l+1] == triplet[2] {
						list = append([]int{triplet[0]}, list...)
						delete(triplets, k)
					}
				} else if list[r] == triplet[1] {
					if list[r-1] == triplet[0] {
						list = append(list, triplet[2])
						delete(triplets, k)

					} else if list[r-1] == triplet[2] {
						list = append(list, triplet[0])
						delete(triplets, k)
					}
				}

				if len(list) == n {
					break loop
				}
			}
		}

		for f, h := 0, len(list)/2; f < len(list)/2; f, h = f+1, h+1 {
			fmt.Fprintln(os.Stdout, list[f], list[h])
		}
	}
}
