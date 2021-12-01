package main

import (
	"fmt"
)

func main() {
	i := 0  // Number of increases
	o := -1 // Negative means there was not old value
	for {
		// Read until EOF
		d := 0 // Current depth
		n, _ := fmt.Scanf("%d", &d)
		if n == 0 {
			break
		}

		if o != -1 && d > o {
			i++
		}
		o = d
	}
	fmt.Println("Number of increases was:", i)
}
