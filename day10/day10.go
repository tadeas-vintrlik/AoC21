package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// Character stack
type stack struct {
	depth  int
	values []byte
}

func (s stack) Top() byte {
	if s.depth == 0 {
		fmt.Fprintf(os.Stderr, "Top on empty stack.\n")
		return ' '
	}
	return s.values[s.depth-1]
}

func (s *stack) Pop() {
	if s.depth == 0 {
		fmt.Fprintf(os.Stderr, "Pop on empty stack.\n")
		return
	}
	s.depth--
	s.values = s.values[:s.depth]
}

func (s stack) IsEmpty() bool {
	return s.depth == 0
}

func (s *stack) Push(v byte) {
	s.values = append(s.values, v)
	s.depth++
}

func matching_paren(b1 byte, b2 byte) bool {
	if b1 == '(' && b2 == ')' {
		return true
	} else if b1 == '[' && b2 == ']' {
		return true
	} else if b1 == '{' && b2 == '}' {
		return true
	} else if b1 == '<' && b2 == '>' {
		return true
	}
	return false
}

func paren_points(b byte) int {
	switch b {
	case ')':
		return 3
	case ']':
		return 57
	case '}':
		return 1197
	case '>':
		return 25137
	}
	return 0
}

func paren_points_incomplete(r int, b byte) int {
	r *= 5
	switch b {
	case '(':
		r += 1
	case '[':
		r += 2
	case '{':
		r += 3
	case '<':
		r += 4
	}
	return r
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	st := stack{}
	r1 := 0
	r2 := make([]int, 0)
	for sc.Scan() {
		// Process the line
		r := 0
		for _, v := range sc.Text() {
			b := byte(v)
			switch b {
			case '(', '[', '{', '<':
				st.Push(byte(b))
			default:
				// Closing parentheses
				if t := st.Top(); !matching_paren(t, b) {
					r += paren_points(b)
				}
				st.Pop()
			}
		}
		// Empty the stack
		r1 += r
		ls := 0 // Line score from completion
		for !st.IsEmpty() {
			if r == 0 {
				// Incomplete line otherwise correct
				t := st.Top()
				ls = paren_points_incomplete(ls, t)
			}
			st.Pop()
		}
		if ls != 0 {
			r2 = append(r2, ls)
		}
	}
	fmt.Println("Part 1 result:", r1)
	sort.Ints(r2)
	fmt.Println("Part 2 result:", r2[len(r2)/2])
}
