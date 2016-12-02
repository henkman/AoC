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
	pad.Width = 5
	pad.Keys = []byte("  1   234 56789 ABC   D  ")
	pad.Height = len(pad.Keys) / pad.Width
	w := Point{0, 2}
	bin := bufio.NewReader(os.Stdin)
	for {
		line, err := bin.ReadString('\n')
		for _, ins := range line {
			switch ins {
			case 'U':
				if w.Y-1 >= 0 && pad.At(w.X, w.Y-1) != ' ' {
					w.Y -= 1
				}
			case 'D':
				if w.Y+1 < pad.Height && pad.At(w.X, w.Y+1) != ' ' {
					w.Y += 1
				}
			case 'L':
				if w.X-1 >= 0 && pad.At(w.X-1, w.Y) != ' ' {
					w.X -= 1
				}
			case 'R':
				if w.X+1 < pad.Width && pad.At(w.X+1, w.Y) != ' ' {
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
