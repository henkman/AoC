package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	groups := []Group{}
	{
		bin := bufio.NewReader(os.Stdin)
		var group Group
		for {
			line, err := bin.ReadString('\n')
			line = strings.TrimRight(line, "\r\n")
			if len(line) > 0 {
				answers := Answers{}
				for _, a := range line {
					answers[byte(a)] = true
				}
				group.Answers = append(group.Answers, answers)
			} else {
				groups = append(groups, group)
				group = Group{}
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
	for _, group := range groups {
		first += group.CountGroupAnswers()
		second += group.CountGroupConsensus()
	}
	fmt.Println("first:", first)
	fmt.Println("second:", second)
}

type Answers map[byte]bool

type Group struct {
	Answers []Answers
}

func (g *Group) CountGroupAnswers() int {
	c := map[byte]bool{}
	for _, answers := range g.Answers {
		for b, _ := range answers {
			c[b] = true
		}
	}
	return len(c)
}

func (g *Group) CountGroupConsensus() int {
	c := 0
next:
	for i := 'a'; i <= 'z'; i++ {
		for _, answers := range g.Answers {
			if _, ok := answers[byte(i)]; !ok {
				continue next
			}
		}
		c++
	}
	return c
}
