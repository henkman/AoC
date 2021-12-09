package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var hm HMap
	{
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			hm.W = len(line)
			for _, c := range []byte(line) {
				hm.Locs = append(hm.Locs, int(c-'0'))
			}
			hm.H++
			hm.Visited = make([]bool, hm.W*hm.H)
		}
	}
	{
		var first int
		for y := 0; y < hm.H; y++ {
		next:
			for x := 0; x < hm.W; x++ {
				p := hm.At(x, y)
				for _, dir := range DIRS {
					tx := x + dir.X
					ty := y + dir.Y
					op := hm.At(tx, ty)
					if p >= op {
						continue next
					}
				}
				first += p + 1
			}
		}
		fmt.Println("first:", first)
	}
	{
		basins := []int{}
		for y := 0; y < hm.H; y++ {
			for x := 0; x < hm.W; x++ {
				if hm.At(x, y) == 9 || hm.WasVisited(x, y) {
					continue
				}
				basin := hm.CountBasinSize(x, y)
				basins = append(basins, basin)
			}
		}
		sort.Ints(basins)
		second := basins[len(basins)-3]
		for _, b := range basins[len(basins)-2:] {
			second *= b
		}
		fmt.Println("second:", second)
	}
}

var (
	DIRS = []Vec{
		{X: 0, Y: -1},
		{X: 0, Y: 1},
		{X: -1, Y: 0},
		{X: 1, Y: 0},
	}
)

type Vec struct {
	X, Y int
}

type HMap struct {
	Locs    []int
	Visited []bool
	W, H    int
}

func (hm *HMap) WasVisited(x, y int) bool {
	if x >= hm.W || x < 0 || y >= hm.H || y < 0 {
		return true
	}
	return hm.Visited[y*hm.W+x]
}

func (hm *HMap) At(x, y int) int {
	if x >= hm.W || x < 0 || y >= hm.H || y < 0 {
		return 9
	}
	return hm.Locs[y*hm.W+x]
}

func (hm *HMap) CountBasinSize(x, y int) int {
	basin := 1
	hm.Visited[y*hm.W+x] = true
	for _, dir := range DIRS {
		tx := x + dir.X
		ty := y + dir.Y
		if hm.At(tx, ty) == 9 || hm.WasVisited(tx, ty) {
			continue
		}
		basin += hm.CountBasinSize(tx, ty)
	}
	return basin
}
