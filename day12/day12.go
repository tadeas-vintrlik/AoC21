package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func is_small(s string) bool {
	return s[0] >= 'a' && s[0] <= 'z'
}

func no_visits(s string, r []string) uint {
	c := uint(0)
	for _, v := range r {
		if v == s {
			c++
		}
	}
	return c
}

// If this path is valid for part 2: one small cave visisted twice at most
func valid_p2(v []string) bool {
	two_visits := make([]string, 0)
	for _, x := range v {
		// If a small cave visisted more than once and not yet added to two_visits
		if is_small(x) && no_visits(x, v) > 1 && no_visits(x, two_visits) == 0 {
			two_visits = append(two_visits, x)
		}
	}

	if len(two_visits) > 1 {
		return false
	}

	if no_visits("start", v) == 2 {
		return false
	}

	return true
}

// c is the current cave
// m is the map of the caves
// v is the return list of all paths tried
// p is the part of the problem to solve
func try_path_r(c string, m map[string][]string, v [][]string, p uint) [][]string {
	// Remember index of the path before recursive calls
	i := len(v) - 1

	// Filter out invalid paths for part 2
	if p == 2 && !valid_p2(v[i]) {
		return v
	}

	for _, n := range m[c] {
		// Small caves can only be visited once
		if is_small(n) && no_visits(n, v[i]) == p {
			continue
		}
		// Create a copy of the current path for each node leading from c
		cpy := make([]string, len(v[i]))
		copy(cpy, v[i])
		cpy = append(cpy, n)
		v = append(v, cpy)
		// Stop recursion on end node
		if n == "end" {
			continue
		}
		v = try_path_r(n, m, v, p)
	}
	return v
}

// m is map of the caves
// p is part of the problem
func find_all_paths(m map[string][]string, p uint) [][]string {
	// Try and get all paths leading from start
	rs := make([][]string, 0)
	rs = append(rs, []string{"start"})
	rs = try_path_r("start", m, rs, p)

	// Filter out the ones that do not lead to end
	n := make([][]string, 0)
	for _, v := range rs {
		if no_visits("end", v) == 1 {
			n = append(n, v)
		}
	}

	return n
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	gs := make(map[string][]string) // Create a map of all paths indexed by starting cave
	for s.Scan() {
		t := strings.Split(s.Text(), "-")
		// Add starting point
		g := gs[t[0]]
		g = append(g, t[1])
		gs[t[0]] = g
		// Add end point
		g = gs[t[1]]
		g = append(g, t[0])
		gs[t[1]] = g
	}

	p1 := find_all_paths(gs, 1)
	fmt.Println("Part 1 solution:", len(p1))
	p2 := find_all_paths(gs, 2)
	fmt.Println("Part 2 solution:", len(p2))
}
