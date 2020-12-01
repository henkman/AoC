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
	var first, second int
	for _, a := range entries {
		for _, b := range entries {
			if a+b == 2020 {
				first = a * b
				continue
			}
			for _, c := range entries {
				if a+b+c == 2020 {
					second = a * b * c
				}
			}
		}
	}
	fmt.Println(first, second)
}
