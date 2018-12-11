package main

import "fmt"

func main() {
	const W = 300
	const H = 300
	const GRID_SERIAL_INPUT = 3214

	grid := make([]int, W*H)
	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			o := y*W + x
			rackid := (x + 1) + 10
			powerlevel := rackid * (y + 1)
			powerlevel += GRID_SERIAL_INPUT
			powerlevel *= rackid
			powerlevel = (powerlevel % 1000) / 100
			powerlevel -= 5
			grid[o] = powerlevel
		}
	}

	{ // first
		const SW = 3
		const SH = 3
		var largest struct {
			X, Y int
			Sum  int
		}
		for y := 0; y < H-SH; y++ {
			for x := 0; x < W-SW; x++ {
				o := y*W + x
				sum := grid[o] + grid[o+1] + grid[o+2] +
					grid[o+W] + grid[o+W+1] + grid[o+W+2] +
					grid[o+W*2] + grid[o+W*2+1] + grid[o+W*2+2]
				if sum > largest.Sum {
					largest.Sum = sum
					largest.X = x
					largest.Y = y
				}
			}
		}
		fmt.Printf("first: %d,%d\n", largest.X+1, largest.Y+1)
	}

	{ // second, slow
		var largest struct {
			X, Y int
			Side int
			Sum  int
		}
		for side := 1; side < W; side++ {
			for y := 0; y < H-side; y++ {
				for x := 0; x < W-side; x++ {
					o := y*W + x
					sum := 0
					for ys := 0; ys < side; ys++ {
						for xs := 0; xs < side; xs++ {
							sum += grid[o+ys*W+xs]
						}
					}
					if sum > largest.Sum {
						largest.Sum = sum
						largest.X = x
						largest.Y = y
						largest.Side = side
					}
				}
			}
		}
		fmt.Printf("second: %d,%d,%d\n", largest.X+1, largest.Y+1, largest.Side)
	}

}
