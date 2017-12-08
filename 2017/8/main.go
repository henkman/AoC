package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

type Op uint8

const (
	OpInc Op = iota
	OpDec
)

type LogOp uint8

const (
	LogOpSmaller LogOp = iota
	LogOpSmallerOrEqual
	LogOpGreater
	LogOpGreaterOrEqual
	LogOpEqual
	LogOpNotEqual
)

type Instruction struct {
	Name   string
	Op     Op
	Amount int
	Cond   struct {
		Name   string
		LogOp  LogOp
		Amount int
	}
}

func main() {
	instrs := make([]Instruction, 0, 16)
	{
		reInst := regexp.MustCompile(
			`([a-z]+) (dec|inc) (-?\d+) if ([a-z]+) ([=!><]+) (-?\d+)`)
		bin := bufio.NewReader(os.Stdin)
		for {
			line, err := bin.ReadBytes('\n')
			if len(line) > 0 {
				m := reInst.FindAllSubmatch(bytes.TrimSpace(line), -1)
				var instr Instruction
				instr.Name = string(m[0][1])
				instr.Op = map[string]Op{
					"dec": OpDec,
					"inc": OpInc,
				}[string(m[0][2])]
				instr.Amount, _ = strconv.Atoi(string(m[0][3]))
				instr.Cond.Name = string(m[0][4])
				instr.Cond.LogOp = map[string]LogOp{
					">":  LogOpGreater,
					"<":  LogOpSmaller,
					"==": LogOpEqual,
					"!=": LogOpNotEqual,
					">=": LogOpGreaterOrEqual,
					"<=": LogOpSmallerOrEqual,
				}[string(m[0][5])]
				instr.Cond.Amount, _ = strconv.Atoi(string(m[0][6]))
				instrs = append(instrs, instr)
			}
			if err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			}
		}
	}
	regs := map[string]int{}
	{
		max := 0
		for _, instr := range instrs {
			creg, ok := regs[instr.Cond.Name]
			if !ok {
				creg = 0
			}
			if instr.Cond.LogOp == LogOpEqual &&
				creg != instr.Cond.Amount {
				continue
			} else if instr.Cond.LogOp == LogOpNotEqual &&
				creg == instr.Cond.Amount {
				continue
			} else if instr.Cond.LogOp == LogOpGreater &&
				creg <= instr.Cond.Amount {
				continue
			} else if instr.Cond.LogOp == LogOpGreaterOrEqual &&
				creg < instr.Cond.Amount {
				continue
			} else if instr.Cond.LogOp == LogOpSmaller &&
				creg >= instr.Cond.Amount {
				continue
			} else if instr.Cond.LogOp == LogOpSmallerOrEqual &&
				creg > instr.Cond.Amount {
				continue
			}
			sreg, ok := regs[instr.Name]
			if !ok {
				sreg = 0
			}
			if instr.Op == OpInc {
				sreg += instr.Amount
			} else {
				sreg -= instr.Amount
			}
			if sreg > max {
				max = sreg
			}
			regs[instr.Name] = sreg
		}
		fmt.Println("second:", max)
	}
	{
		max := 0
		for _, reg := range regs {
			if reg > max {
				max = reg
			}
		}
		fmt.Println("first:", max)
	}
}
