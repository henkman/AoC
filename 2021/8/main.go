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
		input := strings.Split(line[:pipe-1], " ")
		output := strings.Split(line[pipe+2:], " ")
		form := decodeFormula(input, output)
		e := 1000
		on := 0
		for _, d := range form.Output {
			if d == 1 || d == 4 || d == 7 || d == 8 {
				first++
			}
			on += e * d
			e /= 10
		}
		second += on
	}
	fmt.Println("first:", first)
	fmt.Println("second:", second)
}

type Formula struct {
	Input  []int
	Output []int
}

func decodeFormula(input, output []string) Formula {
	indigit := parseSegments(input)
	outdigit := parseSegments(output)

	abcdefg := Segments(0xFF >> 1)
	var cf, bcdf, acf Segments
	digits := map[Segments]int{}
	digits[abcdefg] = 8

	pot069 := [3]Segments{}
	{
		mpot069 := map[Segments]bool{}
		for _, d := range indigit {
			switch d.Count() {
			case 2:
				cf = d
			case 3:
				acf = d
			case 4:
				bcdf = d
			case 6:
				mpot069[d] = true
			}
		}
		for _, d := range outdigit {
			switch d.Count() {
			case 2:
				cf = d
			case 3:
				acf = d
			case 4:
				bcdf = d
			case 6:
				mpot069[d] = true
			}
		}
		o := 0
		for d, _ := range mpot069 {
			pot069[o] = d
			o++
		}
	}
	digits[cf] = 1
	digits[acf] = 7
	digits[bcdf] = 4

	bd := bcdf ^ cf

	var d, c, e Segments
	for _, p := range pot069 {
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

	in := make([]int, len(input))
	for i, d := range indigit {
		in[i] = digits[d]
	}
	out := make([]int, len(output))
	for i, d := range outdigit {
		out[i] = digits[d]
	}

	return Formula{Input: in, Output: out}
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
