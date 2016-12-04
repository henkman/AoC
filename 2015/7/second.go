package main

import (
	"fmt"
	"os"
)

func main() {
	wires := ParseWires(os.Stdin)
	v := wires["a"].Value(wires)
	b := wires["b"]
	b.Op = Op_Move
	b.Operands = []Operand{{Type: Operand_Signal, Signal: v}}
	for _, w := range wires {
		w.Reset()
	}
	fmt.Println(wires["a"].Value(wires))
}
