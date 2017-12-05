package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

func run(jumps []int, incr func(int) int) int {
	jm := make([]int, len(jumps))
	copy(jm, jumps)
	n := 0
	ip := 0
	for {
		nip := ip + jm[ip]
		jm[ip] = incr(jm[ip])
		ip = nip
		n++
		if ip >= len(jm) {
			return n
		}
	}
}

func first(jumps []int) int {
	return run(jumps, func(x int) int {
		return x + 1
	})
}

func second(jumps []int) int {
	return run(jumps, func(x int) int {
		if x >= 3 {
			return x - 1
		} else {
			return x + 1
		}
	})
}

func main() {
	jumps := []int{}
	bin := bufio.NewReader(os.Stdin)
	for {
		line, err := bin.ReadBytes('\n')
		if len(line) != 0 {
			x, err := strconv.Atoi(
				string(bytes.TrimSpace(line)))
			if err != nil {
				panic(err)
			}
			jumps = append(jumps, x)
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
	}
	fmt.Println(first(jumps))
	fmt.Println(second(jumps))
}
