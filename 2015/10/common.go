package main

import "fmt"

func lookAndSay(s []byte) []byte {
	r := make([]byte, 0, len(s))
	n := 1
	l := s[0]
	for i := 1; i < len(s); i++ {
		if l != s[i] {
			r = append(r, []byte(fmt.Sprint(n))...)
			r = append(r, l)
			l = s[i]
			n = 1
		} else {
			n++
		}
	}
	if n > 0 {
		r = append(r, []byte(fmt.Sprint(n))...)
		r = append(r, l)
		l = s[len(s)-1]
	}
	return r
}
