package main

import (
	"fmt"
	"io"
	"os"
)

type Banks []int

func (bs *Banks) String() string {
	s := ""
	for _, c := range *bs {
		s += fmt.Sprint(c)
	}
	return s
}

func main() {
	var banks Banks
	for {
		var v int
		_, err := fmt.Fscanf(os.Stdin, "%d", &v)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		banks = append(banks, v)
	}
	cycles := 0
	seen := map[string]int{}
	for {
		var max struct {
			n int
			o int
		}
		for i, _ := range banks {
			if banks[i] > max.n {
				max.n = banks[i]
				max.o = i
			}
		}
		s := max.n / len(banks)
		r := max.n % len(banks)
		banks[max.o] = 0
		o := (max.o + 1) % len(banks)
		for i := 0; i < r; i++ {
			banks[o]++
			o = (o + 1) % len(banks)
		}
		for i, _ := range banks {
			banks[i] += s
		}
		cycles++
		cs := banks.String()
		if sb, ok := seen[cs]; ok {
			fmt.Println("first:", cycles)
			fmt.Println("second:", cycles-sb)
			break
		}
		seen[cs] = cycles
	}
}
