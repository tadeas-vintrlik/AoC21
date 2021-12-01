package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

func hexa2bin_char(b byte) string {
	switch b {
	case '0':
		return "0000"
	case '1':
		return "0001"
	case '2':
		return "0010"
	case '3':
		return "0011"
	case '4':
		return "0100"
	case '5':
		return "0101"
	case '6':
		return "0110"
	case '7':
		return "0111"
	case '8':
		return "1000"
	case '9':
		return "1001"
	case 'A':
		return "1010"
	case 'B':
		return "1011"
	case 'C':
		return "1100"
	case 'D':
		return "1101"
	case 'E':
		return "1110"
	case 'F':
		return "1111"
	}

	return "INVALID"
}

func prepare_bin(s *string, left *string, desired int) string {
	r := ""
	l := 0

	if ll := len(*left); ll != 0 {
		r += *left
		l += ll
		*left = ""
	}

	for l < desired {
		r += hexa2bin_char((*s)[0])
		l += 4
		*s = (*s)[1:]
	}

	if l > desired {
		*left += r[desired:]
	}

	return r[:desired]
}

// Returns version and packet type
func get_header(s *string, left *string) (uint, uint) {
	local := prepare_bin(s, left, 6)

	v, _ := strconv.ParseInt(local[:3], 2, 64)
	t, _ := strconv.ParseInt(local[3:6], 2, 64)
	return uint(v), uint(t)
}

// Returns the parsed number and number of read bits
func get_number(s *string, left *string) (uint, uint) {
	r := ""
	no_bits := uint(0)
	end := false

	for !end {
		local := prepare_bin(s, left, 5)
		r += local[1:5]
		no_bits += 5
		if local[0] == '0' {
			end = true
		}
	}

	num, _ := strconv.ParseInt(r, 2, 64)
	return uint(num), no_bits
}

// Returns the parsed length of an operator packet if number of bits (number of packets otherwise)
// and number of bits read
func get_length(s *string, left *string) (uint, bool, uint) {
	local := prepare_bin(s, left, 1)
	l := 0
	if local[0] == '0' {
		l = 15
	} else {
		l = 11
	}
	local = prepare_bin(s, left, l)
	i, _ := strconv.ParseInt(local, 2, 64)
	// + 1 for the length type ID bit
	return uint(i), l == 15, uint(l + 1)
}

func process_operator_packet(type_id, no_packets uint, stack *list.List) {
	switch type_id {
	case 0: // Sum
		s := uint(0)
		for i := uint(0); i < no_packets; i++ {
			t := stack.Front()
			stack.Remove(t)
			s += t.Value.(uint)
		}
		stack.PushFront(s)
	case 1: // Product
		p := uint(1)
		for i := uint(0); i < no_packets; i++ {
			t := stack.Front()
			stack.Remove(t)
			p *= t.Value.(uint)
		}
		stack.PushFront(p)
	case 2: // Minimum
		m := -1
		for i := uint(0); i < no_packets; i++ {
			t := stack.Front()
			stack.Remove(t)
			v := t.Value.(uint)
			if m == -1 || uint(m) > v {
				m = int(v)
			}
		}
		stack.PushFront(uint(m))
	case 3: // Maximum
		m := -1
		for i := uint(0); i < no_packets; i++ {
			t := stack.Front()
			stack.Remove(t)
			v := t.Value.(uint)
			if m == -1 || uint(m) < v {
				m = int(v)
			}
		}
		stack.PushFront(uint(m))
	case 5: // Greater than
		t2 := stack.Front()
		stack.Remove(t2)
		t1 := stack.Front()
		stack.Remove(t1)
		v2 := t2.Value.(uint)
		v1 := t1.Value.(uint)
		if v1 > v2 {
			stack.PushFront(uint(1))
		} else {
			stack.PushFront(uint(0))
		}
	case 6: // Less than
		t2 := stack.Front()
		stack.Remove(t2)
		t1 := stack.Front()
		stack.Remove(t1)
		v2 := t2.Value.(uint)
		v1 := t1.Value.(uint)
		if v1 < v2 {
			stack.PushFront(uint(1))
		} else {
			stack.PushFront(uint(0))
		}
	case 7: // Equal to
		t2 := stack.Front()
		stack.Remove(t2)
		t1 := stack.Front()
		stack.Remove(t1)
		v2 := t2.Value.(uint)
		v1 := t1.Value.(uint)
		if v1 == v2 {
			stack.PushFront(uint(1))
		} else {
			stack.PushFront(uint(0))
		}
	default:
		fmt.Println("INVALID")
	}
}

// Returns number of read bytes
func process_packet(hex *string, left *string, v_sum *uint, stack *list.List) uint {
	r := uint(0)
	v, t := get_header(hex, left)
	*v_sum += v
	r += 6
	if t == 4 {
		num, bits := get_number(hex, left)
		stack.PushFront(num)
		r += bits
	} else { // Operator packet
		length, use_bit, bits := get_length(hex, left)
		r += bits

		no_packets := uint(0)
		if use_bit {
			i := uint(0)
			for i != length {
				i += process_packet(hex, left, v_sum, stack)
				no_packets++
			}
			r += i
		} else {
			// if length type ID was in number of packets
			no_packets = length
			for i := uint(0); i < length; i++ {
				r += process_packet(hex, left, v_sum, stack)
			}
		}
		process_operator_packet(t, no_packets, stack)
	}

	return r
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	hex := s.Text()
	left := ""
	version_sum := uint(0)
	stack := list.New()
	stack.Init()
	process_packet(&hex, &left, &version_sum, stack)
	fmt.Println("Part 1 solution:", version_sum)
	fmt.Println("Part 2 solution:", stack.Front().Value.(uint))
}
