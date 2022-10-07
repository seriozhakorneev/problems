package main

import (
	"fmt"
)

// university task
func main() {
	var input int
	fmt.Printf("Input: ")
	_, err := fmt.Scanf("%d", &input)
	if err != nil {
		return
	}

	turn := 1 // ход игры, начинаем с 1, будем инкрементировать его каждый раз на 1

	// Запускаем функцию decomposeNumber и декомпозируем входящее значение (input)
	// получаем первоначальный массив значений разложенных слагаемых input
	// который мы будем использовать в основном цикле
	arr := decomposeNumber(input)

	// выводим входящее значение (input), первый шаг и первоначальный массив arr в терминал
	fmt.Println("\nINPUT:", input)
	printResult(arr, turn)

	for {
		turn += 1
		arr = decomposeArray(arr)

		if len(arr) == 0 {
			fmt.Println("\nEND")
			break
		}
		printResult(arr, turn)
	}
}

// printResult выводит результат в консоль
func printResult(arr [][]int, turn int) {
	fmt.Printf("\nTurn %v:", turn)
	var number int
	for _, el := range arr {
		if number != el[0] {
			number = el[0]
			fmt.Printf("\n%v = ", number)
		}
		fmt.Printf("%v+%v, ", el[1], el[2])
	}
}

// decomposeArray возвращает массив разложенных значений inputArray
func decomposeArray(inputArray [][]int) (outputArray [][]int) {
	for _, insideArray := range inputArray {
		for _, number := range insideArray[1:] {
			inputArray = decomposeNumber(number)
			outputArray = append(outputArray, inputArray...)
		}
	}
	return
}

// decomposeNumber раскладывает число на слагаемые
// возвращает массив массивов каждый из которых в сумме даст число
func decomposeNumber(num int) (arr [][]int) {
	for first := 1; first < num; first++ {
		arr = addToSet(arr, []int{num, first, num - first})
	}
	return
}

// addToSet проверяет наличие в array уже существующие значения из addArray
// и если их нет добавляет новые, иначе возвращает неизмененный массив array
func addToSet(array [][]int, addArray []int) [][]int {
	for _, inArray := range array {
		for _, ele := range inArray {
			if ele == addArray[1] {
				return array
			}
		}
	}
	return append(array, addArray)
}
