package main

import (
	"fmt"
	"strconv"
	"strings"
)

type bboard struct {
	board [5][5]int
}

func (b bboard) get(x, y int) int {
	return b.board[y][x]
}

func (b *bboard) set(x, y, v int) {
	b.board[y][x] = v
}

func (b bboard) find(v int) (int, int) {
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			if b.get(x, y) == v {
				return x, y
			}
		}
	}
	return -1, -1
}

func (b *bboard) mark(v int) {
	x, y := b.find(v)
	if x != -1 && y != -1 {
		b.set(x, y, -1)
	}
}

func (b bboard) won() bool {
	for x := 0; x < 5; x++ {
		m := 0
		for y := 0; y < 5; y++ {
			if b.get(x, y) == -1 {
				m++
			}
		}
		if m == 5 {
			return true
		}
	}

	for y := 0; y < 5; y++ {
		m := 0
		for x := 0; x < 5; x++ {
			if b.get(x, y) == -1 {
				m++
			}
		}
		if m == 5 {
			return true
		}
	}

	return false
}

func (b bboard) get_score() int {
	s := 0
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			v := b.get(x, y)
			if v != -1 {
				s += v
			}
		}
	}
	return s
}

func load_bboard() (bboard, bool) {
	b := bboard{}
	fmt.Scanf("\n")
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			var v int
			n, _ := fmt.Scanf("%d", &v)
			if n == 0 {
				return b, false
			}
			b.set(x, y, v)
		}
	}
	return b, true
}

func load_draw_numbers() []int {
	// Load line
	var l string
	n := make([]int, 0)
	fmt.Scanf("%s\n", &l)

	// Split and convert to numbers, append to slice
	r := strings.Split(l, ",")
	for _, v := range r {
		a, _ := strconv.ParseInt(v, 10, 32)
		n = append(n, int(a))
	}

	return n
}

func part1(ns []int, bs []bboard, p int) {
	for _, n := range ns {
		for i, b := range bs {
			b.mark(n)
			bs[i] = b
			if b.won() {
				fmt.Println("Part", p, "solution:", n*b.get_score())
				return
			}
		}
	}
}

func part2(ns []int, bs []bboard) {
	for i, n := range ns {
		nbs := make([]bboard, 0)
		for j, b := range bs {
			b.mark(n)
			bs[j] = b
			if !b.won() {
				nbs = append(nbs, b)
			}
		}
		bs = nbs
		if len(bs) == 1 {
			part1(ns[i:], bs, 2)
			return
		}
	}
}

func main() {
	ns := load_draw_numbers()
	bs := make([]bboard, 0)
	for {
		b, e := load_bboard()
		if e {
			bs = append(bs, b)
		} else {
			break
		}
	}

	bs1 := make([]bboard, 0)
	copy(bs1, bs)
	part1(ns, bs, 1)
	part2(ns, bs)
}
