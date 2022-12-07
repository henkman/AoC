package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

type Stack []byte

type Instruction struct {
	Amount int
	From   int
	To     int
}

var (
	reInitialStacks = regexp.MustCompile(`(?:\[[A-Z]\]|\s{3}) ?`)
)

func reverse(input []byte) {
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}
}

func main() {
	const DEBUG = false

	var stacks []Stack
	var instructions []Instruction
	{
		bin := bufio.NewReader(os.Stdin)
		for {
			line, err := bin.ReadString('\n')
			if len(line) > 0 {
				line = strings.TrimRight(line, "\r\n")
				if len(line) == 0 {
					break
				}
				fm := reInitialStacks.FindAllString(line, -1)
				if cap(stacks) == 0 {
					stacks = make([]Stack, len(fm))
				}
				for i, m := range fm {
					x := strings.TrimSpace(m)
					if len(x) != 0 {
						stacks[i] = append(stacks[i], m[1])
					}
				}
			}
			if err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			}
		}
		for {
			line, err := bin.ReadString('\n')
			if len(line) > 0 {
				line = strings.TrimRight(line, "\r\n")
				var instr Instruction
				fmt.Sscanf(line, "move %d from %d to %d",
					&instr.Amount, &instr.From, &instr.To)
				instructions = append(instructions, instr)
			}
			if err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			}
		}
		for _, stack := range stacks {
			reverse(stack)
			if DEBUG {
				fmt.Println(string(stack))
			}
		}
	}

	{ // first
		wss := make([]Stack, len(stacks))
		copy(wss, stacks)
		for _, instr := range instructions {
			from := &wss[instr.From-1]
			to := &wss[instr.To-1]
			take := (*from)[len(*from)-instr.Amount:]
			if DEBUG {
				fmt.Printf("take %s from %s(%d) to %s(%d)\n",
					string(take), string(*from), instr.From, string(*to), instr.To)
			}
			for i := 0; i < instr.Amount; i++ {
				c := (*from)[len(*from)-1-i]
				*to = append(*to, c)
			}
			*from = (*from)[:len(*from)-instr.Amount]
			if DEBUG {
				fmt.Println("---")
				for _, stack := range wss {
					fmt.Println(string(stack))
				}
				fmt.Println("---")
			}
		}
		var first []byte
		for _, stack := range wss {
			l := stack[len(stack)-1]
			first = append(first, l)
		}
		fmt.Println("first:", string(first))
	}
	{ // second
		wss := make([]Stack, len(stacks))
		copy(wss, stacks)

		for _, instr := range instructions {
			from := &wss[instr.From-1]
			to := &wss[instr.To-1]
			take := (*from)[len(*from)-instr.Amount:]
			if DEBUG {
				fmt.Printf("take %s from %s(%d) to %s(%d)\n",
					string(take), string(*from), instr.From, string(*to), instr.To)
			}
			*to = append(*to, take...)
			*from = (*from)[:len(*from)-instr.Amount]
			if DEBUG {
				fmt.Println("---")
				for _, stack := range wss {
					fmt.Println(string(stack))
				}
				fmt.Println("---")
			}
		}
		var first []byte
		for _, stack := range wss {
			l := stack[len(stack)-1]
			first = append(first, l)
		}
		fmt.Println("second:", string(first))
	}
}
