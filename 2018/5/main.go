package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
)

func main() {
	var source []byte

	{
		fd, err := os.Open("input.txt")
		if err != nil {
			panic(err)
		}
		defer fd.Close()
		raw, err := ioutil.ReadAll(fd)
		if err != nil {
			panic(err)
		}
		source = raw
	}

	{ // first
		polymer := make([]byte, 0, len(source))
		polymer = append(polymer[0:0], source...)
		for {
			bl := len(polymer)
			var o int
			for o < len(polymer)-1 {
				if polymer[o] == polymer[o+1]+0x20 ||
					polymer[o]+0x20 == polymer[o+1] {
					polymer = append(polymer[:o], polymer[o+2:]...)
				} else {
					o++
				}
			}
			if bl == len(polymer) {
				break
			}
		}
		fmt.Println("first:", len(polymer))
	}

	{ // second
		chars := map[byte]interface{}{}
		for _, c := range source {
			if c >= 0x41 && c <= 0x5A {
				if _, ok := chars[c]; !ok {
					chars[c] = true
				}
			}
		}
		tiniest := struct {
			Alpha byte
			Size  uint32
		}{
			Size: math.MaxUint32,
		}
		polymer := make([]byte, 0, len(source))
		for c, _ := range chars {
			polymer = append(polymer[0:0], source...)
			var o int
			for o < len(polymer) {
				if polymer[o] == c || polymer[o] == c+0x20 {
					polymer = append(polymer[:o], polymer[o+1:]...)
				} else {
					o++
				}
			}
			for {
				bl := len(polymer)
				o = 0
				for o < len(polymer)-1 {
					if polymer[o] == polymer[o+1]+0x20 ||
						polymer[o]+0x20 == polymer[o+1] {
						polymer = append(polymer[:o], polymer[o+2:]...)
					} else {
						o++
					}
				}
				if bl == len(polymer) {
					break
				}
			}
			if uint32(len(polymer)) < tiniest.Size {
				tiniest.Alpha = c
				tiniest.Size = uint32(len(polymer))
			}
		}
		fmt.Println("second:", tiniest.Size)
	}
}
