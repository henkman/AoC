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
	adapters := []int{}
	{
		bin := bufio.NewReader(os.Stdin)
		for {
			line, err := bin.ReadString('\n')
			if len(line) > 0 {
				line = strings.TrimRight(line, "\r\n")
				adapter, err := strconv.Atoi(line)
				if err != nil {
					panic(err)
				}
				adapters = append(adapters, adapter)
			}
			if err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			}
		}
	}
	sort.Ints(adapters)
	{
		last := 0
		joltdiff := map[int]int{}
		for _, adapter := range adapters {
			jd := adapter - last
			joltdiff[jd]++
			last = adapter
		}
		joltdiff[3]++
		fmt.Println("first:", joltdiff[1]*joltdiff[3])
	}

}
