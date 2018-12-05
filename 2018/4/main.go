package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"
)

type EventType uint8

const (
	EventType_ShiftBegin EventType = iota
	EventType_FallsAsleep
	EventType_WakesUp
)

type Event struct {
	Time       time.Time
	Type       EventType
	ShiftBegin struct {
		Guard uint32
	}
}

type Range struct {
	Start time.Time
	End   time.Time
}

type Events []Event

func (a Events) Len() int           { return len(a) }
func (a Events) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Events) Less(i, j int) bool { return a[i].Time.Before(a[j].Time) }

type MostMin struct {
	Minute byte
	Max    uint32
}

type GuardMostMin struct {
	Guard uint32
	MostMin
}

type GuardMostMins []GuardMostMin

func (a GuardMostMins) Len() int           { return len(a) }
func (a GuardMostMins) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a GuardMostMins) Less(i, j int) bool { return a[i].Max > a[j].Max }

func readEvents() ([]Event, error) {
	fd, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	reEvent := regexp.MustCompile(`\[([^\]]+)\] (.*)`)
	input := make([]Event, 0, 64)
	bin := bufio.NewReader(fd)
	for {
		line, err := bin.ReadString('\n')
		if len(line) > 0 {
			var ev Event
			m := reEvent.FindStringSubmatch(strings.TrimSpace(line))
			t, err := time.Parse("2006-01-02 15:04", m[1])
			if err != nil {
				panic(err)
			}
			ev.Time = t
			if m[2] == "wakes up" {
				ev.Type = EventType_WakesUp
			} else if m[2] == "falls asleep" {
				ev.Type = EventType_FallsAsleep
			} else {
				ev.Type = EventType_ShiftBegin
				fmt.Sscanf(m[2], "Guard #%d begins shift", &ev.ShiftBegin.Guard)
			}
			input = append(input, ev)
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return input, nil
}

func main() {
	events, err := readEvents()
	if err != nil {
		panic(err)
	}
	sort.Sort(Events(events))

	table := map[uint32][]Range{}
	var guard uint32
	var rng Range
	for _, ev := range events {
		if ev.Type == EventType_ShiftBegin {
			guard = ev.ShiftBegin.Guard
		} else if ev.Type == EventType_FallsAsleep {
			rng.Start = ev.Time
		} else if ev.Type == EventType_WakesUp {
			rng.End = ev.Time
			if times, ok := table[guard]; ok {
				table[guard] = append(times, rng)
			} else {
				table[guard] = []Range{rng}
			}
		}
	}

	{ // first
		var most struct {
			Guard uint32
			Sum   uint32
		}
		for guard, times := range table {
			var sum uint32
			for _, time := range times {
				sum += uint32(time.End.Sub(time.Start).Minutes())
			}
			if sum > most.Sum {
				most.Guard = guard
				most.Sum = sum
			}
		}
		mins := map[byte]uint32{}
		for _, rng := range table[most.Guard] {
			for cur := rng.Start; cur.Before(rng.End); cur = cur.Add(time.Minute) {
				min := byte(cur.Minute())
				if cnt, ok := mins[min]; ok {
					mins[min] = cnt + 1
				} else {
					mins[min] = 1
				}
			}
		}
		var mostMin MostMin
		for min, cnt := range mins {
			if cnt > mostMin.Max {
				mostMin.Max = cnt
				mostMin.Minute = min
			}
		}
		fmt.Println("first:", most.Guard*uint32(mostMin.Minute))
	}

	{ // second
		mostmins := make([]GuardMostMin, 0, 32)
		for guard, times := range table {
			mins := map[byte]uint32{}
			for _, rng := range times {
				for cur := rng.Start; cur.Before(rng.End); cur = cur.Add(time.Minute) {
					min := byte(cur.Minute())
					if cnt, ok := mins[min]; ok {
						mins[min] = cnt + 1
					} else {
						mins[min] = 1
					}
				}
			}
			var mostMin MostMin
			for min, cnt := range mins {
				if cnt > mostMin.Max {
					mostMin.Max = cnt
					mostMin.Minute = min
				}
			}
			mostmins = append(mostmins, GuardMostMin{
				Guard:   guard,
				MostMin: mostMin,
			})
		}
		sort.Sort(GuardMostMins(mostmins))
		guard := mostmins[0]
		fmt.Println("second:", guard.Guard*uint32(guard.Minute))
	}
}
