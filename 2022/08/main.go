package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Tree = int8

type Visibility = byte

const (
	Visiblity_Left   Visibility = 1
	Visiblity_Right  Visibility = 1 << 1
	Visiblity_Top    Visibility = 1 << 2
	Visiblity_Bottom Visibility = 1 << 3
)

type TreeMap struct {
	Trees          []Tree
	VisibilityMap  []Visibility
	ScenicScoreMap []uint
	Width, Height  int
}

func (tm *TreeMap) At(x, y int) Tree {
	if x < 0 || x >= tm.Width || y < 0 || y >= tm.Height {
		return -1
	}
	return tm.Trees[tm.Width*y+x]
}

func (tm *TreeMap) AddVisibility(x, y int, v Visibility) {
	tm.VisibilityMap[tm.Width*y+x] |= v
}

func (tm *TreeMap) Visibility(x, y int) Visibility {
	return tm.VisibilityMap[tm.Width*y+x]
}

func (tm *TreeMap) IsVisibleFrom(x, y int, v Visibility) bool {
	return tm.VisibilityMap[tm.Width*y+x]&v != 0
}

func (tm *TreeMap) SetScenicScore(x, y int, score uint) {
	tm.ScenicScoreMap[tm.Width*y+x] = score
}

func (tm *TreeMap) ScenicScore(x, y int) uint {
	return tm.ScenicScoreMap[tm.Width*y+x]
}

func (tm *TreeMap) CalculateVisibilityMap() {
	w := tm.Width
	h := tm.Height
	tm.VisibilityMap = make([]Visibility, w*h)
	for y := 0; y < h; y++ {
		{ // left to right
			var ht Tree = -1
			for x := 0; x < w; x++ {
				p := tm.At(x-1, y)
				c := tm.At(x, y)
				if c > p && c > ht {
					tm.AddVisibility(x, y, Visiblity_Left)
				}
				if c > ht {
					ht = c
				}
			}
		}
		{ // right to left
			var ht Tree = -1
			for x := w - 1; x >= 0; x-- {
				p := tm.At(x+1, y)
				c := tm.At(x, y)
				if c > p && c > ht {
					tm.AddVisibility(x, y, Visiblity_Right)
				}
				if c > ht {
					ht = c
				}
			}
		}
	}
	for x := 0; x < w; x++ {
		{ // top to bottom
			var ht Tree = -1
			for y := 0; y < h; y++ {
				p := tm.At(x, y-1)
				c := tm.At(x, y)
				if c > p && c > ht {
					tm.AddVisibility(x, y, Visiblity_Top)
				}
				if c > ht {
					ht = c
				}
			}
		}
		{ // bottom to top
			var ht Tree = -1
			for y := h - 1; y >= 0; y-- {
				p := tm.At(x, y+1)
				c := tm.At(x, y)
				if c > p && c > ht {
					tm.AddVisibility(x, y, Visiblity_Bottom)
				}
				if c > ht {
					ht = c
				}
			}
		}
	}
}

type Vec struct {
	X, Y int
}

func (v *Vec) Add(o Vec) Vec {
	return Vec{X: v.X + o.X, Y: v.Y + o.Y}
}

func (tm *TreeMap) CalculateScenicScoreMap() {
	w := tm.Width
	h := tm.Height
	tm.ScenicScoreMap = make([]uint, w*h)
	dirs := []Vec{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			var dsa [4]uint
			t := tm.At(x, y)
			for i, dir := range dirs {
				cur := Vec{x, y}
				var ds uint
				for {
					cur = cur.Add(dir)
					nt := tm.At(cur.X, cur.Y)
					if nt == -1 {
						break
					}
					ds++
					if nt >= t {
						break
					}
				}
				dsa[i] = ds
			}
			var score uint = dsa[0]
			for i := 1; i < len(dsa); i++ {
				score *= dsa[i]
			}
			tm.SetScenicScore(x, y, score)
		}
	}
}

func main() {
	var tm TreeMap
	{
		bin := bufio.NewReader(os.Stdin)
		for {
			line, err := bin.ReadString('\n')
			if len(line) > 0 {
				line = strings.TrimRight(line, "\r\n")
				tm.Width = len(line)
				ts := make([]int8, len(line))
				for i := 0; i < tm.Width; i++ {
					ts[i] = int8(line[i] - '0')
				}
				tm.Trees = append(tm.Trees, ts...)
			}
			if err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			}
		}
		tm.Height = len(tm.Trees) / tm.Width
	}
	tm.CalculateVisibilityMap()
	first := 0
	for _, t := range tm.VisibilityMap {
		if t != 0 {
			first++
		}
	}
	fmt.Println("first:", first)

	tm.CalculateScenicScoreMap()
	var second uint
	for _, t := range tm.ScenicScoreMap {
		if t > second {
			second = t
		}
	}
	fmt.Println("second:", second)
}
