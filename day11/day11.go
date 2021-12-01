package main

import (
	"bufio"
	"fmt"
	"os"
)

// Matrix (2D slice)
type m struct {
	s     [][]int
	max_x int
	max_y int
}

func (m m) String() string {
	r := ""
	ly := m.max_y
	lx := m.max_x
	for i := 0; i < ly; i++ {
		for j := 0; j < lx; j++ {
			r += fmt.Sprintf("%d", m.s[i][j])
		}
		r += "\n"
	}
	return r
}

func (m *m) Append(s string) {
	l := len(s)
	n := make([]int, 0)

	// Convert the ASCII values into integers from 0-9
	for i := 0; i < l; i++ {
		n = append(n, int(s[i]-'0'))
	}

	// Append the line to the matrix, update size
	m.s = append(m.s, n)
	if m.max_x == 0 {
		// Asuming all lines will have the same length
		m.max_x = l
	}
	m.max_y++
}

func (m m) Get(x, y int) int {
	return m.s[y][x]
}

func (m *m) Set(x, y, v int) {
	m.s[y][x] = v
}

func (m *m) TryFlash(x, y int) int {
	// Check coordinates
	if x < 0 || y < 0 || x >= m.max_x || y >= m.max_y {
		return 0
	}

	v := m.Get(x, y)
	r := 0
	if v == 0 {
		return 0 // Already flashed this step
	} else if v >= 9 {
		r += 1
		m.Set(x, y, 0)
		r += m.TryFlash(x+1, y)
		r += m.TryFlash(x, y+1)
		r += m.TryFlash(x+1, y+1)
		r += m.TryFlash(x-1, y)
		r += m.TryFlash(x, y-1)
		r += m.TryFlash(x-1, y-1)
		r += m.TryFlash(x-1, y+1)
		r += m.TryFlash(x+1, y-1)
	} else {
		m.Set(x, y, v+1)
	}
	return r
}

func (m *m) ProcessStep() int {
	f := make([][2]int, 0) // List of indexes to flash
	for y := 0; y < m.max_y; y++ {
		for x := 0; x < m.max_x; x++ {
			v := m.Get(x, y) + 1
			m.Set(x, y, v)
			if v > 9 {
				f = append(f, [2]int{x, y})
			}
		}
	}

	r := 0

	for _, v := range f {
		r += m.TryFlash(v[0], v[1])
	}

	return r
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	m := m{}

	for s.Scan() {
		l := s.Text()
		m.Append(l)
	}

	r1 := 0
	r2 := 0
	i := 0
	// See how many flashes there are in a 100 steps
	for ; i < 100; i++ {
		r1 += m.ProcessStep()
	}

	// Find the first step where all octopuses flashed together for the first time
	for {
		if m.ProcessStep() == m.max_x*m.max_y {
			r2 = i + 1
			break
		}
		i++
	}

	fmt.Println("Part 1 result:", r1)
	fmt.Println("Part 2 result:", r2)
}
