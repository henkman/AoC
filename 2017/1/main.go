package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func first(s []byte) string {
	var prev byte
	sum := 0
	for _, c := range s {
		if c == prev {
			sum += int(c - '0')
		}
		prev = c
	}
	if prev == s[0] {
		sum += int(prev - '0')
	}
	return fmt.Sprint(sum)
}

func second(s []byte) string {
	sum := 0
	for i, c := range s {
		o := (i + (len(s) / 2)) % len(s)
		if c == s[o] {
			sum += int(c - '0')
		}
	}
	return fmt.Sprint(sum)
}

func main() {
	s, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	fmt.Println(first(s))
	fmt.Println(second(s))
}
