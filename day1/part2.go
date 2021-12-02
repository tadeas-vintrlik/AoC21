package main

import (
	"fmt"
)

func sum_triplet(t []int) int {
	return t[0] + t[1] + t[2]
}

func main() {
	t := make([]int, 0) // Triplets
	o := -1             // Old (negative means there was none)
	i := 0              // Number of increases in depth of triplets
	for {
		d := 0 // Depth
		n, _ := fmt.Scanf("%d", &d)
		if n == 0 {
			break
		}

		// Append to a triplet, if all three values exist compare them and remove first
		t = append(t, d)
		if len(t) == 3 {
			v := sum_triplet(t)
			if o != -1 && v > o {
				i++
			}
			o = v
			t = t[1:]
		}
	}

	fmt.Println("Increases in value of triplets:", i)
}
