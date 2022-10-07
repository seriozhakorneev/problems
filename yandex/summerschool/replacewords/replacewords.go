package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// TODO: Unsolved??
func main() {
	dict, text := stdIn()
	fmt.Fprintln(os.Stdout, replaceWords(dict, text))
}

/*
№	Вердикт	Ресурсы	Баллы
1	ok	3ms / 516.00Kb	-
2	ok	3ms / 512.00Kb	-
3	ok	3ms / 516.00Kb	-
4	ok	3ms / 516.00Kb	-
5	ok	3ms / 520.00Kb	-
6	ok	2ms / 516.00Kb	-
7	ok	3ms / 508.00Kb	-
8	ok	3ms / 636.00Kb	-
9	ok	8ms / 600.00Kb	-
10	ok	6ms / 636.00Kb	-
11	ok	5ms / 640.00Kb	-
12	ok	6ms / 640.00Kb	-
13	ok	9ms / 772.00Kb	-
14	wrong-answer	3ms / 764.00Kb	-
*/

func replaceWords(dict, text []string) string {
	for _, short := range dict {
		for i := 0; i < len(text); i++ {
			if len(text[i]) > len(short) && short == text[i][:len(short)] {
				text[i] = short
			}
		}
	}
	return strings.Join(text, " ")
}

func stdIn() (words []string, text []string) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	for _, el := range strings.Split(scanner.Text(), " ") {
		if string(el) != " " {
			words = append(words, el)
		}
	}
	scanner.Scan()
	for _, el := range strings.Split(scanner.Text(), " ") {
		if string(el) != " " {
			text = append(text, el)
		}
	}
	return
}
