package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var pwcs []PasswordCheck
	{
		bin := bufio.NewReader(os.Stdin)
		for {
			line, err := bin.ReadString('\n')
			if len(line) > 0 {
				line = strings.TrimRight(line, "\r\n")
				var pwc PasswordCheck
				fmt.Sscanf(line, "%d-%d %c: %s",
					&pwc.Min, &pwc.Max, &pwc.Letter, &pwc.Pass)
				pwcs = append(pwcs, pwc)
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
	for _, pwc := range pwcs {
		if pwc.IsValidOld() {
			first++
		}
		if pwc.IsValidNew() {
			second++
		}
	}
	fmt.Println("first:", first)
	fmt.Println("second:", second)
}

type PasswordCheck struct {
	Min    int
	Max    int
	Letter byte
	Pass   string
}

func (pwc *PasswordCheck) IsValidOld() bool {
	l := string([]byte{pwc.Letter})
	c := strings.Count(pwc.Pass, l)
	return c >= pwc.Min && c <= pwc.Max
}

func (pwc *PasswordCheck) IsValidNew() bool {
	b := []byte(pwc.Pass)
	min := b[pwc.Min-1] == pwc.Letter
	max := b[pwc.Max-1] == pwc.Letter
	return min != max
}
