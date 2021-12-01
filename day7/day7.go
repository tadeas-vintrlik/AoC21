package main

import (
	"fmt"
)

// fuel in part increased by 1 for each unit of distance
func compute_fuel_p2(v int) int {
	return (v + v*v) / 2
}

func main() {
	p := make([]int, 0)
	m := 0 // Maximum value used as upper bound

	// Load all crab horizontal positions
	for {
		var v int
		n, _ := fmt.Scanf("%d", &v)
		if n == 0 {
			break
		}

		if v > m {
			m = v
		}
		p = append(p, v)
	}

	// Try to find the minimal sum of all movements
	l := len(p)
	r1 := m * l                  // The resulting amount of fuel needed, defaults to maximum possible (part 1)
	r2 := compute_fuel_p2(m) * l // Same as above for part 2
	for i := 0; i <= m; i++ {
		s1 := 0 // Sum of all fuel needed for i (part 1)
		s2 := 0 // Sum of all fuel needed for i (part 2)

		for j := 0; j < l; j++ {
			// Absolute value of v as it is distance
			v := i - p[j]
			if v < 0 {
				v = -v
			}
			s1 += v
			s2 += compute_fuel_p2(v)
		}

		if s1 < r1 {
			r1 = s1
		}

		if s2 < r2 {
			r2 = s2
		}
	}

	fmt.Println("Part 1 solution:", r1)
	fmt.Println("Part 2 solution:", r2)
}
