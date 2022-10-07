package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"reflect"
	"testing"
)

func TestPairProgramming(t *testing.T) {

	files, _ := ioutil.ReadDir("/Users/korneevsergey/go_code/ozonroute256/pairprogramming/tests")
	fmt.Println("tests count:", len(files)-(len(files)/2))

	var testF, answerF *os.File
	for i := 0; i < len(files); i += 2 {
		testF, _ = os.Open(fmt.Sprintf(
			"/Users/korneevsergey/go_code/ozonroute256/pairprogramming/tests/%s",
			files[i].Name()))

		answerF, _ = os.Open(fmt.Sprintf(
			"/Users/korneevsergey/go_code/ozonroute256/pairprogramming/tests/%s",
			files[i+1].Name()))

		inputTest := bufio.NewReader(testF)
		inputAnswer := bufio.NewReader(answerF)
		var inputCount int
		fmt.Fscan(inputTest, &inputCount)

		for j := 0; j < inputCount; j++ {
			var count int
			fmt.Fscan(inputTest, &count)
			levels := readSliceTest(inputTest, count)

			levelsCopy := make([]int, count)
			copy(levelsCopy, levels)
			testAnswer := readSliceTest(inputAnswer, count)
			answer := makePairsTest(count, levelsCopy)

			if !reflect.DeepEqual(answer, testAnswer) {
				fmt.Println(
					"levels:", levels,
					"result:", makePairsTest(count, levelsCopy),
					"answer:", testAnswer,
				)
				os.Exit(1)
			}

		}

		testF.Close()
		answerF.Close()
	}
}

func makePairsTest(count int, levels []int) (res []int) {
	for i := 0; i < count; i++ {
		if levels[i] == 101+99 {
			continue
		}
		index := closestLevelIndexTest(levels[i+1:], levels[i]) + len(levels[:i])
		levels[i], levels[index] = 101+99, 101+99
		res = append(res, i+1, index+1)
	}
	return
}

func closestLevelIndexTest(array []int, num int) int {
	diff := float64(101)
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

func readSliceTest(reader *bufio.Reader, n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return a
}
