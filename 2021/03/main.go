package main

import (
	"fmt"
	"strconv"
)

func main() {
	entries := []int{}
	bits := 0
	{
		var s string
		for {
			_, err := fmt.Scan(&s)
			if err != nil {
				bits = len(s)
				break
			}
			v, _ := strconv.ParseInt(s, 2, 32)
			entries = append(entries, int(v))
		}
	}

	{
		ones := make([]int, bits)
		for _, entry := range entries {
			for i := 0; i < bits; i++ {
				v := entry & (1 << i)
				if v != 0 {
					ones[bits-i-1]++
				}
			}
		}
		h := len(entries) / 2
		var gamma, epsilon int
		for i := 0; i < bits; i++ {
			gamma <<= 1
			epsilon <<= 1
			if ones[i] >= h {
				gamma |= 1
			} else {
				epsilon |= 1
			}
		}
		fmt.Println("first:", gamma*epsilon)
	}

	{
		var oxygen, co2scrubber int
		filtered := make([]int, len(entries))
		{
			copy(filtered, entries)
			n := len(entries)
			for i := bits - 1; i >= 0; i-- {
				keep := filter(filtered[:n], i, true)
				if len(keep) <= 1 {
					oxygen = keep[0]
					break
				}
				copy(filtered, keep)
				n = len(keep)
			}
		}
		{
			copy(filtered, entries)
			n := len(entries)
			for i := bits - 1; i >= 0; i-- {
				keep := filter(filtered[:n], i, false)
				if len(keep) <= 1 {
					co2scrubber = keep[0]
					break
				}
				copy(filtered, keep)
				n = len(keep)
			}
		}
		fmt.Println("second:", oxygen*co2scrubber)
	}
}

func filter(entries []int, bit int, filterOnes bool) []int {
	mask := 1 << bit
	var ones int
	for _, entry := range entries {
		if entry&mask != 0 {
			ones++
		}
	}
	zeros := len(entries) - ones

	var keepOnes bool
	if filterOnes {
		keepOnes = ones >= zeros
	} else {
		keepOnes = ones < zeros
	}

	keep := []int{}
	for _, entry := range entries {
		one := entry&mask != 0
		if one == keepOnes {
			keep = append(keep, entry)
		}
	}
	return keep
}
