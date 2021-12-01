package main

import (
	"fmt"
)

const max_age = 9

// School of lanternfish
type lf struct {
	s [max_age]int
}

func (o lf) compute_total() int {
	v := 0
	for a := 0; a < max_age; a++ {
		v += o.s[a]
	}
	return v
}

func (o *lf) simulate_day() {
	t := lf{} // Temporary school that will replace the old
	for a := max_age; a >= 0; a-- {
		switch a {
		case 1, 2, 3, 4, 5, 6, 7, 8:
			t.s[a-1] = o.s[a]
			break
		case 0:
			t.s[8] = o.s[a]
			t.s[6] += o.s[a]
			break
		}
	}

	for a := 0; a < max_age; a++ {
		o.s[a] = t.s[a]
	}
}

func main() {
	l := lf{}

	v := 0
	for {
		n, _ := fmt.Scanf("%d", &v)

		if n == 0 {
			break
		}

		l.s[v]++
	}

	for i := 0; i < 80; i++ {
		l.simulate_day()
	}

	fmt.Println("Part 1:", l.compute_total())

	for i := 0; i < 176; i++ {
		l.simulate_day()
	}

	fmt.Println("Part 2:", l.compute_total())
}
