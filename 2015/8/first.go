package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	bin := bufio.NewReader(os.Stdin)
	reString := regexp.MustCompile("\\\\([\"'\\\\]|x[0-9a-f]{2})")
	memory := 0
	str := 0
	for {
		line, err := bin.ReadString('\n')
		if line == "" {
			break
		}
		line = line[:len(line)-1]
		memory += len(line)
		line = line[1 : len(line)-1]
		s := reString.ReplaceAllStringFunc(line, func(m string) string {
			if m[1] == 'x' {
				v, err := strconv.ParseUint(m[2:], 16, 32)
				if err != nil {
					log.Fatal(err)
				}
				return fmt.Sprintf("%c", v%128)
			}
			return m[1:]
		})
		str += len(s)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
	}
	fmt.Printf("%d-%d=%d\n", memory, str, memory-str)
}
