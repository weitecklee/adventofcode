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
}

type Machine struct {
	lightDiagram []bool
	buttons      [][]int
	joltageReqs  []int
}

func NewMachine(line string) *Machine {
	parts := strings.Split(line, " ")
	lightDiagram := make([]bool, len(parts[0])-2)
	for i, ch := range parts[0][1 : len(parts[0])-1] {
		lightDiagram[i] = ch == '#'
	}
	buttons := make([][]int, len(parts)-2)

	for i, s := range parts[1 : len(parts)-1] {
		buttons[i] = utils.ExtractInts(s)
	}

	joltageReqs := utils.ExtractInts(parts[len(parts)-1])

	m := Machine{
		lightDiagram,
		buttons,
		joltageReqs,
	}

	return &m
}

func (m *Machine) doLightsMatch(lights []bool) bool {
	for i := range lights {
		if lights[i] != m.lightDiagram[i] {
			return false
		}
	}
	return true
}

type Value struct {
	lights  []bool
	presses int
}

func (m *Machine) calcFewestPresses() int {
	res := math.MaxInt
	queue := utils.NewMinHeap[Value]()
	heap.Push(queue, &utils.Item[Value]{
		Priority: 0,
		Value:    Value{make([]bool, len(m.lightDiagram)), 0},
	})

	memo := make(map[int]int)

	for len(queue.PriorityQueue) > 0 {
		item := heap.Pop(queue).(*utils.Item[Value])
		lights := item.Value.lights
		presses := item.Value.presses

		lightsInt := lightsToInt(lights)

		if v, ok := memo[lightsInt]; ok && v <= presses {
			continue
		}

		memo[lightsInt] = presses

		if presses >= res {
			continue
		}

		if m.doLightsMatch(lights) {
			if presses < res {
				res = presses
			}
			continue
		}

		for _, button := range m.buttons {
			tmp := slices.Clone(lights)
			for _, n := range button {
				tmp[n] = !tmp[n]
			}
			heap.Push(queue, &utils.Item[Value]{
				Priority: presses + 1,
				Value:    Value{tmp, presses + 1},
			})
		}
	}
	return res
}

func lightsToInt(lights []bool) int {
	res := 0
	for i, light := range lights {
		if light {
			res += 1 << i
		}
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
		res += m.calcFewestPresses()
	}
	return res
}
