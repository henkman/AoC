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
	bags := Bags{}
	{
		input, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
		reContainer := regexp.MustCompile(`^([a-z ]+?) bags contain `)
		bin := bufio.NewReader(bytes.NewBuffer(input))
		for {
			line, err := bin.ReadString('\n')
			if len(line) > 0 {
				line = strings.TrimRight(line, "\r\n")
				cr := reContainer.FindStringSubmatch(line)
				bags = append(bags, Bag{
					Name: cr[1],
				})
			}
			if err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			}
		}
		reContained := regexp.MustCompile(`(\d+) ([a-z ]+?) bags?`)
		bin.Reset(bytes.NewBuffer(input))
		for {
			line, err := bin.ReadString('\n')
			if len(line) > 0 {
				line = strings.TrimRight(line, "\r\n")
				cr := reContainer.FindStringSubmatch(line)
				parent := bags.FindByName(cr[1])
				cd := reContained.FindAllStringSubmatch(line, -1)
				for _, m := range cd {
					n, err := strconv.Atoi(m[1])
					if err != nil {
						panic(err)
					}
					child := bags.FindByName(m[2])
					child.Parents = append(parent.Parents, BagRelation{
						Bag:      parent,
						Quantity: n,
					})
					parent.Children = append(parent.Children, BagRelation{
						Bag:      child,
						Quantity: n,
					})
				}
			}
			if err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			}
		}
	}

	const MYBAG = `shiny gold`
	first := 0
	for _, bag := range bags {
		if CanContainRecursive(&bag, MYBAG) {
			first++
		}
	}
	fmt.Println(first)

	shinygold := bags.FindByName(MYBAG)
	fmt.Println("second:", CountBagsRecursively(shinygold))
}

func CountBagsRecursively(bag *Bag) uint64 {
	var count uint64
	for _, child := range bag.Children {
		count += uint64(child.Quantity)
		count += uint64(child.Quantity) * CountBagsRecursively(child.Bag)
	}
	return count
}

func CanContainRecursive(bag *Bag, search string) bool {
	for _, child := range bag.Children {
		if child.Bag.Name == search {
			return true
		}
		if CanContainRecursive(child.Bag, search) {
			return true
		}
	}
	return false
}

type Bag struct {
	Name     string
	Parents  []BagRelation
	Children []BagRelation
}

type BagRelation struct {
	Bag      *Bag
	Quantity int
}

type Bags []Bag

func (bs *Bags) FindByName(name string) *Bag {
	for i, bag := range *bs {
		if bag.Name == name {
			return &(*bs)[i]
		}
	}
	return nil
}
