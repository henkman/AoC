package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func isNice(s string) bool {
	for _, b := range []string{"ab", "cd", "pq", "xy"} {
		if strings.Contains(s, b) {
			return false
		}
	}
	hasRow := false
	vs := 0
	var last byte
	for _, c := range []byte(s) {
		if c == last {
			hasRow = true
		}
		if strings.ContainsRune("aeiou", rune(c)) {
			vs++
		}
		last = c
	}
	return vs >= 3 && hasRow
}

func main() {
	bin := bufio.NewReader(os.Stdin)
	sum := 0
	for {
		line, err := bin.ReadString('\n')
		if line == "" {
			break
		}
		if isNice(line) {
			sum++
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
	}
	fmt.Println(sum)
}
