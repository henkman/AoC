package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
)

type Node struct {
	Name      string
	Weight    int
	SumWeight int
	Node      []*Node
}

func findNodeByName(name string, nodes []Node) *Node {
	for i, _ := range nodes {
		if name == nodes[i].Name {
			return &nodes[i]
		}
	}
	return nil
}

func second(root *Node) int {
	return 0
}

func main() {
	var root *Node
	{
		bin := bufio.NewReader(os.Stdin)
		reNode := regexp.MustCompile(
			`^([a-z]+ \(\d+)\)(?: -> ((?:[a-z]+, )+[a-z]+))?$`)
		nodes := make([]Node, 0, 16)
		pars := map[string]string{}
		for {
			line, err := bin.ReadBytes('\n')
			if len(line) > 0 {
				m := reNode.FindAllSubmatch(bytes.TrimSpace(line),
					-1)
				if m != nil {
					var node Node
					fmt.Sscanf(string(m[0][1]), "%s (%d)",
						&node.Name, &node.Weight)
					nodes = append(nodes, node)
					if len(m[0][2]) != 0 {
						for _, n := range bytes.Split(m[0][2], []byte(",")) {
							pars[string(bytes.TrimSpace(n))] = node.Name
						}
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
		for i, node := range nodes {
			if _, ok := pars[node.Name]; !ok {
				root = &nodes[i]
			}
		}
		for child, par := range pars {
			parent := findNodeByName(par, nodes)
			child := findNodeByName(child, nodes)
			parent.Node = append(parent.Node, child)
		}
	}
	fmt.Println("first:", root.Name)
	fmt.Println(second(root))
}
