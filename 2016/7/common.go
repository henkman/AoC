package main

import "regexp"

var (
	reInside  = regexp.MustCompile("\\[([^\\]]+)\\]")
	reOutside = regexp.MustCompile("(?:^|\\])([^\\[]+)(?:$|\\[)")
)

func isPalindrome(arr []byte) bool {
	l := len(arr)
	h := l / 2
	e := l - 1
	for i := 0; i < h; i++ {
		if arr[i] != arr[e-i] {
			return false
		}
	}
	return true
}
