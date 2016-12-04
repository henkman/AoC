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
		o := 0
		var ss [2]Point
		houses[ss[0]] = 1
		for _, c := range []byte(line) {
			s := &ss[o]
			switch c {
			case '^':
				s.Y -= 1
			case 'v':
				s.Y += 1
			case '>':
				s.X += 1
			case '<':
				s.X -= 1
			}
			if p, ok := houses[*s]; ok {
				houses[*s] = p + 1
			} else {
				houses[*s] = 1
			}
			if o == 1 {
				o = 0
			} else {
				o = 1
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
