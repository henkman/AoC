package main

import (
	"regexp"
	"sort"
	"strconv"

	"bufio"
	"fmt"
	"io"
	"os"
)

func first(sheet [][]int) int {
	sum := 0
	for _, line := range sheet {
		sort.Ints(line)
		sum += line[len(line)-1] - line[0]
	}
	return sum
}

func second(sheet [][]int) int {
	sum := 0
loop:
	for _, line := range sheet {
		sort.Ints(line)
		for i := len(line) - 1; i >= 0; i-- {
			for e := i - 1; e >= 0; e-- {
				if line[i]%line[e] == 0 {
					sum += line[i] / line[e]
					continue loop
				}
			}
		}
	}
	return sum
}

func main() {
	reValue := regexp.MustCompile(`\d+`)
	bin := bufio.NewReader(os.Stdin)
	sheet := make([][]int, 0, 10)
	for {
		line, err := bin.ReadSlice('\n')
		if len(line) != 0 {
			bvals := reValue.FindAll(line, -1)
			vals := make([]int, 0, len(bvals))
			for _, bval := range bvals {
				x, err := strconv.Atoi(string(bval))
				if err != nil {
					panic(err)
				}
				vals = append(vals, x)
			}
			sheet = append(sheet, vals)
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
	}
	fmt.Println(first(sheet))
	fmt.Println(second(sheet))
}
