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
	var w Point
	dir := Dir_North
	for _, m := range reInstruction.FindAllStringSubmatch(string(raw), -1) {
		dir = dir.Turn(byte(m[1][0]) == 'L')
		n, err := strconv.Atoi(m[2])
		if err != nil {
			log.Fatal(err)
		}
		w = w.WalkStraight(dir, n)
	}
	fmt.Println(w.Distance(Point{0, 0}))
}
