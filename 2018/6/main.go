package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

type Point struct {
	X, Y int32
}

func readPoints() ([]Point, error) {
	fd, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	input := make([]Point, 0, 64)
	bin := bufio.NewReader(fd)
	for {
		line, err := bin.ReadString('\n')
		if len(line) > 0 {
			var p Point
			fmt.Sscanf(strings.TrimSpace(line), "%d, %d", &p.X, &p.Y)
			input = append(input, p)
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return input, nil
}

func abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

type Area struct {
	Size     int32
	Infinite bool
}

func main() {
	points, err := readPoints()
	if err != nil {
		panic(err)
	}

	var mw, mh int32
	for _, p := range points {
		if p.X > mw {
			mw = p.X
		}
		if p.Y > mh {
			mh = p.Y
		}
	}

	{ // first
		areas := make([]Area, len(points))
		for y := int32(0); y <= mh; y++ {
			for x := int32(0); x <= mw; x++ {
				closest := struct {
					Index    int32
					Distance int32
					Tied     bool
				}{
					Distance: math.MaxInt32,
				}
				for i, p := range points {
					d := abs(p.X-x) + abs(p.Y-y)
					if d < closest.Distance {
						closest.Index = int32(i)
						closest.Distance = d
						closest.Tied = false
					} else if d == closest.Distance {
						closest.Index = int32(i)
						closest.Distance = d
						closest.Tied = true
					}
				}
				if !closest.Tied {
					if x == 0 || x == mw || y == 0 || y == mh {
						areas[closest.Index].Infinite = true
					}
					areas[closest.Index].Size++
				}
			}
		}
		var largest int32
		for _, area := range areas {
			if area.Infinite {
				continue
			}
			if area.Size > largest {
				largest = area.Size
			}
		}
		fmt.Println("first:", largest)
	}

	{ // second
		const DIST = 10000
		var size int32
		for y := int32(0); y <= mh; y++ {
			for x := int32(0); x <= mw; x++ {
				var tot int32
				for _, p := range points {
					d := abs(p.X-x) + abs(p.Y-y)
					tot += d
				}
				if tot < DIST {
					size++
				}
			}
		}
		fmt.Println("second:", size)
	}
}
