package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Rucksack = string

type Group [3]Rucksack

func score(c rune) int {
	if c >= 'a' && c <= 'z' {
		return int(c - 'a' + 1)
	}
	return int(c - 'A' + 27)
}

func main() {
	groups := []Group{}
	{
		c := 0
		var group Group
		bin := bufio.NewReader(os.Stdin)
		for {
			line, err := bin.ReadString('\n')
			if len(line) > 0 {
				line = strings.TrimRight(line, "\r\n")
				group[c] = line
				c++
				if c > 2 {
					c = 0
					groups = append(groups, group)
				}
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
	for _, g := range groups {
		for _, rs := range g {
			h := len(rs) / 2
			fc := rs[:h]
			sc := rs[h:]
			for _, c := range fc {
				if strings.ContainsRune(sc, c) {
					first += score(c)
					break
				}
			}
		}
		frs := g[0]
		for _, c := range frs {
			if strings.ContainsRune(g[1], c) && strings.ContainsRune(g[2], c) {
				second += score(c)
				break
			}
		}
	}
	fmt.Println("first:", first)
	fmt.Println("second:", second)
}
