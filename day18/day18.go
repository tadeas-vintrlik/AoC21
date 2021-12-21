package main

import (
	"bufio"
	"container/list"
	"fmt"
	"math"
	"os"
)

type SnailNum struct {
	val   int
	depth uint
}

func list_print(num *list.List) {
	e := num.Front()
	for e != nil {
		fmt.Print(e.Value.(SnailNum), " || ")
		e = e.Next()
	}
	fmt.Print("\n")
}

func (sn SnailNum) String() string {
	return fmt.Sprintf("%d (depth: %d)", sn.val, sn.depth)
}

func parse_snail_number(line string) *list.List {
	num := list.New()
	num.Init()

	depth := uint(0)
	for _, v := range line {
		switch v {
		case '[':
			depth++
		case ']':
			depth--
		case ',':
		default: // number
			num.PushBack(SnailNum{int(v - '0'), depth})
		}
	}

	return num
}

func explode(num *list.List, e *list.Element) {
	ev := e.Value.(SnailNum)
	n := e.Next()
	nv := n.Value.(SnailNum)

	if ev.depth != nv.depth {
		fmt.Println(ev.depth, "!=", nv.depth)
		panic("Depths did not match in explosion of a tuple")
	}

	left := e.Prev()
	if left != nil {
		leftv := left.Value.(SnailNum)
		left.Value = SnailNum{leftv.val + ev.val, leftv.depth}
	}

	right := n.Next()
	if right != nil {
		rightv := right.Value.(SnailNum)
		right.Value = SnailNum{rightv.val + nv.val, rightv.depth}
	}

	num.Remove(n)
	e.Value = SnailNum{0, ev.depth - 1}
}

func split(num *list.List, e *list.Element) {
	ev := e.Value.(SnailNum)
	n1 := int(math.Floor(float64(ev.val) / 2.0))
	n2 := int(math.Ceil(float64(ev.val) / 2.0))
	depth := ev.depth + 1

	e.Value = SnailNum{n1, depth}
	split2 := SnailNum{n2, depth}
	num.InsertAfter(split2, e)
}

func reduce(num *list.List) {
	no_changes := -1

	for no_changes != 0 {
		no_changes = 0

		e := num.Front()
		for e != nil {
			n := e.Value.(SnailNum)

			if n.depth >= 5 {
				no_changes++
				explode(num, e)
				e = num.Front()
			} else {
				e = e.Next()
			}
		}

		e = num.Front()
		for e != nil {
			n := e.Value.(SnailNum)

			if n.val >= 10 {
				no_changes++
				split(num, e)
				break
			} else {
				e = e.Next()
			}
		}
	}
}

func indent(l *list.List) {
	e := l.Front()
	for e != nil {
		v := e.Value.(SnailNum)
		v.depth++
		e.Value = v
		e = e.Next()
	}
}

func add(num, to_add *list.List) {
	indent(num)
	indent(to_add)
	num.PushBackList(to_add)
	reduce(num)
}

func magnitude(num *list.List) uint {
	magn := list.New()
	magn.PushBackList(num)

	e := magn.Front()
	for d := uint(4); d != 0; d-- {
		e = magn.Front()
		for e != nil && e.Next() != nil {
			ev := e.Value.(SnailNum)
			n := e.Next()
			nv := n.Value.(SnailNum)
			if ev.depth == d && ev.depth == nv.depth {
				e.Value = SnailNum{3*ev.val + 2*nv.val, ev.depth - 1}
				magn.Remove(n)
			}
			e = e.Next()
		}
	}

	return uint(magn.Front().Value.(SnailNum).val)
}

func part1(to_add []*list.List) uint {
	first := to_add[0]
	for _, v := range to_add[1:] {
		add(first, v)
	}
	return magnitude(first)
}

func part2(to_add []*list.List) uint {
	max := uint(0)
	for _, v1 := range to_add {
		// need to copies since addition in SnailFish numbers is not commutative
		// and the add function modifies both
		for _, v2 := range to_add {
			v2cpy := list.New()
			v2cpy.PushBackList(v2)
			v1cpy := list.New()
			v1cpy.PushBackList(v1)

			add(v1cpy, v2cpy)
			if magn1 := magnitude(v1cpy); magn1 > max {
				max = magn1
			}

			v1cpy2 := list.New()
			v1cpy2.PushBackList(v1)
			v2cpy2 := list.New()
			v2cpy2.PushBackList(v2)
			add(v2cpy2, v1cpy2)

			if magn2 := magnitude(v2cpy2); magn2 > max {
				max = magn2
			}
		}
	}
	return max
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	p1in := make([]*list.List, 0)
	p2in := make([]*list.List, 0)

	for s.Scan() {
		t := s.Text()
		p1in = append(p1in, parse_snail_number(t))
		p2in = append(p2in, parse_snail_number(t))
	}

	fmt.Println("Part 1 solution:", part1(p1in))
	fmt.Println("Part 2 solution:", part2(p2in))
}
