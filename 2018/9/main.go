package main

import (
	"fmt"
	"os"
)

type Marble struct {
	Value int
	Next  *Marble
	Prev  *Marble
}

func main() {
	var input struct {
		Players    int
		LastMarble int
	}
	{
		fd, err := os.Open("input.txt")
		if err != nil {
			panic(err)
		}
		defer fd.Close()
		fmt.Fscanf(fd, "%d players; last marble is worth %d points",
			&input.Players, &input.LastMarble)
	}
	players := make([]int, input.Players)
	cp := 0
	cm := &Marble{Value: 0}
	cm.Next = cm
	cm.Prev = cm
	for mv := 1; mv < (input.LastMarble * 100); mv++ {
		if mv == input.LastMarble {
			l := 0
			for i := range players {
				if players[i] > players[l] {
					l = i
				}
			}
			fmt.Println("first:", players[l])
		}
		if mv%23 == 0 {
			players[cp] += mv
			rm := cm
			for i := 0; i < 7; i++ {
				rm = rm.Prev
			}
			players[cp] += rm.Value
			prev := rm.Prev
			next := rm.Next
			prev.Next = next
			next.Prev = prev
			rm.Next = nil
			rm.Prev = nil
			cm = next
		} else {
			prev := cm.Next
			next := prev.Next
			cm = &Marble{
				Value: mv,
				Prev:  prev,
				Next:  next,
			}
			prev.Next = cm
			next.Prev = cm
		}
		cp++
		if cp == input.Players {
			cp = 0
		}
	}
	l := 0
	for i := range players {
		if players[i] > players[l] {
			l = i
		}
	}
	fmt.Println("second:", players[l])
}
