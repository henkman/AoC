package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Rect struct {
	X, Y, W, H uint32
}

func (r *Rect) Right() uint32 {
	return r.X + r.W
}

func (r *Rect) Bottom() uint32 {
	return r.Y + r.H
}

type Claim struct {
	ID uint32
	Rect
}

func readClaims() ([]Claim, error) {
	fd, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	input := make([]Claim, 0, 64)
	bin := bufio.NewReader(fd)
	for {
		line, err := bin.ReadString('\n')
		if len(line) > 0 {
			var c Claim
			fmt.Sscanf(strings.TrimSpace(line),
				"#%d @ %d,%d: %dx%d",
				&c.ID, &c.X, &c.Y, &c.W, &c.H)
			input = append(input, c)
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

func main() {
	claims, err := readClaims()
	if err != nil {
		panic(err)
	}

	{ // first
		var mw, mh uint32
		for _, claim := range claims {
			tw := claim.X + claim.W
			if tw > mw {
				mw = tw
			}
			th := claim.Y + claim.H
			if th > mh {
				mh = th
			}
		}

		bm := make([]uint32, mw*mh)
		for _, claim := range claims {
			h := claim.Y + claim.H
			w := claim.X + claim.W
			for y := claim.Y; y < h; y++ {
				for x := claim.X; x < w; x++ {
					o := x + (y * mw)
					bm[o]++
				}
			}
		}

		var n uint32
		for _, si := range bm {
			if si > 1 {
				n++
			}
		}
		fmt.Println("first:", n)
	}

	{ // second
		overlaps := func(a, b *Rect) bool {
			inrange := func(v, min, max uint32) bool {
				return v >= min && v <= max
			}
			xoverlap := inrange(a.X, b.X, b.X+b.W) ||
				inrange(b.X, a.X, a.X+a.W)
			yoverlap := inrange(a.Y, b.Y, b.Y+b.H) ||
				inrange(b.Y, a.Y, a.Y+a.H)
			return xoverlap && yoverlap
		}
	loop:
		for i, claim := range claims {
			for e, oclaim := range claims {
				if i == e {
					continue
				}

				if overlaps(&claim.Rect, &oclaim.Rect) {
					continue loop
				}
			}
			fmt.Println("second:", claim.ID)
			break
		}
	}
}
