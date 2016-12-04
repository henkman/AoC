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
		for i, c := range []byte(line) {
			if c == '(' {
				n++
			} else {
				n--
			}
			if n == -1 {
				fmt.Println(i + 1)
				break
			}
		}
		if err != nil {
			if err == io.EOF {
				return
			}
			log.Fatal(err)
		}
	}

}
