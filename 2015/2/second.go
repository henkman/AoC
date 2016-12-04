package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

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
		d := []int{x, y, z}
		sort.Ints(d)
		sum += d[0]*2 + d[1]*2
		sum += x * y * z
	}
	fmt.Println(sum)
}
