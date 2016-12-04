package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	bin := bufio.NewReader(os.Stdin)
	for {
		line, err := bin.ReadString('\n')
		if line == "" {
			return
		}
		n := 0
		for _, c := range []byte(line) {
			if c == '(' {
				n++
			} else {
				n--
			}
		}
		fmt.Println(n)
		if err != nil {
			if err == io.EOF {
				return
			}
			log.Fatal(err)
		}
	}

}
