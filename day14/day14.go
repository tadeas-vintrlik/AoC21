package main

import (
	"bufio"
	"fmt"
	"os"
)

func pair_insert_step(p map[string]int, m map[string]string) map[string]int {
	r := make(map[string]int)
	for k, v := range p {
		n := m[k]
		r[string(k[0])+n] += v
		r[n+string(k[1])] += v
	}
	return r
}

func count_elements(p map[string]int, s string) map[byte]int {
	r := map[byte]int{'B': 0, 'C': 0, 'N': 0, 'H': 0}
	for k, v := range p {
		r[k[0]] += v
		r[k[1]] += v
	}

	// First and last character were the only ones not counted twice
	r[s[0]]++
	r[s[len(s)-1]]++

	// Divide all values by two since they were accounted for twice
	for k, v := range r {
		r[k] = v / 2
	}

	return r
}

func result(p map[string]int, s string) int {
	max := 0
	min := -1
	m := count_elements(p, s)
	for _, v := range m {
		if v > max {
			max = v
		}
		if v < min || min == -1 {
			min = v
		}
	}
	return max - min
}

func polymer_to_map(s string) map[string]int {
	m := make(map[string]int)
	l := len(s)
	for i := 2; i <= l; i++ {
		m[s[i-2:i]]++
	}
	return m
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	p := s.Text()

	ins := make(map[string]string)
	for s.Scan() {
		var f, t string
		fmt.Sscanf(s.Text(), "%s -> %s", &f, &t)
		ins[f] = t
	}

	m := polymer_to_map(p)
	for i := 0; i < 10; i++ {
		m = pair_insert_step(m, ins)
	}
	fmt.Println("Part 1 solution:", result(m, p))
	for i := 0; i < 30; i++ {
		m = pair_insert_step(m, ins)
	}
	fmt.Println("Part 2 solution:", result(m, p))
}
