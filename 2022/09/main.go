package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type Direction = byte

const (
	Direction_Up    Direction = 'U'
	Direction_Down  Direction = 'D'
	Direction_Left  Direction = 'L'
	Direction_Right Direction = 'R'
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sgn(x int) int {
	if x < 0 {
		return -1
	} else if x > 0 {
		return 1
	}
	return 0
}

type Vector struct {
	X, Y int
}

func (v *Vector) Add(o Vector) Vector {
	return Vector{v.X + o.X, v.Y + o.Y}
}

func (v *Vector) Sub(o Vector) Vector {
	return Vector{v.X - o.X, v.Y - o.Y}
}

func (v *Vector) Sgn() Vector {
	return Vector{sgn(v.X), sgn(v.Y)}
}

func (v *Vector) Distance(o Vector) float64 {
	x := float64(v.X - o.X)
	y := float64(v.Y - o.Y)
	return math.Sqrt(x*x + y*y)
}

type Instruction struct {
	Direction Direction
	Steps     int
}

type Rope struct {
	Knots []Vector
}

func (r Rope) Tail() Vector {
	return r.Knots[len(r.Knots)-1]
}

func (r Rope) Length() int {
	return len(r.Knots)
}

func (r Rope) PullInDirection(d Vector) {
	r.Knots[0] = r.Knots[0].Add(d)
	for i := 1; i < r.Length(); i++ {
		prev := &r.Knots[i-1]
		cur := &r.Knots[i]
		dist := (*prev).Distance(*cur)
		if dist > math.Sqrt2 {
			diff := (*prev).Sub(*cur)
			sgn := diff.Sgn()
			*cur = (*cur).Add(sgn)
		}
	}
}

func main() {
	instrs := []Instruction{}
	{
		bin := bufio.NewReader(os.Stdin)
		for {
			line, err := bin.ReadString('\n')
			if len(line) > 0 {
				line = strings.TrimRight(line, "\r\n")
				tokens := strings.Split(line, " ")
				steps, err := strconv.Atoi(tokens[1])
				if err != nil {
					panic(err)
				}
				instrs = append(instrs, Instruction{tokens[0][0], steps})
			}
			if err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			}
		}
	}
	var (
		firstVisited  = map[Vector]int{}
		firstRope     = Rope{Knots: make([]Vector, 2)}
		secondVisited = map[Vector]int{}
		secondRope    = Rope{Knots: make([]Vector, 10)}
	)
	for _, instr := range instrs {
		v := map[Direction]Vector{
			Direction_Up:    {0, -1},
			Direction_Down:  {0, 1},
			Direction_Left:  {-1, 0},
			Direction_Right: {1, 0},
		}[instr.Direction]
		for step := 0; step < instr.Steps; step++ {
			firstRope.PullInDirection(v)
			firstVisited[firstRope.Tail()]++
			secondRope.PullInDirection(v)
			secondVisited[secondRope.Tail()]++
		}
	}
	fmt.Println("first:", len(firstVisited))
	fmt.Println("second:", len(secondVisited))
}
