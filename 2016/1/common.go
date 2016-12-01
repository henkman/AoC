package main

import (
	"math"
	"regexp"
)

var (
	reInstruction = regexp.MustCompile("([LR])(\\d+)(?:, )?")
)

type Dir uint8

const (
	Dir_North Dir = iota
	Dir_West
	Dir_East
	Dir_South
)

func (d *Dir) Turn(left bool) Dir {
	switch *d {
	case Dir_North:
		if left {
			return Dir_West
		} else {
			return Dir_East
		}
	case Dir_South:
		if left {
			return Dir_East
		} else {
			return Dir_West
		}
	case Dir_West:
		if left {
			return Dir_South
		} else {
			return Dir_North
		}
	case Dir_East:
		if left {
			return Dir_North
		} else {
			return Dir_South
		}
	}
	panic("unreachable")
}

type Point struct {
	X, Y int
}

func (p *Point) WalkStraight(d Dir, n int) Point {
	switch d {
	case Dir_North:
		return Point{p.X, p.Y + -n}
	case Dir_South:
		return Point{p.X, p.Y + n}
	case Dir_West:
		return Point{p.X + -n, p.Y}
	case Dir_East:
		return Point{p.X + n, p.Y}
	}
	panic("unreachable")
}

func (p *Point) WalkStraightPath(d Dir, n int) []Point {
	path := make([]Point, n)
	for i := 0; i < n; i++ {
		switch d {
		case Dir_North:
			path[i] = Point{p.X, p.Y + -(i + 1)}
		case Dir_South:
			path[i] = Point{p.X, p.Y + (i + 1)}
		case Dir_West:
			path[i] = Point{p.X + -(i + 1), p.Y}
		case Dir_East:
			path[i] = Point{p.X + (i + 1), p.Y}
		}
	}
	return path
}

func (p *Point) Distance(o Point) int {
	return int(math.Abs(float64(p.X)-float64(o.X)) + math.Abs(float64(p.Y)-float64(o.Y)))
}
