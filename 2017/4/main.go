package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func valid(lines [][][]byte, test func(a, b []byte) bool) int {
	v := 0
loop:
	for _, line := range lines {
		for i := 0; i < len(line); i++ {
			for e := i + 1; e < len(line); e++ {
				if len(line[i]) == len(line[e]) &&
					test(line[i], line[e]) {
					continue loop
				}
			}
		}
		v++
	}
	return v
}

func first(lines [][][]byte) int {
	return valid(lines, bytes.Equal)
}

func second(lines [][][]byte) int {
	return valid(lines, func(a, b []byte) bool {
		oca := ['z' + 1]int{}
		for _, c := range a {
			oca[c]++
		}
		ocb := ['z' + 1]int{}
		for _, c := range b {
			ocb[c]++
		}
		for i := 'a'; i <= 'z'; i++ {
			if oca[i] != ocb[i] {
				return false
			}
		}
		return true
	})
}

func main() {
	lines := [][][]byte{}
	bin := bufio.NewReader(os.Stdin)
	for {
		line, err := bin.ReadBytes('\n')
		if len(line) != 0 {
			lines = append(lines,
				bytes.Split(bytes.TrimSpace(line), []byte(" ")))
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
	}
	fmt.Println(first(lines))
	fmt.Println(second(lines))
}
