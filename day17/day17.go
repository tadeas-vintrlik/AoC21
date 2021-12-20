package main

import (
	"fmt"
)

func probe_step(coord, velocity *[2]int) {
	(*coord)[0] += (*velocity)[0]
	(*coord)[1] += (*velocity)[1]
	if (*velocity)[0] > 0 {
		(*velocity)[0]--
	} else if (*velocity)[0] < 0 {
		(*velocity)[0]++
	}
	(*velocity)[1]--
}

func in_target(c, bx, by [2]int) bool {
	return c[0] >= bx[0] && c[0] <= bx[1] && c[1] >= by[0] && c[1] <= by[1]
}

func overshot(c, bx, by [2]int) bool {
	if c[0] > bx[1] {
		return true
	}

	if by[0] < 0 && c[1] < by[0] {
		return true
	}

	if by[0] >= 0 && c[1] > by[0] {
		return true
	}

	return false
}

func simulate_probe_shot(velocity, bx, by [2]int) (bool, int) {
	coord := [2]int{0, 0}
	max_y := -1
	for !in_target(coord, bx, by) && !overshot(coord, bx, by) {
		probe_step(&coord, &velocity)
		if coord[1] > max_y || max_y == -1 {
			max_y = coord[1]
		}
	}
	return in_target(coord, bx, by), max_y
}

func main() {
	// Bounds for the target area
	var bx [2]int
	var by [2]int
	fmt.Scanf("target area: x=%d..%d, y=%d..%d\n", &bx[0], &bx[1], &by[0], &by[1])

	max := -1
	total_hit := 0
	for i := 0; i < 500; i++ {
		for j := -500; j < 500; j++ {
			hit, tmp := simulate_probe_shot([2]int{i, j}, bx, by)
			if hit {
				total_hit++
				if tmp > max || max == -1 {
					max = tmp
				}
			}
		}
	}
	fmt.Println("Max:", max)
	fmt.Println("Total hit:", total_hit)
}
