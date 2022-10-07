package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	desirable, received := stdIn()
	fmt.Fprintf(os.Stdout, rejection(desirable, received))
}

func rejection(desirable, received []string) string {
	var result []string

	first := make(map[string]bool)
	second := make(map[string]string)
	third := make(map[string]string)
	for _, d := range desirable {
		first[d] = true

		lower := strings.ToLower(d)
		if _, ok := second[lower]; !ok {
			second[lower] = d
		}

		noVowels := strings.Map(vowelsFunc, lower)
		if _, ok := third[noVowels]; !ok {
			third[noVowels] = d
		}
	}

	for _, r := range received {
		lower := strings.ToLower(r)
		v, ok := "", false

		if _, ok = first[r]; ok {
			result = append(result, r)
		} else if v, ok = second[lower]; ok {
			result = append(result, v)
		} else if v, ok = third[strings.Map(vowelsFunc, lower)]; ok {
			result = append(result, v)
		} else {
			result = append(result, "")
			continue
		}
	}
	return strings.Join(result, " ")
}

var vowelsFunc = func(r rune) rune {
	switch r {
	case 'a':
		return '#'
	case 'e':
		return '#'
	case 'i':
		return '#'
	case 'o':
		return '#'
	case 'u':
		return '#'
	default:
		return r
	}
}

func stdIn() (desirable, received []string) {
	var n int
	fmt.Fscan(os.Stdin, &n)
	desirable = make([]string, n)
	fmt.Fscan(os.Stdin, packAddrs(desirable)...)

	var k int
	fmt.Fscan(os.Stdin, &k)
	received = make([]string, k)
	fmt.Fscan(os.Stdin, packAddrs(received)...)
	return desirable, received
}

func packAddrs(n []string) []interface{} {
	p := make([]interface{}, len(n))
	for i := range n {
		p[i] = &n[i]
	}
	return p
}
