package main

import (
	"fmt"
)

func main() {
	var line string
	fmt.Scanf("%s", &line)
	l := len(line)
	z := make([]int, l) // Zeroes
	o := make([]int, l) // Ones
	for {
		n, _ := fmt.Scanf("%s", &line)
		if n == 0 {
			break
		}

		for i := 0; i < l; i++ {
			if line[i] == '1' {
				o[i]++
			} else {
				z[i]++
			}
		}
	}

	g := 0
	e := 0
	x := 1 << (l - 1)
	for i := 0; i < l; i++ {
		if o[i] > z[i] {
			// One was the most common bit:
			g += x
		} else {
			// One was the least common bit:
			e += x
		}
		x = x >> 1
	}

	fmt.Println("Power consumption:", g*e)
}
