package main

import (
	"fmt"
	"os"
)

func main() {
	wires := ParseWires(os.Stdin)
	fmt.Println(wires["a"].Value(wires))
}
