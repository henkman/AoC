package main

import (
	"fmt"
	"os"
)

func main() {
	program := parseProgram(os.Stdin)
	var m Machine
	m.Register[C] = 1
	m.Execute(program)
	fmt.Println(m.Register[A])
}
