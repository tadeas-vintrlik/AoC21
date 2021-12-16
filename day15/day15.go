package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type Matrix struct {
	t [][]uint
	l uint
}

type Point struct {
	x, y uint
	v    uint
}

type PriorityQueue []Point

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].v < pq[j].v
}

func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
	item := (*pq)[n-1]
	*pq = (*pq)[0 : n-1]
	return item
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Point))
}

func (pq PriorityQueue) Swap(i, j int) {
	tmp := pq[i]
	pq[i] = pq[j]
	pq[j] = tmp
}

func (m Matrix) get_value(x, y uint) uint {
	return m.t[y][x]
}

func (m *Matrix) set_value(x, y, v uint) {
	(*m).t[y][x] = v
}

func (m Matrix) get_neighbours(p Point, visited map[[2]uint]bool) []Point {
	r := make([]Point, 0)
	if int(p.x-1) >= 0 {
		if _, d := visited[[2]uint{p.x - 1, p.y}]; !d {
			r = append(r, Point{p.x - 1, p.y, m.get_value(p.x-1, p.y)})
		}
	}
	if int(p.y-1) >= 0 {
		if _, d := visited[[2]uint{p.x, p.y - 1}]; !d {
			r = append(r, Point{p.x, p.y - 1, m.get_value(p.x, p.y-1)})
		}
	}
	if p.x+1 < m.l {
		if _, d := visited[[2]uint{p.x + 1, p.y}]; !d {
			r = append(r, Point{p.x + 1, p.y, m.get_value(p.x+1, p.y)})
		}
	}
	if p.y+1 < m.l {
		if _, d := visited[[2]uint{p.x, p.y + 1}]; !d {
			r = append(r, Point{p.x, p.y + 1, m.get_value(p.x, p.y+1)})
		}
	}
	return r
}

// Create a Matrix filled with 0 with the same dimensions
func create_risk_Matrix(m Matrix) Matrix {
	n := Matrix{}
	n.t = make([][]uint, m.l)
	for i, _ := range m.t {
		n.l++
		tmp := make([]uint, m.l)
		for j, _ := range tmp {
			tmp[j] = 0
		}
		n.t[i] = tmp
	}

	return n
}

func five_time_matrix(m Matrix) Matrix {
	r := Matrix{}
	r.l = m.l * 5
	for y := uint(0); y < r.l; y++ {
		r.t = append(r.t, make([]uint, r.l))
		for x := uint(0); x < r.l; x++ {
			n := (m.t[y%m.l][x%m.l] + (x / m.l) + (y / m.l))
			if n > 9 {
				n -= 9
			}
			r.t[y][x] = n
		}
	}
	return r
}

func (m Matrix) dijkstra() Matrix {
	r := create_risk_Matrix(m)

	pq := make(PriorityQueue, 0)
	visited := make(map[[2]uint]bool)
	i := 0

	heap.Init(&pq)
	heap.Push(&pq, Point{0, 0, 0})

	for len(pq) != 0 {
		i++
		c := heap.Pop(&pq).(Point)
		for _, v := range m.get_neighbours(c, visited) {
			old := r.get_value(v.x, v.y)
			tenative := c.v + v.v
			if old == 0 || tenative < old {
				v.v = tenative
				r.set_value(v.x, v.y, v.v)
				heap.Push(&pq, v)
			}
		}
		visited[[2]uint{c.x, c.y}] = true
	}

	return r
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	m := Matrix{}
	m.t = make([][]uint, 0)
	for s.Scan() {
		m.t = append(m.t, make([]uint, 0))
		for _, v := range s.Text() {
			m.t[m.l] = append(m.t[m.l], uint(v-'0'))
		}
		m.l++
	}

	r1 := m.dijkstra()
	fmt.Println("Part 1 solution:", r1.t[r1.l-1][r1.l-1])
	b := five_time_matrix(m)
	r2 := b.dijkstra()
	fmt.Println("Part 2 solution:", r2.t[r2.l-1][r2.l-1])
}
