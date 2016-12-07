package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	sum := 0
	dec := json.NewDecoder(os.Stdin)
	for {
		t, err := dec.Token()
		if err == io.EOF {
			break
		}
		switch ot := t.(type) {
		case json.Number:
			f, err := ot.Float64()
			if err != nil {
				log.Fatal(err)
			}
			sum += int(f)
		case float64:
			sum += int(ot)
		}
	}
	fmt.Println(sum)
}
