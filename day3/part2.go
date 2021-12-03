package main

import (
	"fmt"
	"strconv"
)

// Reduces slice of string to just length one
// if b is true it keeps the most common bit
// if b is false it keeps the least common bit
func reduce_slice_len_to_one(s []string, b bool) string {
	i := 0         // Index in one number
	var r []string // Remaining
	for {
		z := 0
		o := 0
		// Find number of ones and zeroes in one column
		for _, v := range s {
			if v[i] == '1' {
				o++
			} else {
				z++
			}
		}

		// Set comparison bit by the most or least common bit
		var c byte
		if b {
			if o >= z {
				c = '1'
			} else {
				c = '0'
			}
		} else {
			if z <= o {
				c = '0'
			} else {
				c = '1'
			}
		}

		// Filter out undesired numbers by comparison bit
		// If there were zero matching it might remove all and cause a crash
		// However we know this is not the case as there must be a solution
		r = make([]string, 0)
		for _, v := range s {
			if v[i] == c {
				r = append(r, v)
			}
		}

		if len(r) == 1 {
			break
		}
		s = r
		i++
	}

	return r[0]
}

func main() {
	var l string
	o := make([]string, 0)
	c := make([]string, 0)
	for {
		n, _ := fmt.Scanf("%s", &l)
		if n == 0 {
			break
		}
		o = append(o, l)
		c = append(c, l)
	}

	// Compute oxygen generator rating and C02 scrubber rating
	o_r, _ := strconv.ParseInt(reduce_slice_len_to_one(o, true), 2, 32)
	c_r, _ := strconv.ParseInt(reduce_slice_len_to_one(c, false), 2, 32)
	fmt.Println(o_r * c_r)
}
