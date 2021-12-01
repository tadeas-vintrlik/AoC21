package main

import (
	"bufio"
	"fmt"
	"os"
)

type fold struct {
	d byte
	l uint
}

func contained(l [][2]uint, x, y uint) bool {
	for _, v := range l {
		if v[0] == x && v[1] == y {
			return true
		}
	}

	return false
}

func print_paper(l [][2]uint) {
	max_x := uint(0)
	max_y := uint(0)
	for _, v := range l {
		if v[0] > max_x {
			max_x = v[0]
		}
		if v[1] > max_y {
			max_y = v[1]
		}
	}

	for y := uint(0); y <= max_y; y++ {
		for x := uint(0); x <= max_x; x++ {
			if contained(l, x, y) {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}

func load_dots(s *bufio.Scanner) [][2]uint {
	dots := make([][2]uint, 0)
	for s.Scan() {
		t := s.Text()
		if t == "" {
			// Coordinates end and folds start
			break
		}

		var x, y uint
		fmt.Sscanf(t, "%d,%d", &x, &y)
		dots = append(dots, [2]uint{x, y})
	}

	return dots
}

func load_folds(s *bufio.Scanner) []fold {
	folds := make([]fold, 0)

	for s.Scan() {
		var d byte
		var l uint
		fmt.Sscanf(s.Text(), "fold along %c=%d", &d, &l)
		folds = append(folds, fold{d, l})
	}

	return folds
}

func fold_dots(d [][2]uint, f fold) [][2]uint {
	n := make([][2]uint, 0)
	i := 0
	if f.d == 'y' {
		i = 1
	}

	for _, v := range d {
		if v[i] >= f.l {
			v[i] = 2*f.l - v[i]
			if v[i] < 0 {
				v[i] = -v[i]
			}
		}
		if !contained(n, v[0], v[1]) {
			n = append(n, [2]uint{v[0], v[1]})
		}
	}
	return n
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	d := load_dots(s)
	f := load_folds(s)

	for i, v := range f {
		d = fold_dots(d, v)
		if i == 0 {
			fmt.Println("Part 1 Solution:", len(d))
		}
	}
	fmt.Println("Part 2 solution:")
	print_paper(d)
}
