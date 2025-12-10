package main

import (
	"container/heap"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"strings"

	"github.com/draffensperger/golp"
	"github.com/weitecklee/adventofcode/utils"
)

func main() {
	_, filename, _, _ := runtime.Caller(0)
	dirname := filepath.Dir(filename)
	inputFilePath := filepath.Join(dirname, "input.txt")
	data, err := os.ReadFile(inputFilePath)
	if err != nil {
		panic(err)
	}
	machines := parseInput(strings.Split(string(data), "\n"))
	fmt.Println(part1(machines))
	fmt.Println(part2(machines))
}

type Machine struct {
	lightDiagram int
	buttons      [][]int
	buttonInts   []int
	joltageReqs  []int
}

func NewMachine(line string) *Machine {
	parts := strings.Split(line, " ")
	var lightDiagram int
	for i, ch := range parts[0][1 : len(parts[0])-1] {
		if ch == '#' {
			lightDiagram += 1 << i
		}
	}
	buttons := make([][]int, len(parts)-2)
	buttonInts := make([]int, len(parts)-2)

	for i, s := range parts[1 : len(parts)-1] {
		nums := utils.ExtractInts(s)
		buttons[i] = nums
		for _, n := range nums {
			buttonInts[i] += 1 << n
		}
	}

	joltageReqs := utils.ExtractInts(parts[len(parts)-1])

	m := Machine{
		lightDiagram,
		buttons,
		buttonInts,
		joltageReqs,
	}

	return &m
}

type LightValue struct {
	lights  int
	presses int
}

func (m *Machine) fewestPressesForLights() int {
	res := math.MaxInt
	queue := utils.NewMinHeap[LightValue]()
	heap.Push(queue, &utils.Item[LightValue]{
		Priority: 0,
		Value:    LightValue{},
	})

	memo := make(map[int]int)

	for len(queue.PriorityQueue) > 0 {
		item := heap.Pop(queue).(*utils.Item[LightValue])
		lights := item.Value.lights
		presses := item.Value.presses

		if v, ok := memo[lights]; ok && v <= presses {
			continue
		}

		memo[lights] = presses

		if presses >= res {
			continue
		}

		if lights == m.lightDiagram {
			if presses < res {
				res = presses
			}
			continue
		}

		for _, button := range m.buttonInts {
			heap.Push(queue, &utils.Item[LightValue]{
				Priority: presses + 1,
				Value:    LightValue{lights ^ button, presses + 1},
			})
		}
	}
	return res
}

/*
	Use external golp package for solving LP equations (that's actually a wrapper
	for lp_solve C library).
	Each machine is a LP problem where objective is to minimize the number of
	button presses.
	Objective function is just slice of 1's with same length as number of buttons.
	Default behavior of solver is already to minimize	objective function.
	For each button, we also add integer constraint.
	We add constraint for each joltage requirement. We use AddConstraintSparse
	because we only need to add buttons that actually affect that joltage counter.
	The constraint then has to be equal to the joltage requirement.
	Finally, solve it!
*/

func (m *Machine) fewestPressesForJoltage() int {
	lp := golp.NewLP(0, len(m.buttons))

	objFn := make([]float64, len(m.buttons))
	for i := range m.buttons {
		lp.SetInt(i, true)
		objFn[i] = 1.0
	}
	lp.SetObjFn(objFn)

	for i, joltage := range m.joltageReqs {
		entries := make([]golp.Entry, 0, len(m.buttons))
		for b, button := range m.buttons {
			if slices.Contains(button, i) {
				entries = append(entries, golp.Entry{
					Col: b,
					Val: 1.0,
				})
			}
		}
		lp.AddConstraintSparse(entries, golp.EQ, float64(joltage))
	}

	lp.Solve()

	presses := lp.Variables()

	res := 0
	for _, n := range presses {
		res += int(n)
	}

	return res
}

func parseInput(data []string) []*Machine {
	res := make([]*Machine, len(data))
	for i, line := range data {
		res[i] = NewMachine(line)
	}
	return res
}

func part1(machines []*Machine) int {
	res := 0
	for _, m := range machines {
		res += m.fewestPressesForLights()
	}
	return res
}

func part2(machines []*Machine) int {
	res := 0
	for _, m := range machines {
		res += m.fewestPressesForJoltage()
	}
	return res
}
