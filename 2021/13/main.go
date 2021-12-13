package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	insts := []Instruction{}
	var paper Paper
	{
		dots := []Vec{}
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.HasPrefix(line, "fold") {
				var c byte
				var v int
				fmt.Sscanf(line, "fold along %c=%d\n", &c, &v)
				insts = append(insts, Instruction{
					FoldUp: c == 'y',
					Val:    v,
				})
				continue
			}
			if line == "" {
				continue
			}
			var vec Vec
			fmt.Sscanf(line, "%d,%d\n", &vec.X, &vec.Y)
			dots = append(dots, vec)
		}
		var mx, my int
		for _, dot := range dots {
			if dot.X > mx {
				mx = dot.X
			}
			if dot.Y > my {
				my = dot.Y
			}
		}
		mx++
		my++
		paper = Paper{
			Dots: make([]bool, mx*my),
			W:    mx,
			H:    my,
		}
		for _, dot := range dots {
			paper.Dots[dot.Y*paper.W+dot.X] = true
		}
	}

	for i, inst := range insts {
		if inst.FoldUp {
			nh := paper.H - inst.Val - 1
			bottom := paper.Dots[(nh+1)*paper.W:]
			for y := 0; y < nh; y++ {
				for x := 0; x < paper.W; x++ {
					o := y*paper.W + x
					paper.Dots[o] = paper.Dots[o] || bottom[(nh-y-1)*paper.W+x]
				}
			}
			paper.H = nh
		} else {
			nw := paper.W - inst.Val - 1
			for y := 0; y < paper.H; y++ {
				for x := 0; x < nw; x++ {
					paper.Dots[y*nw+x] = paper.Dots[y*paper.W+x]
				}
				for x := 0; x < nw; x++ {
					o := y*nw + x
					paper.Dots[o] = paper.Dots[o] || paper.Dots[(y+1)*paper.W-x-1]
				}
			}
			paper.W = nw
		}

		if i == 0 {
			first := 0
			for _, d := range paper.Dots[:paper.W*paper.H] {
				if d {
					first++
				}
			}
			fmt.Println("first:", first)
		}
	}
	paper.Print()
}

type Paper struct {
	Dots []bool
	W, H int
}

func (p *Paper) Print() {
	for y := 0; y < p.H; y++ {
		for x := 0; x < p.W; x++ {
			if p.Dots[y*p.W+x] {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

type Vec struct {
	X, Y int
}

type Instruction struct {
	FoldUp bool
	Val    int
}
