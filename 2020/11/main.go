package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	prev := Map{}
	{
		bin := bufio.NewReader(os.Stdin)
		for {
			line, err := bin.ReadString('\n')
			if len(line) > 0 {
				line = strings.TrimRight(line, "\r\n")
				prev.Width = len(line)
				prev.Height++
				prev.Position = append(prev.Position, []Position(line)...)
			}
			if err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			}
		}
	}

	cur := Map{
		Position: make([]Position, len(prev.Position)),
		Width:    prev.Width,
		Height:   prev.Height,
	}
	copy(cur.Position, prev.Position)
	for {
		for y := 0; y < cur.Height; y++ {
			for x := 0; x < cur.Width; x++ {
				sp := cur.At(x, y)
				p := prev.At(x, y)
				occ := prev.AdjacentOccupiedSeats(x, y)
				if *p == PositionEmptySeat && occ == 0 {
					*sp = PositionOccupiedSeat
				} else if *p == PositionOccupiedSeat && occ >= 4 {
					*sp = PositionEmptySeat
				}
			}
		}
		// fmt.Println(cur.String())
		if bytes.Equal(prev.Position, cur.Position) {
			first := 0
			for _, p := range cur.Position {
				if p == PositionOccupiedSeat {
					first++
				}
			}
			fmt.Println("first:", first)
			break
		}
		copy(prev.Position, cur.Position)
	}
}

type Position = byte

const (
	PositionFloor        Position = '.'
	PositionEmptySeat    Position = 'L'
	PositionOccupiedSeat Position = '#'
)

type Map struct {
	Position      []Position
	Width, Height int
}

func (m *Map) String() string {
	var sb strings.Builder
	o := 0
	for y := 0; y < m.Height; y++ {
		for x := 0; x < m.Width; x++ {
			sb.WriteByte(m.Position[o])
			o++
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

func (m *Map) At(x, y int) *Position {
	if x >= m.Width || x < 0 || y >= m.Height || y < 0 {
		return nil
	}
	return &m.Position[y*m.Width+x]
}

func (m *Map) AdjacentOccupiedSeats(x, y int) int {
	s := 0
	if p := m.At(x-1, y-1); p != nil && *p == PositionOccupiedSeat {
		s++
	}
	if p := m.At(x, y-1); p != nil && *p == PositionOccupiedSeat {
		s++
	}
	if p := m.At(x+1, y-1); p != nil && *p == PositionOccupiedSeat {
		s++
	}
	if p := m.At(x+1, y); p != nil && *p == PositionOccupiedSeat {
		s++
	}
	if p := m.At(x+1, y+1); p != nil && *p == PositionOccupiedSeat {
		s++
	}
	if p := m.At(x, y+1); p != nil && *p == PositionOccupiedSeat {
		s++
	}
	if p := m.At(x-1, y+1); p != nil && *p == PositionOccupiedSeat {
		s++
	}
	if p := m.At(x-1, y); p != nil && *p == PositionOccupiedSeat {
		s++
	}
	return s
}
