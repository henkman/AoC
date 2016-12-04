package main

import "regexp"

var (
	reLine = regexp.MustCompile("([a-z-]+)-([0-9]+)\\[([a-z]+)\\]")
)

func isValidRoom(cipher, checksum []byte) bool {
	var ls ['z' + 1]int
	for _, c := range cipher {
		if c == '-' {
			continue
		}
		ls[c]++
	}
	var top [5]struct {
		Char  byte
		Count int
	}
	var c byte
	for c = 'a'; c <= 'z'; c++ {
		for i, _ := range top {
			t := &top[i]
			n := ls[c]
			if n > t.Count {
				for e := len(top) - 1; e > i; e-- {
					o := &top[e]
					o.Char = top[e-1].Char
					o.Count = top[e-1].Count
				}
				t.Char = c
				t.Count = n
				break
			}
		}
	}
	for i, c := range checksum {
		if c != top[i].Char {
			return false
		}
	}
	return true
}
