package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var nodes Nodes
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		dash := strings.IndexByte(line, '-')
		bn := line[:dash]
		en := line[dash+1:]
		bi := nodes.ByName(bn)
		if bi == -1 {
			bi = nodes.Create(bn)
		}
		ei := nodes.ByName(en)
		if ei == -1 {
			ei = nodes.Create(en)
		}
		if en != "start" && bn != "end" {
			nodes[bi].Paths = append(nodes[bi].Paths, ei)
		}
		if bn != "start" && en != "end" {
			nodes[ei].Paths = append(nodes[ei].Paths, bi)
		}
	}

	start := nodes.ByName("start")
	end := nodes.ByName("end")
	visited := make([]bool, len(nodes))
	count := DFS(nodes, start, end, visited)
	fmt.Println("first:", count)
}

func DFS(nodes Nodes, cur, end int, visited []bool) int {
	if cur == end {
		return 1
	}
	var count int
	for _, n := range nodes[cur].Paths {
		if !visited[n] {
			if nodes[n].Small {
				visited[n] = true
			}
			count += DFS(nodes, n, end, visited)
			visited[n] = false
		}
	}
	return count
}

type Nodes []Node

func (ns *Nodes) Create(name string) int {
	n := []byte(name)
	*ns = append(*ns, Node{
		Name:  name,
		Small: n[0] >= 'a' && n[0] <= 'z',
	})
	return len(*ns) - 1
}

func (ns Nodes) ByName(name string) int {
	for i, n := range ns {
		if n.Name == name {
			return i
		}
	}
	return -1
}

type Node struct {
	Name  string
	Small bool
	Paths []int
}
