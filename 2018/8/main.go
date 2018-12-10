package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	Children []Node
	Metadata []int32
}

func readNode(bin *bufio.Reader) (Node, error) {
	var header struct {
		Children int32
		Metadata int32
	}
	n, err := fmt.Fscanf(bin, "%d %d ", &header.Children, &header.Metadata)
	if n == 0 {
		return Node{}, err
	}
	var node Node
	for i := int32(0); i < header.Children; i++ {
		c, err := readNode(bin)
		if err == nil {
			node.Children = append(node.Children, c)
		}
	}
	for i := int32(0); i < header.Metadata; i++ {
		var v int32
		fmt.Fscanf(bin, "%d ", &v)
		node.Metadata = append(node.Metadata, v)
	}
	return node, nil
}

func firstSum(node Node) int64 {
	var sum int64
	for _, md := range node.Metadata {
		sum += int64(md)
	}
	for _, c := range node.Children {
		sum += firstSum(c)
	}
	return sum
}

func secondSum(node Node) int64 {
	var sum int64
	if len(node.Children) == 0 {
		for _, md := range node.Metadata {
			sum += int64(md)
		}
	} else {
		for _, md := range node.Metadata {
			idx := md - 1
			if idx >= 0 && idx < int32(len(node.Children)) {
				sum += secondSum(node.Children[idx])
			}
		}
	}
	return sum
}

func main() {
	var root Node
	{
		fd, err := os.Open("input.txt")
		if err != nil {
			panic(err)
		}
		defer fd.Close()
		bin := bufio.NewReader(fd)
		node, err := readNode(bin)
		if err != nil {
			panic(err)
		}
		root = node
	}

	{ // first
		fmt.Println(firstSum(root))
	}

	{ // second
		fmt.Println(secondSum(root))
	}
}
