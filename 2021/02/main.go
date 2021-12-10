package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func main() {
	commands := []Command{}
	{
		reCommand := regexp.MustCompile(`(forward|up|down) (\d+)`)
		bin := bufio.NewReader(os.Stdin)
		for {
			line, err := bin.ReadString('\n')
			if len(line) > 0 {
				m := reCommand.FindStringSubmatch(line)
				if m == nil {
					panic("failed reading command")
				}
				units, err := strconv.Atoi(m[2])
				if err != nil {
					panic(err)
				}
				commands = append(commands, Command{
					Dir:   Dir(m[1]),
					Units: units,
				})
			}
			if err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			}
		}
	}
	var second Sub
	var first Point
	for _, cmd := range commands {
		switch cmd.Dir {
		case DirForward:
			first.X += cmd.Units
			second.Pos.X += cmd.Units
			second.Pos.Y += second.Aim * cmd.Units
		case DirUp:
			first.Y -= cmd.Units
			second.Aim -= cmd.Units
		case DirDown:
			first.Y += cmd.Units
			second.Aim += cmd.Units
		}
	}
	fmt.Println("first:", first.X*first.Y)
	fmt.Println("second:", second.Pos.X*second.Pos.Y)
}

type Sub struct {
	Pos Point
	Aim int
}

type Point struct {
	X, Y int
}

type Dir string

const (
	DirForward Dir = "forward"
	DirUp      Dir = "up"
	DirDown    Dir = "down"
)

type Command struct {
	Dir   Dir
	Units int
}
