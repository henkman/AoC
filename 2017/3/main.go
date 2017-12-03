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
	return int(math.Abs(float64(p.X)) +
		math.Abs(float64(p.Y)))
}

func (p *Point) Neighbour(op Point) bool {
	return int(math.Abs(float64(p.X-op.X))) <= 1 &&
		int(math.Abs(float64(p.Y-op.Y))) <= 1
}

func first(x int) int {
	d := Point{1, 0}
	n := 1
	turns := 0
	straight := 1
	p := Point{0, 0}
	for {
		for i := 1; i <= straight; i++ {
			p.Add(d)
			n++
			if n == x {
				return p.Abs()
			}
		}
		d.TurnLeft()
		turns++
		if turns == 2 {
			straight++
			turns = 0
		}
	}
}

func second(x int) int {
	type ValPoint struct {
		Point
		Val int
	}
	d := Point{1, 0}
	turns := 0
	straight := 1
	p := Point{0, 0}
	points := []ValPoint{
		ValPoint{Point{0, 0}, 1},
	}
	for {
		for i := 1; i <= straight; i++ {
			p.Add(d)
			v := 0
			for _, op := range points {
				if p.Neighbour(op.Point) {
					v += op.Val
				}
			}
			if v > x {
				return v
			}
			points = append(points, ValPoint{p, v})
		}
		d.TurnLeft()
		turns++
		if turns == 2 {
			straight++
			turns = 0
		}
	}
}

func main() {
	const INPUT = 361527
	fmt.Println(first(INPUT))
	fmt.Println(second(INPUT))
}
