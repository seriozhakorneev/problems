package main

import (
	"fmt"
	"os"
)

func main() {
	codes := map[string]uint8{
		"00":  'a',
		"100": 'b',
		"101": 'c',
		"11":  'd',
	}

	var t int
	fmt.Fscan(os.Stdin, &t)

	for i := 0; i < t; i++ {
		var encoded string
		fmt.Fscan(os.Stdin, &encoded)

		var decoded []byte
		for l, r := 0, 2; r <= len(encoded); r++ {
			if v, ok := codes[encoded[l:r]]; ok {
				decoded = append(decoded, v)
				l = r
				r++
			}
		}
		fmt.Fprintln(os.Stdout, string(decoded))
	}
}
