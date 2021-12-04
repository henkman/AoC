package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const BOARDSIZE = 5

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	line := scanner.Text()
	reNumber := regexp.MustCompile(`[0-9]+`)
	m := reNumber.FindAllString(line, -1)
	draws := make([]int, 0, len(m))
	for _, e := range m {
		v, _ := strconv.Atoi(e)
		draws = append(draws, v)
	}

	boards := []Board{}
	for scanner.Scan() {
		board := Board{
			Numbers: make([]int, BOARDSIZE*BOARDSIZE),
			Marked:  make([]bool, BOARDSIZE*BOARDSIZE),
		}
		o := 0
		for i := 0; i < BOARDSIZE; i++ {
			if !scanner.Scan() {
				panic("strange board")
			}
			line := scanner.Text()
			m := reNumber.FindAllString(line, -1)
			for _, e := range m {
				v, _ := strconv.Atoi(e)
				board.Numbers[o] = v
				o++
			}
		}
		boards = append(boards, board)
	}

	{
	winner:
		for _, draw := range draws {
			for _, board := range boards {
				board.MarkDraw(draw)
				if board.HasBingo() {
					s := board.SumUnmarked()
					fmt.Println("first:", s*draw)
					break winner
				}
			}
		}
	}

	{
		boardsInGame := map[int]Board{}
		for i, board := range boards {
			boardsInGame[i] = board
		}
	loser:
		for _, draw := range draws {
			for i, board := range boardsInGame {
				board.MarkDraw(draw)
				if board.HasBingo() {
					if len(boardsInGame) > 1 {
						delete(boardsInGame, i)
						continue
					}
					s := board.SumUnmarked()
					fmt.Println("second:", s*draw)
					break loser
				}
			}
		}
	}
}

type Board struct {
	Numbers []int
	Marked  []bool
}

func (b *Board) MarkDraw(draw int) {
	for i, n := range b.Numbers {
		if draw == n {
			b.Marked[i] = true
			break
		}
	}
}

func (b *Board) HasBingo() bool {
horizontal:
	for y := 0; y < BOARDSIZE; y++ {
		for x := 0; x < BOARDSIZE; x++ {
			o := y*BOARDSIZE + x
			if !b.Marked[o] {
				continue horizontal
			}
		}
		return true
	}
vertical:
	for x := 0; x < BOARDSIZE; x++ {
		for y := 0; y < BOARDSIZE; y++ {
			o := y*BOARDSIZE + x
			if !b.Marked[o] {
				continue vertical
			}
		}
		return true
	}
	return false
}

func (b *Board) SumUnmarked() int {
	s := 0
	for i, v := range b.Marked {
		if !v {
			s += b.Numbers[i]
		}
	}
	return s
}
