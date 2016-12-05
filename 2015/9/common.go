package main

import (
	"bufio"
	"io"
	"regexp"
	"sort"
	"strconv"
)

type Route struct {
	Node     *Node
	Distance int
}

type Node struct {
	Name   string
	Routes []Route
}

var (
	reRoute = regexp.MustCompile("([^ ]+) to ([^ ]+) = (\\d+)")
)

func (n *Node) Cost(to *Node) int {
	for _, r := range n.Routes {
		if r.Node.Name == to.Name {
			return r.Distance
		}
	}
	panic("unreachable")
}

func ParseNodes(in io.Reader) ([]*Node, error) {
	bin := bufio.NewReader(in)
	nodes := make([]*Node, 0, 8)
	noderef := map[string]*Node{}
	for {
		line, err := bin.ReadString('\n')
		if line == "" {
			break
		}
		m := reRoute.FindStringSubmatch(line)
		if m == nil {
			continue
		}
		d, err := strconv.Atoi(m[3])
		if err != nil {
			return nil, err
		}
		a, hasA := noderef[m[1]]
		if !hasA {
			a = &Node{m[1], []Route{}}
			noderef[m[1]] = a
			nodes = append(nodes, a)
		}
		b, hasB := noderef[m[2]]
		if !hasB {
			b = &Node{m[2], []Route{}}
			noderef[m[2]] = b
			nodes = append(nodes, b)
		}
		a.Routes = append(a.Routes, Route{b, d})
		b.Routes = append(b.Routes, Route{a, d})
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return nodes, nil
}

func NextPermutation(data sort.Interface) bool {
	var k, l int
	for k = data.Len() - 2; ; k-- {
		if k < 0 {
			return false
		}
		if data.Less(k, k+1) {
			break
		}
	}
	for l = data.Len() - 1; !data.Less(k, l); l-- {
	}
	data.Swap(k, l)
	for i, j := k+1, data.Len()-1; i < j; i++ {
		data.Swap(i, j)
		j--
	}
	return true
}
