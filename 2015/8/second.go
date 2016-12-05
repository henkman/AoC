package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	bin := bufio.NewReader(os.Stdin)
	reString := regexp.MustCompile("[\"'\\\\]")
	memory := 0
	str := 0
	for {
		line, err := bin.ReadString('\n')
		if line == "" {
			break
		}
		line = strings.TrimSpace(line)
		str += len(line)
		s := "\"" + reString.ReplaceAllStringFunc(line, func(m string) string {
			return "\\" + m
		}) + "\""
		memory += len(s)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
	}
	fmt.Printf("%d-%d=%d\n", memory, str, memory-str)
}
