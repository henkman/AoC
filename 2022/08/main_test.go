package main

import (
	"fmt"
	"testing"
)

var (
	tm = TreeMap{
		Trees: []Tree{
			3, 0, 3, 7, 3,
			2, 5, 5, 1, 2,
			6, 5, 3, 3, 2,
			3, 3, 5, 4, 9,
			3, 5, 3, 9, 0,
		},
		Width:  5,
		Height: 5,
	}
)

func TestVisibilityMap(t *testing.T) {
	return
	tm.CalculateVisibilityMap()
	vs := []Visibility{Visiblity_Left, Visiblity_Right,
		Visiblity_Top, Visiblity_Bottom}
	vsm := map[Visibility]byte{
		Visiblity_Left: 'L', Visiblity_Right: 'R',
		Visiblity_Top: 'T', Visiblity_Bottom: 'B'}
	for y := 0; y < tm.Height; y++ {
		for x := 0; x < tm.Width; x++ {
			fmt.Print("(")
			for _, v := range vs {
				if tm.IsVisibleFrom(x, y, v) {
					fmt.Printf("%c", vsm[v])
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Print(") ")
		}
		fmt.Println()
	}
	sum := 0
	for _, t := range tm.VisibilityMap {
		if t != 0 {
			sum++
		}
	}
	fmt.Println(sum)
}

func TestScenicScoreMap(t *testing.T) {
	tm.CalculateScenicScoreMap()
	for y := 0; y < tm.Height; y++ {
		for x := 0; x < tm.Width; x++ {
			fmt.Printf("%03d ", tm.ScenicScore(x, y))
		}
		fmt.Println()
	}
	var highest uint
	for _, t := range tm.ScenicScoreMap {
		if t > highest {
			highest = t
		}
	}
	fmt.Println(highest)
}
