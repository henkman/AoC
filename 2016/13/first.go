package main

import (
	"fmt"
	"math"

	astar "github.com/beefsack/go-astar"
)

func bitsSet(v uint) uint {
	if v < 256 {
		return [256]uint{
			0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4,
			1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
			1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
			2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
			1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
			2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
			2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
			3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
			1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
			2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
			2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
			3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
			2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
			3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
			3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
			4, 5, 5, 6, 5, 6, 6, 7, 5, 6, 6, 7, 6, 7, 7, 8,
		}[v]
	}
	var c uint
	for ; v != 0; v >>= 1 {
		c += v & 1
	}
	return c
}

type Tile struct {
	X, Y int
	Off  uint
}

func (t Tile) PathNeighborCost(to astar.Pather) float64 {
	return 1
}

func (t Tile) PathEstimatedCost(to astar.Pather) float64 {
	dx := to.(Tile).X - t.X
	dy := to.(Tile).Y - t.Y
	return math.Sqrt(float64(dx*dx + dy*dy))
}

func (t Tile) PathNeighbors() []astar.Pather {
	ns := make([]astar.Pather, 0, 4)
	check := func(tx, ty int) {
		x := t.X + tx
		y := t.Y + ty
		if x < 0 || y < 0 {
			return
		}
		b := uint(x*x+3*x+2*x*y+y+y*y) + t.Off
		ot := Tile{x, y, t.Off}
		bits := bitsSet(b)
		if bits&1 == 0 {
			ns = append(ns, ot)
		}
	}
	check(-1, 0)
	check(0, -1)
	check(1, 0)
	check(0, 1)
	return ns
}

func main() {
	const N = 1358
	_, d, ok := astar.Path(Tile{1, 1, N}, Tile{31, 39, N})
	if ok {
		fmt.Println(d)
	}
}
