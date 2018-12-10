package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

type Condition struct {
	Pre  byte
	Then byte
}

type Step struct {
	Name       byte
	Pre        []*Step
	Next       []*Step
	Done       bool
	InProgress bool
}

type ByAlpha []*Step

func (a ByAlpha) Len() int           { return len(a) }
func (a ByAlpha) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAlpha) Less(i, j int) bool { return a[i].Name < a[j].Name }

func contains(steps []*Step, needle *Step) bool {
	for _, step := range steps {
		if step == needle {
			return true
		}
	}
	return false
}

func (step *Step) GetNext() *Step {
	pn := make([]*Step, 0, len(step.Next))
	for _, s := range step.Next {
		if !s.Done && !s.InProgress {
			predone := true
			for _, pre := range s.Pre {
				if !pre.Done {
					predone = false
					break
				}
			}
			if predone && !contains(pn, s) {
				pn = append(pn, s)
			}
		} else if len(s.Next) != 0 {
			ns := s.GetNext()
			if ns != nil && !contains(pn, ns) {
				pn = append(pn, ns)
			}
		}
	}
	if len(pn) == 0 {
		return nil
	}
	sort.Sort(ByAlpha(pn))
	return pn[0]
}

func readInput() ([]Condition, error) {
	fd, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	input := make([]Condition, 0, 64)
	bin := bufio.NewReader(fd)
	for {
		line, err := bin.ReadString('\n')
		if len(line) > 0 {
			line = strings.TrimSpace(line)
			var c Condition
			fmt.Sscanf(line,
				"Step %c must be finished before step %c can begin.",
				&c.Pre, &c.Then)
			input = append(input, c)
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

type Steps []Step

func (ss Steps) FindByName(name byte) *Step {
	for i, s := range ss {
		if s.Name == name {
			return &ss[i]
		}
	}
	return nil
}

func FindNext(begins []*Step) []*Step {
	pn := make([]*Step, 0, 8)
	for _, s := range begins {
		if !s.Done && !s.InProgress {
			predone := true
			for _, pre := range s.Pre {
				if !pre.Done {
					predone = false
					break
				}
			}
			if predone && !contains(pn, s) {
				pn = append(pn, s)
			}
		} else if len(s.Next) != 0 {
			ns := s.GetNext()
			if ns != nil && !contains(pn, ns) {
				pn = append(pn, ns)
			}
		}
	}
	return pn
}

type Worker struct {
	Task     *Step
	Progress uint32
	Work     uint32
}

func main() {
	conds, err := readInput()
	if err != nil {
		panic(err)
	}

	uniq := map[byte]interface{}{}
	for _, cond := range conds {
		if _, ok := uniq[cond.Pre]; !ok {
			uniq[cond.Pre] = true
		}
		if _, ok := uniq[cond.Then]; !ok {
			uniq[cond.Then] = true
		}
	}

	steps := make(Steps, len(uniq))
	{
		i := 0
		for v, _ := range uniq {
			steps[i].Name = v
			i++
		}
	}

	var begins []*Step
	{
		for _, cond := range conds {
			step := steps.FindByName(cond.Pre)
			next := steps.FindByName(cond.Then)
			next.Pre = append(next.Pre, step)
			step.Next = append(step.Next, next)
		}
		for i, step := range steps {
			if len(step.Pre) == 0 {
				begins = append(begins, &steps[i])
			}
		}
	}

	{ // first
		var order strings.Builder
		for {
			pn := FindNext(begins)
			if len(pn) == 0 {
				break
			}
			sort.Sort(ByAlpha(pn))
			step := pn[0]
			step.Done = true
			order.WriteByte(step.Name)
		}
		fmt.Println("first:", order.String())
		for i, _ := range steps {
			steps[i].Done = false
		}
	}

	{ // second
		var sec uint32
		workers := make([]Worker, 5)
		for {
			for i, _ := range workers {
				worker := &workers[i]
				if worker.Task != nil {
					worker.Progress++
					if worker.Progress >= worker.Work {
						worker.Task.Done = true
						worker.Task = nil
					}
				}
			}
			done := true
			for _, step := range steps {
				if !step.Done {
					done = false
					break
				}
			}
			if done {
				break
			}
			pn := FindNext(begins)
			if len(pn) > 0 {
				sort.Sort(ByAlpha(pn))
				for i, _ := range workers {
					worker := &workers[i]
					if worker.Task == nil {
						worker.Task = pn[0]
						worker.Task.InProgress = true
						worker.Work = 60 + uint32(worker.Task.Name-'A') + 1
						worker.Progress = 0
						pn = FindNext(begins)
						if len(pn) == 0 {
							break
						}
						sort.Sort(ByAlpha(pn))
					}
				}
			}
			sec++
		}
		fmt.Println("second:", sec)
	}
}
