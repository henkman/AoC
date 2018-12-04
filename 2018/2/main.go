package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func readInput() ([]string, error) {
	fd, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	input := make([]string, 0, 64)
	bin := bufio.NewReader(fd)
	for {
		line, err := bin.ReadString('\n')
		if len(line) > 0 {
			line = strings.TrimSpace(line)
			input = append(input, line)
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return input, nil
}

func main() {
	input, err := readInput()
	if err != nil {
		panic(err)
	}

	{ // first
		idxs := map[int]int{}
		for _, line := range input {
			seen := map[byte]int{}
			for _, c := range []byte(line) {
				if _, ok := seen[c]; ok {
					seen[c]++
				} else {
					seen[c] = 1
				}
			}
			nums := map[int]int{}
			for _, n := range seen {
				if n == 1 {
					continue
				}
				if _, ok := nums[n]; !ok {
					nums[n] = 1
				}
			}
			for n, _ := range nums {
				idxs[n]++
			}
		}
		fmt.Println("first:", idxs[2]*idxs[3])
	}

	{ // second
		hasOneDifference := func(a, b string) bool {
			diff := false
			for i, _ := range []byte(a) {
				if a[i] != b[i] {
					if diff {
						return false
					} else {
						diff = true
					}
				}
			}
			return diff
		}

		intersection := func(a, b string) string {
			var sb strings.Builder
			for i, _ := range []byte(a) {
				if a[i] == b[i] {
					sb.WriteByte(a[i])
				}
			}
			return sb.String()
		}

	loop:
		for i, line := range input {
			for e, ol := range input {
				if i == e {
					continue
				}
				if hasOneDifference(line, ol) {
					fmt.Println(intersection(line, ol))
					break loop
				}
			}
		}
	}
}
