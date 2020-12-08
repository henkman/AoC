package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	program := []Instruction{}
	{
		input, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
		reInstruction := regexp.MustCompile(`^(acc|jmp|nop) ([+-])(\d+)$`)
		bin := bufio.NewReader(bytes.NewBuffer(input))
		for {
			line, err := bin.ReadString('\n')
			if len(line) > 0 {
				line = strings.TrimRight(line, "\r\n")
				m := reInstruction.FindStringSubmatch(line)
				n, err := strconv.Atoi(m[3])
				if err != nil {
					panic(err)
				}
				if m[2] == "-" {
					n = -n
				}
				program = append(program, Instruction{
					Operation: Operation(m[1]),
					Argument:  n,
				})
			}
			if err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			}
		}
	}

	var c Console
	c.Execute(program)
	fmt.Println("first:", c.Acc)

	pc := make([]Instruction, len(program))
	o := 0
	for {
		c = Console{}
		copy(pc, program)
		var pins *Instruction
		e := 0
		for i, ins := range pc {
			if ins.Operation == OperationNop || ins.Operation == OperationJmp {
				if e == o {
					pins = &pc[i]
					break
				} else {
					e++
				}
			}
		}
		switch pins.Operation {
		case OperationJmp:
			pins.Operation = OperationNop
		case OperationNop:
			pins.Operation = OperationJmp
		}
		c.Execute(pc)
		if c.IP >= len(pc) {
			fmt.Println("second:", c.Acc)
			break
		}
		o++
	}
}

type Console struct {
	Acc int
	IP  int
}

func (c *Console) Execute(program []Instruction) {
	used := map[int]bool{}
	for c.IP < len(program) {
		if _, ok := used[c.IP]; ok {
			return
		}
		ins := program[c.IP]
		used[c.IP] = true
		switch ins.Operation {
		case OperationAcc:
			c.Acc += ins.Argument
			c.IP++
		case OperationJmp:
			c.IP += ins.Argument
		case OperationNop:
			c.IP++
		}
	}
}

type Instruction struct {
	Operation Operation
	Argument  int
}

type Operation string

const (
	OperationAcc Operation = "acc"
	OperationJmp Operation = "jmp"
	OperationNop Operation = "nop"
)
