package main

import (
	"fmt"
	"io"
	"os"
)

func findMarker(buf []byte, mlen int) int {
next:
	for o := 0; o < len(buf)-mlen; o++ {
		end := o + mlen
		for i := o; i < end; i++ {
			for e := o; e < end; e++ {
				if i == e {
					continue
				}
				if buf[i] == buf[e] {
					continue next
				}
			}
		}
		return o + mlen
	}
	return -1
}

func main() {
	buf, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	fmt.Println("first:", findMarker(buf, 4))
	fmt.Println("second:", findMarker(buf, 14))
}
