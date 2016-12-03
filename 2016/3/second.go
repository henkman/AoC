package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	const N = 3
	reNum := regexp.MustCompile("(\\d+ ?)+")
	bin := bufio.NewReader(os.Stdin)
	sum := 0
	o := 0
	tris := [N][3]int{}
	for {
		line, err := bin.ReadString('\n')
		m := reNum.FindAllStringSubmatch(line, -1)
		for i := 0; i < N; i++ {
			n, err := strconv.Atoi(strings.TrimSpace(m[i][1]))
			if err != nil {
				log.Fatal(err)
			}
			tris[i][o] = n
		}
		o++
		if o == N {
			o = 0
			for i := 0; i < N; i++ {
				if isPossibleTriangle(tris[i][0], tris[i][1], tris[i][2]) {
					sum++
				}
			}
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
