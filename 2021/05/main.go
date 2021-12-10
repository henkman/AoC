package main

import (
	"fmt"
)

func main() {
	vents := []Line{}
	for {
		var vent Line
		if _, err := fmt.Scanf("%d,%d -> %d,%d\n",
			&vent.Begin.X, &vent.Begin.Y,
			&vent.End.X, &vent.End.Y); err != nil {
			break
		}
		vents = append(vents, vent)
	}

	var max Point
	for _, vent := range vents {
		if vent.Begin.X > max.X {
			max.X = vent.Begin.X
		}
		if vent.End.X > max.X {
			max.X = vent.End.X
		}
		if vent.Begin.Y > max.Y {
			max.Y = vent.Begin.Y
		}
		if vent.End.Y > max.Y {
			max.Y = vent.End.Y
		}
	}
	max.X++
	max.Y++

	fd := make([]int, max.X*max.Y)
	sd := make([]int, max.X*max.Y)
	for _, vent := range vents {
		renderOnlyHorizontalAndVertical(vent.Begin, vent.End, max, fd)
		renderLines(vent.Begin, vent.End, max, sd)
	}

	var first, second int
	for i, _ := range fd {
		if fd[i] > 1 {
			first++
		}
		if sd[i] > 1 {
			second++
		}
	}
	fmt.Println("first:", first)
	fmt.Println("second:", second)
}

func renderLines(b, e Point, max Point, diagram []int) {
	var step Point
	if b.X < e.X {
		step.X++
	} else if b.X > e.X {
		step.X--
	}
	if b.Y < e.Y {
		step.Y++
	} else if b.Y > e.Y {
		step.Y--
	}
	p := b
	for {
		o := p.Y*max.X + p.X
		diagram[o]++
		if p.X == e.X && p.Y == e.Y {
			break
		}
		p.X += step.X
		p.Y += step.Y
	}
}

func renderOnlyHorizontalAndVertical(b, e Point, max Point, diagram []int) {
	if b.Y == e.Y {
		if e.X < b.X {
			e, b = b, e
		}
		for x := b.X; x <= e.X; x++ {
			o := b.Y*max.X + x
			diagram[o]++
		}
	} else if b.X == e.X {
		if e.Y < b.Y {
			e, b = b, e
		}
		for y := b.Y; y <= e.Y; y++ {
			o := y*max.X + b.X
			diagram[o]++
		}
	}
}

type Line struct {
	Begin, End Point
}

type Point struct {
	X, Y int
}
