package main

import (
	"fmt"
)

type submarine struct {
	horizontal, depth int
}

func (s *submarine) forward(distance int) {
	s.horizontal += distance
}

func (s *submarine) down(distance int) {
	s.depth += distance
}

func (s *submarine) up(distance int) {
	s.depth -= distance
}

func main() {
	s := submarine{depth: 0, horizontal: 0}
	for {
		var movement string
		var distance int
		n, _ := fmt.Scanf("%s %d", &movement, &distance)
		if n == 0 {
			break
		}
		switch movement {
		case "forward":
			s.forward(distance)
			break
		case "down":
			s.down(distance)
			break
		case "up":
			s.up(distance)
			break
		}
	}
	fmt.Println("Depth:", s.depth)
	fmt.Println("Horizontal:", s.horizontal)
	fmt.Println("Result:", s.depth*s.horizontal)
}
