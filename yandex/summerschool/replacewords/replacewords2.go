package main

// Вариант списанный с решения лектора по задачам у него на питоне проходит все тесты
// у меня на го ломается на том же тесте как и в первом варианте

/*
func main() {
	dict, text := stdIn()
	fmt.Fprintln(os.Stdout, replaceWords(dict, text))
}

func replaceWords(dict map[string]bool, text []string) string {
	var result []string
	for _, word := range text {
		i := 1
		for ; i < len(word); i++ {
			part := word[:i]
			if _, ok := dict[word[:i]]; ok {
				result = append(result, part)
				break
			}
		}
		if i == len(word) {
			result = append(result, word)
		}
	}
	return strings.Join(result, " ")
}

func stdIn() (map[string]bool, []string) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	dict := make(map[string]bool)
	for _, el := range strings.Split(scanner.Text(), " ") {
		if string(el) != " " {
			dict[el] = true
		}
	}

	scanner.Scan()
	var text []string
	for _, el := range strings.Split(scanner.Text(), " ") {
		if string(el) != " " {
			text = append(text, el)
		}
	}
	return dict, text
}
*/
