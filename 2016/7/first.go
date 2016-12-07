package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func containsValidPalindrome(b []byte) bool {
	test := make([]byte, 0, 4)
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
			return true
		}
	}
	return false
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
		m := reInside.FindAllSubmatch(line, -1)
		if m != nil {
			for _, g := range m {
				if containsValidPalindrome(g[1]) {
					continue nextline
				}
			}
		}
		m = reOutside.FindAllSubmatch(line, -1)
		if m == nil {
			continue
		}
		for _, g := range m {
			if containsValidPalindrome(g[1]) {
				sum++
				continue nextline
			}
		}
	}
	fmt.Println(sum)
}
