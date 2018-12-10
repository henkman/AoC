package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"math"
	"os"
	"path/filepath"
	"strings"
)

type Vec2 struct {
	X, Y int
}

func (v Vec2) Add(o Vec2) Vec2 {
	return Vec2{
		X: v.X + o.X,
		Y: v.Y + o.Y,
	}
}

type Light struct {
	Position Vec2
	Velocity Vec2
}

func main() {
	lights := make([]Light, 0, 32)
	{
		fd, err := os.Open("input.txt")
		if err != nil {
			panic(err)
		}
		defer fd.Close()
		bin := bufio.NewReader(fd)
		for {
			line, err := bin.ReadString('\n')
			if len(line) > 0 {
				line = strings.TrimSpace(line)
				var l Light
				fmt.Sscanf(line,
					"position=<%d,%d> velocity=<%d,%d>",
					&l.Position.X, &l.Position.Y,
					&l.Velocity.X, &l.Velocity.Y)
				lights = append(lights, l)
			}
			if err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			}
		}
	}

	red := color.RGBA{0xFF, 0x00, 0x00, 0xFF}
	for step := 0; step < 500000; step++ {
		mw := math.MaxInt32
		mh := math.MaxInt32
		var w, h int
		for i, _ := range lights {
			lights[i].Position = lights[i].Position.Add(lights[i].Velocity)
			p := lights[i].Position
			if p.X >= 0 && p.X < mw {
				mw = p.X
			}
			if p.X > w {
				w = p.X
			}
			if p.Y >= 0 && p.Y < mh {
				mh = p.Y
			}
			if p.Y > h {
				h = p.Y
			}
		}
		if w > len(lights) && h > len(lights) {
			continue
		}
		frame := image.NewRGBA(image.Rect(0, 0, w-mw+1, h-mh+1))
		for _, light := range lights {
			x := light.Position.X
			y := light.Position.Y
			if x < 0 || y < 0 {
				continue
			}
			frame.Set(int(x-mw), int(y-mh), red)
		}
		{
			os.Mkdir("frames", 0750)
			fd, err := os.OpenFile(
				filepath.Join("frames", fmt.Sprintf("frame_%d.png", step)),
				os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0750)
			if err != nil {
				panic(err)
			}
			err = png.Encode(fd, frame)
			if err != nil {
				panic(err)
			}
		}
	}

}
