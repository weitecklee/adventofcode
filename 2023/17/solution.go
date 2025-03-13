package main

import (
	"container/heap"
	"fmt"
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
	puzzleInput := parseInput(strings.Split(string(data), "\n"))
	fmt.Println(part1(puzzleInput))
	fmt.Println(part2(puzzleInput))
}

type Value struct {
	heatLoss           int
	pos                [2]int
	dir                [2]int
	stepsSinceLastTurn int
}

type State struct {
	pos                [2]int
	dir                [2]int
	stepsSinceLastTurn int
}

var directions = [][2]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func parseInput(data []string) [][]int {
	puzzleInput := make([][]int, len(data))
	for i, line := range data {
		row := make([]int, len(line))
		for j, ch := range line {
			row[j] = int(ch - '0')
		}
		puzzleInput[i] = row
	}
	return puzzleInput
}

func distanceBetweenPoints(p1, p2 [2]int) int {
	return utils.AbsInt(p1[0]-p2[0]) + utils.AbsInt(p1[1]-p2[1])
}

func solve(puzzleInput [][]int, minSteps, maxSteps int) int {
	queue := utils.NewMinHeap[Value]()
	heap.Push(queue, &utils.Item[Value]{
		Priority: 0,
		Value:    Value{0, [2]int{0, 0}, [2]int{0, 0}, minSteps},
	})
	visited := make(map[State]int)

	rMax := len(puzzleInput) - 1
	cMax := len(puzzleInput[0]) - 1
	target := [2]int{rMax, cMax}
	var r, c, r2, c2, heatLoss, steps int
	var pos, dir [2]int
	var state State

	for len(queue.PriorityQueue) > 0 {
		item := heap.Pop(queue).(*utils.Item[Value])
		if item.Value.pos == target {
			return item.Value.heatLoss
		}
		r, c = item.Value.pos[0], item.Value.pos[1]
		dir = item.Value.dir
		for _, d := range directions {
			if d[0] == -dir[0] && d[1] == -dir[1] {
				continue
			}
			r2, c2 = r+d[0], c+d[1]
			if r2 < 0 || c2 < 0 || r2 > rMax || c2 > cMax {
				continue
			}
			steps = 1
			if d[0] == dir[0] && d[1] == dir[1] {
				steps = item.Value.stepsSinceLastTurn + 1
			} else if item.Value.stepsSinceLastTurn < minSteps {
				continue
			}
			if steps > maxSteps {
				continue
			}
			pos = [2]int{r2, c2}
			heatLoss = item.Value.heatLoss + puzzleInput[r2][c2]
			state = State{pos, d, steps}
			if h, ok := visited[state]; ok && h <= heatLoss {
				continue
			}
			visited[state] = heatLoss
			heap.Push(queue,
				&utils.Item[Value]{
					Priority: heatLoss + distanceBetweenPoints(pos, target),
					Value:    Value{heatLoss, pos, d, steps},
				})
		}
	}
	return -1
}

func part1(puzzleInput [][]int) int {
	return solve(puzzleInput, 0, 3)
}

func part2(puzzleInput [][]int) int {
	return solve(puzzleInput, 4, 10)
}
