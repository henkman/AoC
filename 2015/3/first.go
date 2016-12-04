package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type Point struct {
	X, Y int
}

func main() {
	bin := bufio.NewReader(os.Stdin)
	for {
		line, err := bin.ReadString('\n')
		if line == "" {
			break
		}
		houses := map[Point]int{}
		var w Point
		houses[w] = 1
		for _, c := range []byte(line) {
			switch c {
			case '^':
				w.Y -= 1
			case 'v':
				w.Y += 1
			case '>':
				w.X += 1
			case '<':
				w.X -= 1
			}
			if p, ok := houses[w]; ok {
				houses[w] = p + 1
			} else {
				houses[w] = 1
			}
		}
		fmt.Println(len(houses))
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
	}
}
