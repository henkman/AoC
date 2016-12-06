package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	lines := make([][]byte, 0, 10)
	bin := bufio.NewReader(os.Stdin)
	for {
		line, err := bin.ReadString('\n')
		if line == "" {
			break
		}
		lines = append(lines, []byte(strings.TrimSpace(line)))
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
	}
	cols := len(lines[0])
	for i := 0; i < cols; i++ {
		occ := ['z' + 1]int{}
		for _, l := range lines {
			occ[l[i]]++
		}
		var c, mc byte
		max := 0
		for c = 'a'; c <= 'z'; c++ {
			if occ[c] > max {
				max = occ[c]
				mc = c
			}
		}
		fmt.Printf("%c", mc)
	}

}
