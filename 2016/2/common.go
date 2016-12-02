package main

type Point struct {
	X, Y int
}

type Pad struct {
	Keys   []byte
	Width  int
	Height int
}

func (p *Pad) At(x, y int) byte {
	return p.Keys[y*p.Width+x]
}
