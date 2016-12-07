package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	abc   = []byte("abcdefghijklmnopqrstuvwxyz")
	rule1 = func() [][]byte {
		const LEN = 3
		arr := make([][]byte, 0, 8)
		for i := 0; i <= len(abc)-LEN; i++ {
			arr = append(arr, abc[i:i+LEN])
		}
		return arr
	}()
	rule2 = func() []byte {
		return []byte("iol")
	}()
	rule3 = func() [][]byte {
		arr := make([][]byte, 0, 8)
		for _, c := range abc {
			arr = append(arr, []byte{c, c})
		}
		return arr
	}()
)

func isValid(pass []byte) bool {
	for _, c := range pass {
		for _, r := range rule2 {
			if c == r {
				return false
			}
		}
	}
	{
		has := false
		for _, r := range rule1 {
			if bytes.Contains(pass, r) {
				has = true
				break
			}
		}
		if !has {
			return false
		}
	}
	{
		pairs := 0
		for _, r := range rule3 {
			if bytes.Contains(pass, r) {
				pairs++
			}
		}
		if pairs < 2 {
			return false
		}
	}
	return true
}

func main() {
	var pass []byte
	_, err := fmt.Fscanf(os.Stdin, "%s\n", &pass)
	if err != nil {
		if err == io.EOF {
			return
		}
		log.Fatal(err)
	}
	first := true
	for {
		if isValid(pass) {
			if first {
				fmt.Println("first:", string(pass))
				first = false
			} else {
				fmt.Println("second:", string(pass))
				break
			}
		}
		pass[len(pass)-1]++
		for i := len(pass) - 2; i >= 0; i-- {
			if pass[i+1] > 'z' {
				pass[i+1] = 'a'
				pass[i]++
			}
		}
	}
}
