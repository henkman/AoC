package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func getValidPalindromes(b []byte) []string {
	ps := make([]string, 0, 2)
	test := make([]byte, 0, 3)
	for _, c := range b {
		if len(test) == cap(test) {
			for i := 0; i < len(test)-1; i++ {
				test[i] = test[i+1]
			}
			test[len(test)-1] = c
		} else {
			test = append(test, c)
		}
		if len(test) == cap(test) && test[0] != test[1] && isPalindrome(test) {
			ps = append(ps, string(test))
		}
	}
	return ps
}

func main() {
	bin := bufio.NewReader(os.Stdin)
	sum := 0
nextline:
	for {
		line, _ := bin.ReadBytes('\n')
		if len(line) == 0 {
			break
		}
		line = bytes.TrimSpace(line)
		is := []string{}
		m := reInside.FindAllSubmatch(line, -1)
		if m != nil {
			for _, g := range m {
				is = append(is, getValidPalindromes(g[1])...)
			}
		}
		m = reOutside.FindAllSubmatch(line, -1)
		if m == nil {
			continue
		}
		os := []string{}
		for _, g := range m {
			os = append(os, getValidPalindromes(g[1])...)
		}
		for _, s := range is {
			for _, o := range os {
				if s[0] == o[1] && s[1] == o[0] && s[1] == o[2] {
					sum++
					continue nextline
				}
			}
		}
	}
	fmt.Println(sum)
}
