package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func fuelRequired(mass int) int {
	return int(math.Floor(float64(mass/3)) - 2)
}

func main() {
	first := 0
	second := 0
	bin := bufio.NewReader(os.Stdin)
	for {
		line, err := bin.ReadString('\n')
		if len(line) > 0 {
			line = strings.TrimRight(line, "\r\n")
			mass, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			fuel := fuelRequired(mass)
			first += fuel

			fuelsum := fuel
			for {
				nf := fuelRequired(fuel)
				if nf < 0 {
					break
				}
				fuelsum += nf
				fuel = nf
			}

			second += fuelsum
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
	}
	fmt.Println("first:", first)
	fmt.Println("second:", second)
}
