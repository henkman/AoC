package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var pad Pad
	pad.Width = 3
	pad.Keys = []byte("123456789")
	pad.Height = len(pad.Keys) / pad.Width
	w := Point{1, 1}
	bin := bufio.NewReader(os.Stdin)
	for {
		line, err := bin.ReadString('\n')
		for _, ins := range line {
			switch ins {
			case 'U':
				if w.Y-1 >= 0 {
					w.Y -= 1
				}
			case 'D':
				if w.Y+1 < pad.Height {
					w.Y += 1
				}
			case 'L':
				if w.X-1 >= 0 {
					w.X -= 1
				}
			case 'R':
				if w.X+1 < pad.Width {
					w.X += 1
				}
			}
		}
		fmt.Printf("%c", pad.At(w.X, w.Y))
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
	}
}
