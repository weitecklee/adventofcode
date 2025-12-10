package main

import (
	"container/heap"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"runtime"
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
	lightDiagram int
	buttons      []int
	joltageReqs  []int
}

func NewMachine2(line string) *Machine {
	parts := strings.Split(line, " ")
	var lightDiagram int
	for i, ch := range parts[0][1 : len(parts[0])-1] {
		if ch == '#' {
			lightDiagram += 1 << i
		}
	}
	buttons := make([]int, len(parts)-2)

	for i, s := range parts[1 : len(parts)-1] {
		nums := utils.ExtractInts(s)
		for _, n := range nums {
			buttons[i] += 1 << n
		}
	}

	joltageReqs := utils.ExtractInts(parts[len(parts)-1])

	m := Machine{
		lightDiagram,
		buttons,
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

		for _, button := range m.buttons {
			heap.Push(queue, &utils.Item[LightValue]{
				Priority: presses + 1,
				Value:    LightValue{lights ^ button, presses + 1},
			})
		}
	}
	return res
}

func parseInput(data []string) []*Machine {
	res := make([]*Machine, len(data))
	for i, line := range data {
		res[i] = NewMachine2(line)
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
