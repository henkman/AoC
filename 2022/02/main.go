package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Match struct {
	First  byte
	Second byte
}

func (m *Match) ScoreFirst() int {
	shape := int(m.Second - 'X' + 1)
	won := m.First == 'A' && m.Second == 'Y' ||
		m.First == 'B' && m.Second == 'Z' ||
		m.First == 'C' && m.Second == 'X'
	if won {
		return 6 + shape
	}
	draw := m.First == 'A' && m.Second == 'X' ||
		m.First == 'B' && m.Second == 'Y' ||
		m.First == 'C' && m.Second == 'Z'
	if draw {
		return 3 + shape
	}
	return shape
}

func (m *Match) ScoreSecond() int {
	win := m.Second == 'Z'
	if win {
		switch m.First {
		case 'A':
			return 6 + 2
		case 'B':
			return 6 + 3
		case 'C':
			return 6 + 1
		}
	}
	lose := m.Second == 'X'
	if lose {
		switch m.First {
		case 'A':
			return 3
		case 'B':
			return 1
		case 'C':
			return 2
		}
	}
	return 3 + int(m.First-'A'+1)
}

func main() {
	matches := []Match{}
	{
		bin := bufio.NewReader(os.Stdin)
		for {
			line, err := bin.ReadString('\n')
			if len(line) > 0 {
				line = strings.TrimRight(line, "\r\n")
				var m Match
				fmt.Sscanf(line, "%c %c", &m.First, &m.Second)
				matches = append(matches, m)
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
	for _, m := range matches {
		first += m.ScoreFirst()
		second += m.ScoreSecond()
	}
	fmt.Println("first:", first)
	fmt.Println("second:", second)
}
