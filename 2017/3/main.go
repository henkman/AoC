package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y int
}

func (p *Point) Add(op Point) {
	p.X += op.X
	p.Y += op.Y
}

func (p *Point) TurnLeft() {
	if p.X != 0 {
		p.Y = -p.X
		p.X = 0
	} else {
		p.X, p.Y = p.Y, p.X
	}
}

func (p *Point) Abs() int {
	return int(math.Abs(float64(p.X)) + math.Abs(float64(p.Y)))
}

func first(x int) int {
	d := Point{1, 0}
	n := 1
	turns := 0
	straight := 1
	p := Point{0, 0}
loop:
	for {
		for i := 1; i <= straight; i++ {
			p.Add(d)
			n++
			if n == x {
				break loop
			}
		}
		d.TurnLeft()
		turns++
		if turns == 2 {
			straight++
			turns = 0
		}
	}
	return p.Abs()
}

func second(x int) int {
	// TODO
	return 0
}

func main() {
	const INPUT = 361527
	fmt.Println(first(INPUT))
	fmt.Println(second(INPUT))
}
