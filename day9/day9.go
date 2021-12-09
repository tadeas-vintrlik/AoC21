package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// Heightmap
type hm struct {
	m     []string
	max_x int
	max_y int
}

// Assumes all strings added will have same length
func (h *hm) add(s string) {
	h.m = append(h.m, s)
	if h.max_x == 0 {
		h.max_x = len(s)
	}
	h.max_y++
}

func (h hm) get(x int, y int) int {
	return int(h.m[y][x]) - int('0')
}

func (h hm) is_low_point(x int, y int) bool {
	v := h.get(x, y)

	// Check all neighbors
	if x+1 < h.max_x {
		if h.get(x+1, y) <= v {
			return false
		}
	}
	if x-1 >= 0 {
		if h.get(x-1, y) <= v {
			return false
		}
	}
	if y+1 < h.max_y {
		if h.get(x, y+1) <= v {
			return false
		}
	}
	if y-1 >= 0 {
		if h.get(x, y-1) <= v {
			return false
		}
	}

	return true
}

func (h hm) get_low_points() [][2]int {
	r := make([][2]int, 0)
	for y := 0; y < h.max_y; y++ {
		for x := 0; x < h.max_x; x++ {
			if h.is_low_point(x, y) {
				l := [2]int{x, y}
				r = append(r, l)
			}
		}
	}
	return r
}

func (h hm) get_risk_level(l [][2]int) int {
	r := 0
	for _, v := range l {
		r += h.get(v[0], v[1]) + 1
	}

	return r
}

func coord_in_basin(x, y int, b [][2]int) bool {
	for _, c := range b {
		if c[0] == x && c[1] == y {
			return true
		}
	}
	return false
}

// Use the flood fill algorithm
func (h hm) add_basin_coord_r(x, y int, b [][2]int) [][2]int {
	if x >= h.max_x || x < 0 {
		return b
	}
	if y >= h.max_y || y < 0 {
		return b
	}
	if h.get(x, y) == 9 {
		return b
	}
	if coord_in_basin(x, y, b) {
		// Already seen before
		return b
	}

	c := [2]int{x, y}
	b = append(b, c)

	b = h.add_basin_coord_r(x+1, y, b)
	b = h.add_basin_coord_r(x-1, y, b)
	b = h.add_basin_coord_r(x, y+1, b)
	b = h.add_basin_coord_r(x, y-1, b)

	return b
}

// returns sizes of all basins
func (h hm) get_basins(l [][2]int) []int {

	r := make([]int, 0)
	for _, v := range l {
		r = append(r, len(h.add_basin_coord_r(v[0], v[1], make([][2]int, 0))))
	}

	return r
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	hm := hm{}
	for s.Scan() {
		hm.add(s.Text())
	}

	lp := hm.get_low_points()
	fmt.Println("Part 1 solution:", hm.get_risk_level(lp))

	bs := hm.get_basins(lp)
	// Get three largest basins (sort in descending)
	sort.Sort(sort.Reverse(sort.IntSlice(bs)))
	fmt.Println("Part 2 solution:", bs[0]*bs[1]*bs[2])
}
