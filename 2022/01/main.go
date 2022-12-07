package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	elves := []int{}
	{
		bin := bufio.NewReader(os.Stdin)
		var elf int
		for {
			line, err := bin.ReadString('\n')
			line = strings.TrimRight(line, "\r\n")
			if len(line) > 0 {
				cals, err := strconv.Atoi(line)
				if err != nil {
					panic(err)
				}
				elf += cals
			} else {
				elves = append(elves, elf)
				elf = 0
			}
			if err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			}
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(elves)))
	fmt.Println("first:", elves[0])
	second := 0
	for i := 0; i < 3; i++ {
		second += elves[i]
	}
	fmt.Println("second:", second)
}
