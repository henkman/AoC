package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var key string
	_, err := fmt.Fscanf(os.Stdin, "%s\n", &key)
	if err != nil {
		if err == io.EOF {
			return
		}
		log.Fatal(err)
	}
	b := []byte(key)
	for i := 0; i < 40; i++ {
		b = lookAndSay(b)
	}
	fmt.Println(len(b))
}
