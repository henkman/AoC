package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var first int
	scores := []int{}
next:
	for scanner.Scan() {
		line := scanner.Bytes()
		stack := []byte{}
		for _, c := range line {
			switch c {
			case '(':
				stack = append(stack, c+1)
			case '[', '{', '<':
				stack = append(stack, c+2)
			case ')', ']', '}', '>':
				if c != stack[len(stack)-1] {
					switch c {
					case ')':
						first += 3
					case ']':
						first += 57
					case '}':
						first += 1197
					case '>':
						first += 25137
					}
					continue next
				} else {
					stack = stack[:len(stack)-1]
				}
			}
		}
		score := 0
		for i := len(stack) - 1; i >= 0; i-- {
			switch stack[i] {
			case ')':
				score = score*5 + 1
			case ']':
				score = score*5 + 2
			case '}':
				score = score*5 + 3
			case '>':
				score = score*5 + 4
			}
		}
		scores = append(scores, score)
	}
	fmt.Println("first:", first)
	sort.Ints(scores)
	fmt.Println("second:", scores[len(scores)/2])
}
