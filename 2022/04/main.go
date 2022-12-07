package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Pair struct {
	First, Second Assignment
}

type Assignment struct {
	Low  int
	High int
}

func (ass *Assignment) Contains(o Assignment) bool {
	return ass.Low <= o.Low && ass.High >= o.High
}

func (ass *Assignment) Overlaps(o Assignment) bool {
	return (ass.Low <= o.High && ass.High >= o.Low)
}

func main() {
	pairs := []Pair{}
	{
		bin := bufio.NewReader(os.Stdin)
		for {
			line, err := bin.ReadString('\n')
			if len(line) > 0 {
				line = strings.TrimRight(line, "\r\n")
				var pair Pair
				fmt.Sscanf(line, "%d-%d,%d-%d",
					&pair.First.Low, &pair.First.High,
					&pair.Second.Low, &pair.Second.High)
				pairs = append(pairs, pair)
			}
			if err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			}
		}
	}
	var first, second int
	for _, p := range pairs {
		if p.First.Contains(p.Second) || p.Second.Contains(p.First) {
			first++
		}
		if p.First.Overlaps(p.Second) {
			second++
		}
	}
	fmt.Println("first:", first)
	fmt.Println("second:", second)
}
