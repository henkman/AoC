package main

import (
	"fmt"
)

func main() {
	fish := [9]int{}
	for {
		var v int
		if _, err := fmt.Scan(&v); err != nil {
			break
		}
		fish[v]++
	}
	for i := 0; i < 256; i++ {
		if i == 80 {
			s := 0
			for _, f := range fish {
				s += f
			}
			fmt.Println("first:", s)
		}

		nf := fish[0]
		for e := 0; e < 8; e++ {
			fish[e] = fish[e+1]
		}
		fish[8] = nf
		fish[6] += nf
	}
	s := 0
	for _, f := range fish {
		s += f
	}
	fmt.Println("second:", s)
}
