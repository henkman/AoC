package main

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type Unit uint32

type OpType uint8

type ValueType uint8

type Register uint8

const (
	OpType_Cpy OpType = iota
	OpType_Inc
	OpType_Dec
	OpType_Jnz
)

const (
	A Register = iota
	B
	C
	D
)

const (
	ValueType_Register ValueType = iota
	ValueType_Constant
)

type Value struct {
	Type     ValueType
	Register Register
	Constant Unit
}

func (v Value) Get(m Machine) Unit {
	if v.Type == ValueType_Constant {
		return v.Constant
	}
	return m.Register[v.Register]
}

type Op struct {
	Type   OpType
	Values []Value
}

type Machine struct {
	Register [4]Unit
}

func (m *Machine) Execute(prog []Op) {
	ip := 0
	for ip < len(prog) {
		op := prog[ip]
		switch op.Type {
		case OpType_Cpy:
			m.Register[op.Values[1].Register] = op.Values[0].Get(*m)
			ip++
		case OpType_Inc:
			m.Register[op.Values[0].Register]++
			ip++
		case OpType_Dec:
			m.Register[op.Values[0].Register]--
			ip++
		case OpType_Jnz:
			if m.Register[op.Values[0].Register] != 0 {
				ip += int(op.Values[1].Get(*m))
			} else {
				ip++
			}
		}
	}
}

func parseValue(s string) Value {
	if s[0] >= 'a' && s[0] <= 'd' {
		return Value{
			ValueType_Register,
			Register(s[0] - 'a'),
			0,
		}
	}
	v, _ := strconv.Atoi(s)
	return Value{
		ValueType_Constant,
		A,
		Unit(v),
	}
}

var (
	reOp = regexp.MustCompile("(cpy|inc|dec|jnz) (-?\\d+|[abcd])(?: (-?\\d+|[abcd]))?")
)

func parseProgram(in io.Reader) []Op {
	prog := make([]Op, 0, 16)
	bin := bufio.NewReader(in)
	for {
		line, _ := bin.ReadString('\n')
		if len(line) == 0 {
			break
		}
		line = strings.TrimRight(line, "\n")
		m := reOp.FindStringSubmatch(line)
		if m == nil {
			continue
		}
		var op Op
		switch m[1] {
		case "cpy":
			op.Type = OpType_Cpy
			op.Values = []Value{parseValue(m[2]), parseValue(m[3])}
		case "inc":
			op.Type = OpType_Inc
			op.Values = []Value{parseValue(m[2])}
		case "dec":
			op.Type = OpType_Dec
			op.Values = []Value{parseValue(m[2])}
		case "jnz":
			op.Type = OpType_Jnz
			op.Values = []Value{parseValue(m[2]), parseValue(m[3])}
		}
		prog = append(prog, op)
	}
	return prog
}
