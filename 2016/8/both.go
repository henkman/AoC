package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	const W = 50
	const H = 6
	var pixels [W * H]bool
	offset := func(x, y int) int {
		return y*W + x
	}
	printPixels := func() {
		for y := 0; y < H; y++ {
			for x := 0; x < W; x++ {
				if pixels[offset(x, y)] {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}
	rotateRow := func(y, n int) {
		for i := 0; i < n; i++ {
			t := pixels[offset(W-1, y)]
			for x := W - 1; x > 0; x-- {
				pixels[offset(x, y)] = pixels[offset(x-1, y)]
			}
			pixels[offset(0, y)] = t
		}
	}
	rotateColumn := func(x, n int) {
		for i := 0; i < n; i++ {
			t := pixels[offset(x, H-1)]
			for y := H - 1; y > 0; y-- {
				pixels[offset(x, y)] = pixels[offset(x, y-1)]
			}
			pixels[offset(x, 0)] = t
		}
	}
	bin := bufio.NewReader(os.Stdin)
	for {
		line, _ := bin.ReadString('\n')
		if len(line) == 0 {
			break
		}
		if strings.HasPrefix(line, "rect") {
			var w, h int
			fmt.Sscanf(line, "rect %dx%d\n", &w, &h)
			fmt.Printf("rect %dx%d\n", w, h)
			for y := 0; y < h; y++ {
				for x := 0; x < w; x++ {
					pixels[y*W+x] = true
				}
			}
		} else if strings.HasPrefix(line, "rotate row") {
			var y, o int
			fmt.Sscanf(line, "rotate row y=%d by %d\n", &y, &o)
			fmt.Printf("rotate row y=%d by %d\n", y, o)
			rotateRow(y, o)
		} else if strings.HasPrefix(line, "rotate column") {
			var x, o int
			fmt.Sscanf(line, "rotate column x=%d by %d\n", &x, &o)
			fmt.Printf("rotate column x=%d by %d\n", x, o)
			rotateColumn(x, o)
		}
		printPixels()
	}
	sum := 0
	for _, p := range pixels {
		if p {
			sum++
		}
	}
	fmt.Println("pixels lit:", sum)
}
