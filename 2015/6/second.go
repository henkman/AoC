package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

type Point struct {
	X, Y int
}

func main() {
	const W = 1000
	const H = 1000
	lights := make([]int, W*H)
	bin := bufio.NewReader(os.Stdin)
	reInst := regexp.MustCompile("(turn on|toggle|turn off) (\\d+,\\d+) through (\\d+,\\d+)")
	for {
		line, err := bin.ReadString('\n')
		if line == "" {
			break
		}
		m := reInst.FindStringSubmatch(line)
		var tl, br Point
		fmt.Sscanf(m[2], "%d,%d", &tl.X, &tl.Y)
		fmt.Sscanf(m[3], "%d,%d", &br.X, &br.Y)
		switch m[1] {
		case "turn on":
			for y := tl.Y; y <= br.Y; y++ {
				for x := tl.X; x <= br.X; x++ {
					o := y*W + x
					lights[o] += 1
				}
			}
		case "turn off":
			for y := tl.Y; y <= br.Y; y++ {
				for x := tl.X; x <= br.X; x++ {
					o := y*W + x
					n := lights[o] - 1
					if n >= 0 {
						lights[o] = n
					}
				}
			}
		case "toggle":
			for y := tl.Y; y <= br.Y; y++ {
				for x := tl.X; x <= br.X; x++ {
					o := y*W + x
					lights[o] += 2
				}
			}
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
	}
	sum := 0
	for _, l := range lights {
		sum += l
	}
	fmt.Println(sum)
}
