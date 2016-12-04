package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func min(a, b int) int {
	if b > a {
		return a
	}
	return b
}

func main() {
	bin := bufio.NewReader(os.Stdin)
	sum := 0
	for {
		var x, y, z int
		_, err := fmt.Fscanf(bin, "%dx%dx%d\n", &x, &y, &z)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		a := x * y
		b := y * z
		c := x * z
		sum += 2*a + 2*b + 2*c
		sum += min(a, min(b, c))
	}
	fmt.Println(sum)
}
