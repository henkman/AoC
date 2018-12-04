package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func readInput() ([]int64, error) {
	fd, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	input := make([]int64, 0, 64)
	bin := bufio.NewReader(fd)
	for {
		line, err := bin.ReadString('\n')
		if len(line) > 0 {
			line = strings.TrimSpace(line)
			v, err := strconv.ParseInt(line, 10, 64)
			if err != nil {
				return nil, err
			}
			input = append(input, v)
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
		var freq int64
		for _, v := range input {
			freq += v
		}
		fmt.Println("first:", freq)
	}

	{ // second
		contains := func(haystack []int64, needle int64) bool {
			for _, v := range haystack {
				if v == needle {
					return true
				}
			}
			return false
		}
		var freq int64
		seen := make([]int64, 0, 64)
	loop:
		for {
			for _, v := range input {
				freq += v
				if contains(seen, freq) {
					fmt.Println("second:", freq)
					break loop
				} else {
					seen = append(seen, freq)
				}
			}
		}
	}
}
