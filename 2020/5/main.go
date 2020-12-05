package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	entries := []string{}
	{
		bin := bufio.NewReader(os.Stdin)
		for {
			line, err := bin.ReadString('\n')
			if len(line) > 0 {
				line = strings.TrimRight(line, "\r\n")
				entries = append(entries, line)
			}
			if err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			}
		}
	}
	seatsTaken := make([]bool, 8*128)
	seatsTakenRange := Range{Low: len(seatsTaken), High: 0}
	first := 0
	for _, entry := range entries {
		bpi := decodeBoardingPass(entry)
		if bpi.SeatId > first {
			first = bpi.SeatId
		}
		if bpi.SeatId < seatsTakenRange.Low {
			seatsTakenRange.Low = bpi.SeatId
		}
		if bpi.SeatId > seatsTakenRange.High {
			seatsTakenRange.High = bpi.SeatId
		}
		seatsTaken[bpi.SeatId] = true
	}
	fmt.Println("first:", first)

	for i := seatsTakenRange.Low; i < seatsTakenRange.High; i++ {
		if !seatsTaken[i] {
			fmt.Println("second:", i)
			break
		}
	}
}

type BoardingPassInfo struct {
	Row    int
	Column int
	SeatId int
}

type Range struct {
	Low  int
	High int
}

func decodeBoardingPass(s string) BoardingPassInfo {
	row := Range{Low: 0, High: 127}
	for _, b := range s[:7] {
		d := ((row.High - row.Low) / 2) + 1
		switch b {
		case 'F':
			row.High -= d
		case 'B':
			row.Low += d
		}
	}
	col := Range{Low: 0, High: 7}
	for _, b := range s[7:] {
		d := ((col.High - col.Low) / 2) + 1
		switch b {
		case 'L':
			col.High -= d
		case 'R':
			col.Low += d
		}
	}
	return BoardingPassInfo{
		Row:    row.Low,
		Column: col.Low,
		SeatId: row.Low*8 + col.Low,
	}
}
