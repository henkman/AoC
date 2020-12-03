package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var m Map
	{
		bin := bufio.NewReader(os.Stdin)
		for {
			line, err := bin.ReadString('\n')
			if len(line) > 0 {
				line = strings.TrimRight(line, "\r\n")
				m.Width = len(line)
				m.Height++
				m.Terrain = append(m.Terrain, []Terrain(line)...)
			}
			if err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			}
		}
	}

	type SlopeTest struct {
		Pos   Vector
		Slope Vector
		Trees int
	}

	tests := []SlopeTest{
		SlopeTest{
			Slope: Vector{X: 1, Y: 1},
		},
		SlopeTest{
			Slope: Vector{X: 3, Y: 1},
		},
		SlopeTest{
			Slope: Vector{X: 5, Y: 1},
		},
		SlopeTest{
			Slope: Vector{X: 7, Y: 1},
		},
		SlopeTest{
			Slope: Vector{X: 1, Y: 2},
		},
	}

	second := 1
	for i, _ := range tests {
		t := &tests[i]
		for t.Pos.Y < m.Height {
			if m.At(t.Pos.X, t.Pos.Y) == TerrainTrees {
				t.Trees++
			}
			t.Pos.Add(t.Slope)
			if t.Pos.X >= m.Width {
				t.Pos.X -= m.Width
			}
		}
		second *= t.Trees
	}

	fmt.Println("first:", tests[1].Trees)
	fmt.Println("second:", second)
}

type Vector struct {
	X, Y int
}

func (v *Vector) Add(o Vector) {
	v.X += o.X
	v.Y += o.Y
}

type Map struct {
	Terrain       []Terrain
	Width, Height int
}

func (m *Map) At(x, y int) Terrain {
	return m.Terrain[y*m.Width+x]
}

type Terrain = byte

const (
	TerrainOpen  Terrain = '.'
	TerrainTrees Terrain = '#'
)
