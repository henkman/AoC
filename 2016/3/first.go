package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	bin := bufio.NewReader(os.Stdin)
	sum := 0
	for {
		var a, b, c int
		_, err := fmt.Fscanf(bin, "%d %d %d\n", &a, &b, &c)
		if err == io.EOF {
			break
		}
		if isPossibleTriangle(a, b, c) {
			sum++
		}
	}
	fmt.Println(sum)
}
