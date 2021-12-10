package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var first, second int
	for scanner.Scan() {
		line := scanner.Text()
		pipe := strings.IndexByte(line, '|')
		in := strings.Split(line[:pipe-1], " ")
		out := strings.Split(line[pipe+2:], " ")
		form := decodeFormula(append(in, out...))
		op := form[len(in):]
		e := 1
		on := 0
		for i := len(op) - 1; i >= 0; i-- {
			d := op[i]
			if d == 1 || d == 4 || d == 7 || d == 8 {
				first++
			}
			on += e * d
			e *= 10
		}
		second += on
	}
	fmt.Println("first:", first)
	fmt.Println("second:", second)
}

func decodeFormula(s []string) []int {
	segs := parseSegments(s)
	abcdefg := Segments(0xFF >> 1)
	var cf, bcdf, acf Segments
	digits := map[Segments]int{}
	digits[abcdefg] = 8

	pot069 := map[Segments]bool{}
	for _, d := range segs {
		switch d.Count() {
		case 2:
			cf = d
		case 3:
			acf = d
		case 4:
			bcdf = d
		case 6:
			pot069[d] = true
		}
	}
	digits[cf] = 1
	digits[acf] = 7
	digits[bcdf] = 4

	bd := bcdf ^ cf

	var d, c, e Segments
	for p, _ := range pot069 {
		t := abcdefg ^ p
		if t&bd != 0 {
			d = t
			digits[p] = 0
		} else if t&acf != 0 {
			c = t
			digits[p] = 6
		} else {
			e = t
			digits[p] = 9
		}
	}

	b := bd ^ d
	digits[abcdefg^(b|e)] = 3
	digits[abcdefg^(c|e)] = 5
	f := cf ^ c
	digits[abcdefg^(b|f)] = 2

	form := make([]int, len(segs))
	for i, d := range segs {
		form[i] = digits[d]
	}
	return form
}

func parseSegments(s []string) []Segments {
	ba := make([]Segments, len(s))
	for i, de := range s {
		var ds Segments
		for _, c := range []byte(de) {
			ds |= (1 << (c - 'a'))
		}
		ba[i] = ds
	}
	return ba
}

type Segments uint8

func (s *Segments) Count() int {
	v := *s
	var c int
	for v != 0 {
		v &= v - 1
		c++
	}
	return c
}
