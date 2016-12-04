package main

import (
	"bufio"
	"io"
	"log"
	"regexp"
	"strconv"
)

type Op uint8

const (
	Op_Undefined Op = iota
	Op_Move
	Op_Not
	Op_And
	Op_Or
	Op_Lshift
	Op_Rshift
)

type OperandType uint8

const (
	Operand_Identifier OperandType = iota
	Operand_Signal
	Operand_SignaledIdentifier
)

type Operand struct {
	Type       OperandType
	Identifier string
	Signal     uint16
}

type Wire struct {
	Op       Op
	Operands []Operand
}

func (o *Operand) Value(wires map[string]*Wire) uint16 {
	if o.Type == Operand_Identifier {
		if w, ok := wires[o.Identifier]; ok {
			v := w.Value(wires)
			o.Type = Operand_SignaledIdentifier
			o.Signal = v
			return v
		} else {
			return 0
		}
	}
	return o.Signal
}

func (o *Operand) Reset() {
	if o.Type == Operand_SignaledIdentifier {
		o.Type = Operand_Identifier
		o.Signal = 0
	}
}

func (w *Wire) Value(wires map[string]*Wire) uint16 {
	switch w.Op {
	case Op_Undefined:
		return 0
	case Op_Move:
		return w.Operands[0].Value(wires)
	case Op_Not:
		v := w.Operands[0].Value(wires)
		return ^v
	case Op_And:
		return w.Operands[0].Value(wires) & w.Operands[1].Value(wires)
	case Op_Or:
		return w.Operands[0].Value(wires) | w.Operands[1].Value(wires)
	case Op_Lshift:
		return w.Operands[0].Value(wires) << w.Operands[1].Value(wires)
	case Op_Rshift:
		return w.Operands[0].Value(wires) >> w.Operands[1].Value(wires)
	}
	panic("unreachable")
}

func (w *Wire) Reset() {
	for i, _ := range w.Operands {
		w.Operands[i].Reset()
	}
}

func ParseOperand(s string) Operand {
	if s[0] >= 'a' && s[0] <= 'z' {
		return Operand{Type: Operand_Identifier, Identifier: s}
	}
	v, err := strconv.ParseUint(s, 10, 16)
	if err != nil {
		log.Fatal(err)
	}
	return Operand{Type: Operand_Signal, Signal: uint16(v)}
}

func ParseWires(in io.Reader) map[string]*Wire {
	bin := bufio.NewReader(in)
	wires := map[string]*Wire{}
	reInst := regexp.MustCompile("(?:([a-z]+|\\d+) )?(?:(AND|OR|LSHIFT|RSHIFT|NOT) )?([a-z]+|\\d+) -> ([a-z]+)")
	for {
		line, err := bin.ReadString('\n')
		if line == "" {
			break
		}
		m := reInst.FindStringSubmatch(line)
		if m == nil {
			continue
		}
		w := new(Wire)
		switch m[2] {
		case "AND":
			w.Op = Op_And
			w.Operands = []Operand{ParseOperand(m[1]), ParseOperand(m[3])}
		case "OR":
			w.Op = Op_Or
			w.Operands = []Operand{ParseOperand(m[1]), ParseOperand(m[3])}
		case "LSHIFT":
			w.Op = Op_Lshift
			w.Operands = []Operand{ParseOperand(m[1]), ParseOperand(m[3])}
		case "RSHIFT":
			w.Op = Op_Rshift
			w.Operands = []Operand{ParseOperand(m[1]), ParseOperand(m[3])}
		case "NOT":
			w.Op = Op_Not
			w.Operands = []Operand{ParseOperand(m[3])}
		default:
			w.Op = Op_Move
			w.Operands = []Operand{ParseOperand(m[3])}
		}
		wires[m[4]] = w
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
	}
	return wires
}
