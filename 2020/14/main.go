package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	assignments := []Assignment{}
	{
		var mask Mask
		reMask := regexp.MustCompile(`mask = ([X10]+)`)
		reAssignment := regexp.MustCompile(`mem\[(\d+)\] = (\d+)`)
		bin := bufio.NewReader(os.Stdin)
		for {
			line, err := bin.ReadString('\n')
			if len(line) > 0 {
				line = strings.TrimRight(line, "\r\n")
				m := reMask.FindStringSubmatch(line)
				if m != nil {
					mask = parseMask(m[1])
					continue
				}
				m = reAssignment.FindStringSubmatch(line)
				address, err := strconv.ParseUint(m[1], 10, 64)
				if err != nil {
					panic(err)
				}
				value, err := strconv.ParseUint(m[2], 10, 64)
				if err != nil {
					panic(err)
				}
				assignments = append(assignments, Assignment{
					Mask:    mask,
					Address: address,
					Value:   value,
				})
			}
			if err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			}
		}
	}
	{
		mem := map[uint64]uint64{}
		for _, a := range assignments {
			mem[a.Address] = a.MaskedValue()
		}
		var s uint64
		for _, a := range mem {
			s += a
		}
		fmt.Println("first:", s)
	}
}

func parseMask(s string) Mask {
	mask := Mask{}
	for i, b := range s {
		switch b {
		case '1':
			mask.Set = append(mask.Set, 35-i)
		case '0':
			mask.UnSet = append(mask.UnSet, 35-i)
		}
	}
	return mask
}

type Assignment struct {
	Mask    Mask
	Address uint64
	Value   uint64
}

func (a *Assignment) MaskedValue() uint64 {
	v := a.Value
	for _, bit := range a.Mask.Set {
		v |= 1 << bit
	}
	for _, bit := range a.Mask.UnSet {
		v &= ^(1 << bit)
	}
	return v
}

type Mask struct {
	Set   []int
	UnSet []int
}
