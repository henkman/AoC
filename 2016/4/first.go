package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	bin := bufio.NewReader(os.Stdin)
	sum := 0
	for {
		line, err := bin.ReadString('\n')
		if line == "" {
			break
		}
		m := reLine.FindStringSubmatch(line)
		if isValidRoom([]byte(m[1]), []byte(m[3])) {
			id, err := strconv.Atoi(m[2])
			if err != nil {
				log.Fatal(err)
			}
			sum += id
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
