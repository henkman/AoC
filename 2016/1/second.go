package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {
	raw, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	past := map[Point]interface{}{}
	var w Point
	dir := Dir_North
	for _, m := range reInstruction.FindAllStringSubmatch(string(raw), -1) {
		dir = dir.Turn(byte(m[1][0]) == 'L')
		n, err := strconv.Atoi(m[2])
		if err != nil {
			log.Fatal(err)
		}
		path := w.WalkStraightPath(dir, n)
		for _, p := range path {
			if _, ok := past[p]; ok {
				fmt.Println(p, p.Distance(Point{0, 0}))
				return
			} else {
				past[p] = true
			}
		}
		w = path[len(path)-1]
	}
}
