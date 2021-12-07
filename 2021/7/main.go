package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	crabs := []int{}
	for {
		var crab int
		_, err := fmt.Scan(&crab)
		if err != nil {
			break
		}
		crabs = append(crabs, crab)
	}
	sort.Ints(crabs)

	median := crabs[len(crabs)/2]
	fuel := 0
	for _, crab := range crabs {
		fuel += abs(crab - median)
	}
	fmt.Println("first:", fuel)

	min := crabs[0]
	max := crabs[len(crabs)-1]
	least := math.MaxInt32
	for s := min; s < max; s++ {
		fuel := 0
		for _, crab := range crabs {
			steps := abs(crab - s)
			fuel += triangular(steps)
		}
		if fuel < least {
			least = fuel
		}
	}
	fmt.Println("second:", least)
}

func triangular(n int) int {
	return n * (n + 1) / 2
}

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}
