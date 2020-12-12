package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	steps := []Step{}
	{
		bin := bufio.NewReader(os.Stdin)
		for {
			line, err := bin.ReadString('\n')
			if len(line) > 0 {
				line = strings.TrimRight(line, "\r\n")
				amount, err := strconv.ParseInt(line[1:], 10, 64)
				if err != nil {
					panic(err)
				}
				steps = append(steps, Step{
					Type:   StepType(line[0]),
					Amount: amount,
				})
			}
			if err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			}
		}
	}
	{
		start := Location{}
		ship := Ship{
			Location: start,
			Facing:   90,
		}
		for _, step := range steps {
			ship.DoStep(step)
		}
		fmt.Println("first:", manhattanDistance(ship.Location, start))
	}
	{
		start := Location{}
		ship := Ship{
			Location: start,
			Waypoint: Location{
				X: 10,
				Y: -1,
			},
		}
		for _, step := range steps {
			ship.MoveTowardsWaypoint(step)
		}
		fmt.Println("second:", manhattanDistance(ship.Location, start))
	}
}

func manhattanDistance(a, b Location) int64 {
	return abs64(a.X-b.X) + abs64(a.Y-b.Y)
}

func abs64(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}

type StepType byte

const (
	StepTypeNorth   StepType = 'N'
	StepTypeSouth   StepType = 'S'
	StepTypeEast    StepType = 'E'
	StepTypeWest    StepType = 'W'
	StepTypeLeft    StepType = 'L'
	StepTypeRight   StepType = 'R'
	StepTypeForward StepType = 'F'
)

const (
	DirectionNorth int = 0
	DirectionEast  int = 90
	DirectionSouth int = 180
	DirectionWest  int = 270
)

type Step struct {
	Type   StepType
	Amount int64
}

type Location struct {
	X, Y int64
}

type Ship struct {
	Location
	Facing   int
	Waypoint Location
}

func (ship *Ship) DoStep(step Step) {
	switch step.Type {
	case StepTypeNorth:
		ship.Location.Y -= step.Amount
	case StepTypeSouth:
		ship.Location.Y += step.Amount
	case StepTypeEast:
		ship.Location.X += step.Amount
	case StepTypeWest:
		ship.Location.X -= step.Amount
	case StepTypeLeft:
		ship.Facing -= int(step.Amount)
		if ship.Facing < 0 {
			ship.Facing = 360 + ship.Facing
		}
	case StepTypeRight:
		ship.Facing += int(step.Amount)
		if ship.Facing >= 360 {
			ship.Facing = ship.Facing - 360
		}
	case StepTypeForward:
		switch ship.Facing {
		case DirectionNorth:
			ship.Y -= step.Amount
		case DirectionEast:
			ship.X += step.Amount
		case DirectionSouth:
			ship.Y += step.Amount
		case DirectionWest:
			ship.X -= step.Amount
		default:
			panic(fmt.Sprint("bad direction", ship.Facing, step))
		}
	}
}

func (ship *Ship) MoveTowardsWaypoint(step Step) {
	switch step.Type {
	case StepTypeNorth:
		ship.Waypoint.Y -= step.Amount
	case StepTypeSouth:
		ship.Waypoint.Y += step.Amount
	case StepTypeEast:
		ship.Waypoint.X += step.Amount
	case StepTypeWest:
		ship.Waypoint.X -= step.Amount
	case StepTypeLeft:
		switch step.Amount {
		case 90:
			ship.Waypoint.X, ship.Waypoint.Y = ship.Waypoint.Y, -ship.Waypoint.X
		case 180:
			ship.Waypoint.X, ship.Waypoint.Y = -ship.Waypoint.X, -ship.Waypoint.Y
		case 270:
			ship.Waypoint.X, ship.Waypoint.Y = -ship.Waypoint.Y, ship.Waypoint.X
		}
	case StepTypeRight:
		switch step.Amount {
		case 90:
			ship.Waypoint.X, ship.Waypoint.Y = -ship.Waypoint.Y, ship.Waypoint.X
		case 180:
			ship.Waypoint.X, ship.Waypoint.Y = -ship.Waypoint.X, -ship.Waypoint.Y
		case 270:
			ship.Waypoint.X, ship.Waypoint.Y = ship.Waypoint.Y, -ship.Waypoint.X
		}
	case StepTypeForward:
		ship.Location.X += step.Amount * ship.Waypoint.X
		ship.Location.Y += step.Amount * ship.Waypoint.Y
	}
}
