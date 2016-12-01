package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {
	const (
		W = 2000
		H = 2000
	)
	img := image.NewRGBA(image.Rect(0, 0, W, H))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.RGBA{0, 0, 0, 0xFF}}, image.ZP, draw.Src)
	raw, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	var w Point
	dir := Dir_North
	for _, m := range reInstruction.FindAllStringSubmatch(string(raw), -1) {
		dir = dir.Turn(byte(m[1][0]) == 'L')
		n, err := strconv.Atoi(m[2])
		if err != nil {
			log.Fatal(err)
		}
		path := w.WalkStraightPath(dir, n)
		for _, p := range path {
			img.Set(p.X+W/2, p.Y+H/2, color.RGBA{0xFF, 0, 0, 0xFF})
		}
		w = path[len(path)-1]
	}
	{
		fd, err := os.OpenFile("path.png", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0750)
		if err != nil {
			log.Fatal(err)
		}
		defer fd.Close()
		if err := png.Encode(fd, img); err != nil {
			log.Fatal(err)
		}
	}
}
