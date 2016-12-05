package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

func main() {
	nodes, err := ParseNodes(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	if len(nodes) == 0 {
		return
	}
	min := math.MaxInt32
	test := make([]int, len(nodes))
	for i := 1; i < len(test); i++ {
		test[i] = i
	}
	for {
		sum := 0
		n := nodes[test[0]]
		for i := 1; i < len(test); i++ {
			o := nodes[test[i]]
			sum += n.Cost(o)
			n = o
		}
		if sum < min {
			min = sum
		}
		if !NextPermutation(sort.IntSlice(test)) {
			break
		}
	}
	fmt.Println(min)
}
