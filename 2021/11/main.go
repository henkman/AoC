package main

import (
	"bufio"
	"fmt"
	"os"
)

const SIZE = 10

func main() {

	grid := [SIZE * SIZE]int{}
	{
		o := 0
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Bytes()
			for _, c := range []byte(line) {
				grid[o] = int(c - '0')
				o++
			}
		}
	}

	first := 0
	second := 1
	for {
		flashes := Step(grid[:])
		first += flashes
		if second == 100 {
			fmt.Println("first:", first)
		}
		if flashes == 100 {
			fmt.Println("second: ", second)
			break
		}
		second++
	}
}

func Step(grid []int) int {
	for o := range grid {
		grid[o]++
	}

	flashed := [SIZE * SIZE]bool{}
	for y := 0; y < SIZE; y++ {
		for x := 0; x < SIZE; x++ {
			o := y*SIZE + x
			if grid[o] > 9 && !flashed[o] {
				Flash(o, x, y, grid, flashed[:])
			}
		}
	}

	flashes := 0
	for i, f := range flashed {
		if f {
			grid[i] = 0
			flashes++
		}
	}
	return flashes
}

func Flash(o, x, y int, grid []int, flashed []bool) {
	flashed[o] = true
	for _, d := range DIRS {
		ox := x + d.X
		oy := y + d.Y
		if ox >= 0 && ox < SIZE && oy >= 0 && oy < SIZE {
			oo := oy*SIZE + ox
			grid[oo]++
			if grid[oo] > 9 && !flashed[oo] {
				Flash(oo, ox, oy, grid, flashed)
			}
		}
	}
}

var (
	DIRS = []Vec{
		{X: -1, Y: -1},
		{X: 0, Y: -1},
		{X: 1, Y: -1},
		{X: 1, Y: 0},
		{X: 1, Y: 1},
		{X: 0, Y: 1},
		{X: -1, Y: 1},
		{X: -1, Y: 0},
	}
)

type Vec struct {
	X, Y int
}
