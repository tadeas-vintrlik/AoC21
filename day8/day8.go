package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Structure with mappings of signal letters to digits
type digits struct {
	table [10]string // Table used to decode index is digit containing one of the permutations of
	// signal letters
}

func (d digits) String() string {
	s := "Mapping of current 7-segment:\n"
	for i := 0; i < 10; i++ {
		t := fmt.Sprintf("%d : %s\n", i, d.table[i])
		s += t
	}
	return s
}

func no_matching_chars(s1 string, s2 string) int {
	r := 0
	l1 := len(s1)
	l2 := len(s1)
	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			if s1[i] == s2[j] {
				r++
			}
		}
	}

	return r
}

func filter_out_shared(s []string) []string {
	l := len(s[0])
	for i := 0; i < l; i++ {
		i1 := strings.Index(s[1], string(s[0][i]))
		i2 := strings.Index(s[2], string(s[0][i]))
		if i1 != -1 && i2 != -1 {
			// Remove the shared segment, length changed and must check the same index again
			s[0] = s[0][:i] + s[0][i+1:]
			s[1] = s[1][:i1] + s[1][i1+1:]
			s[2] = s[2][:i2] + s[2][i2+1:]
			l--
			i--
		}
	}
	return s
}

// Takes the part of the input before delimiter and creates a corresponding 7digit structure
func create_digits(s string) digits {
	r := digits{}

	all := strings.Split(s, " ")
	l5 := make([]string, 0)
	l6 := make([]string, 0)
	for _, v := range all {
		// First try if we can determine just by length
		d := decode_unique(v)
		if d != 0 {
			r.table[d] = v
			continue
		}
		// The rest should be strings of length 5 or 6 split them accordingly
		if len(v) == 5 {
			l5 = append(l5, v)
		} else {
			l6 = append(l6, v)
		}
	}

	// Create copies to mutate the strings and determine which number is on each index
	c5 := make([]string, len(l5))
	c6 := make([]string, len(l6))
	copy(c5, l5)
	copy(c6, l6)

	// Each number has 3 shared segments, remove them to make it easier to determine
	c5 = filter_out_shared(c5)

	// Each number has 4 shared segments, remove them to make it easier to determine
	c6 = filter_out_shared(c6)

	f := false
	for i := 0; i < 3 && !f; i++ {
		for j := 0; j < 3 && !f; j++ {
			if no_matching_chars(c5[i], c6[j]) == 2 {
				// Found 0 in l5 and 2 in l6 since they share same segments
				r.table[2] = l5[i]
				r.table[0] = l6[j]
				c5 = append(c5[:i], c5[i+1:]...)
				l5 = append(l5[:i], l5[i+1:]...)
				c6 = append(c6[:j], c6[j+1:]...)
				l6 = append(l6[:j], l6[j+1:]...)
				f = true
			}
		}
	}

	f = false
	for i := 0; i < 2 && !f; i++ {
		m := 0
		for j := 0; j < 2 && !f; j++ {
			m += no_matching_chars(c5[i], c6[j])
		}
		if m == 0 {
			// We found number 5 since it shares no segments, the other is 3
			r.table[5] = l5[i]
			i3 := i - 1
			if i3 < 0 {
				i3 = 1
			}
			r.table[3] = l5[i3]

			// 3 shares one segment with number 9
			if no_matching_chars(c6[0], c5[i3]) == 1 {
				r.table[9] = l6[0]
				r.table[6] = l6[1]
			} else {
				r.table[9] = l6[1]
				r.table[6] = l6[0]
			}
			f = true
		}
	}

	return r
}

// Decode just numbers with unique number of signals on 7-digit display (1, 4, 7, 8)
// Returns 0 if none of the above
func decode_unique(s string) int {
	n := 0
	switch len(s) {
	case 2:
		n = 1
		break
	case 4:
		n = 4
		break
	case 3:
		n = 7
		break
	case 7:
		n = 8
		break
	}
	return n
}

func decode_number(s string, d digits) int {
	l := len(s)
	for i := 0; i < 10; i++ {
		m := 0
		if len(d.table[i]) != l {
			continue
		}
		for j := 0; j < l; j++ {
			if strings.ContainsRune(d.table[i], rune(s[j])) {
				m++
			}
		}
		if m == l {
			return i
		}
	}
	return -1
}

func main() {
	s1 := 0
	s2 := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " | ")

		// Part 2
		d := create_digits(s[0])

		o := strings.Split(s[1], " ")
		e := 1000 // The first digit is thousands
		for _, v := range o {
			// Part 1
			if decode_unique(v) != 0 {
				s1++
			}
			s2 += e * decode_number(v, d)
			e /= 10
		}
	}

	fmt.Println("Part 1 solution:", s1)
	fmt.Println("Part 2 solution:", s2)
}
