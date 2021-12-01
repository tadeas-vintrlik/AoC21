package main

import (
	"fmt"
)

const size = 1000

type diagram struct {
	board [size][size]int
}

func (d *diagram) get_no_crossed() int {
	c := 0
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			if d.board[y][x] >= 2 {
				c++
			}
		}
	}
	return c
}

// s if to only mark straight lines
func (d *diagram) mark_vector(v vector, s bool) {
	if v.x1 == v.x2 {
		if v.y1 > v.y2 {
			for i := v.y2; i <= v.y1; i++ {
				d.board[i][v.x1]++
			}
		} else {
			for i := v.y1; i <= v.y2; i++ {
				d.board[i][v.x1]++
			}
		}
		return
	}
	if v.y1 == v.y2 {
		if v.x1 > v.x2 {
			for i := v.x2; i <= v.x1; i++ {
				d.board[v.y1][i]++
			}
		} else {
			for i := v.x1; i <= v.x2; i++ {
				d.board[v.y1][i]++
			}
		}
		return
	}

	if !s {
		if v.x1 > v.x2 && v.y1 > v.y2 {
			for i := 0; i <= v.x1-v.x2; i++ {
				d.board[v.y2+i][v.x2+i]++
			}
		} else if v.x1 > v.x2 && v.y1 < v.y2 {
			for i := 0; i <= v.x1-v.x2; i++ {
				d.board[v.y2-i][v.x2+i]++
			}

		} else if v.x1 < v.x2 && v.y1 > v.y2 {
			for i := 0; i <= v.x2-v.x1; i++ {
				d.board[v.y2+i][v.x2-i]++
			}
		} else if v.x1 < v.x2 && v.y1 < v.y2 {
			for i := 0; i <= v.x2-v.x1; i++ {
				d.board[v.y2-i][v.x2-i]++
			}
		}
	}
}

type vector struct {
	x1, y1, x2, y2 int
}

func load_vector() (vector, bool) {
	var x1, x2, y1, y2 int
	n, _ := fmt.Scanf("%d,%d -> %d,%d\n", &x1, &y1, &x2, &y2)
	if n == 0 {
		return vector{}, false
	}
	return vector{x1, y1, x2, y2}, true
}

func main() {
	d1 := diagram{}
	d2 := diagram{}

	for {
		v, c := load_vector()
		if !c {
			break
		}
		d1.mark_vector(v, true)
		d2.mark_vector(v, false)
	}

	fmt.Println("Part 1 solution:", d1.get_no_crossed())
	fmt.Println("Part 2 solution:", d2.get_no_crossed())
}
