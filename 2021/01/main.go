package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	entries := []int{}
	{
		bin := bufio.NewReader(os.Stdin)
		for {
			line, err := bin.ReadString('\n')
			if len(line) > 0 {
				line = strings.TrimRight(line, "\r\n")
				entry, err := strconv.Atoi(line)
				if err != nil {
					panic(err)
				}
				entries = append(entries, entry)
			}
			if err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			}
		}
	}
	first := 0
	prev := entries[0]
	for i := 1; i < len(entries); i++ {
		cur := entries[i]
		if cur > prev {
			first++
		}
		prev = cur
	}
	fmt.Println("first:", first)

	const WINDOW = 3
	second := 0
	for i := 1; i < len(entries)-WINDOW+1; i++ {
		prev := entries[i-1 : i-1+WINDOW]
		cur := entries[i : i+WINDOW]
		if sum(cur) > sum(prev) {
			second++
		}
	}
	fmt.Println("second:", second)
}

func sum(arr []int) int {
	s := 0
	for _, e := range arr {
		s += e
	}
	return s
}
