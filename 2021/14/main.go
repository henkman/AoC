package main

import (
	"bufio"
	"container/list"
	"fmt"
	"math"
	"os"
)

func main() {
	rules := map[string]byte{}
	poly := list.New()
	{
		scanner := bufio.NewScanner(os.Stdin)
		var templ string
		if scanner.Scan() {
			templ = scanner.Text()
		}
		if !scanner.Scan() {
			panic("empty line missing")
		}

		for scanner.Scan() {
			var match string
			var insert byte
			line := scanner.Text()
			if _, err := fmt.Sscanf(line, "%s -> %c",
				&match, &insert); err != nil {
				break
			}
			rules[match] = insert
		}
		for _, c := range []byte(templ) {
			poly.PushBack(c)
		}
	}

	for i := 0; i < 10; i++ {
		cur := poly.Front()
		for cur.Next() != nil {
			f := cur.Value.(byte)
			s := cur.Next().Value.(byte)
			b := []byte{f, s}
			ins := rules[string(b)]
			n := poly.InsertAfter(ins, cur)
			cur = n.Next()
		}
	}

	cm := map[byte]int{}
	iter := poly.Front()
	for iter != nil {
		cm[iter.Value.(byte)]++
		iter = iter.Next()
	}

	mx := 0
	mn := math.MaxInt64
	for _, c := range cm {
		if c > mx {
			mx = c
		} else if c < mn {
			mn = c
		}
	}

	fmt.Println("first:", mx-mn)
}
