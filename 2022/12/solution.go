package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	_, filename, _, _ := runtime.Caller(0)
	dirname := filepath.Dir(filename)
	inputFilePath := filepath.Join(dirname, "input.txt")
	data, err := os.ReadFile(inputFilePath)
	if err != nil {
		panic(err)
	}
	heights, start, end := parseInput(strings.Split(string(data), "\n"))
	fmt.Println(part1(heights, start, end))
	fmt.Println(part2(heights, end))
}

type QueueEntry struct {
	steps int
	pos   [2]int
	ht    rune
}

var directions = [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func parseInput(data []string) ([][]rune, [2]int, [2]int) {
	heights := make([][]rune, len(data))
	var start, end [2]int
	for i, row := range data {
		heightRow := make([]rune, len(row))
		for j, ch := range row {
			if ch == 'S' {
				start = [2]int{i, j}
				heightRow[j] = 0
			} else if ch == 'E' {
				end = [2]int{i, j}
				heightRow[j] = 'z' - 'a'
			} else {
				heightRow[j] = ch - 'a'
			}
		}
		heights[i] = heightRow
	}
	return heights, start, end
}

func part1(heights [][]rune, start, end [2]int) int {
	rMax := len(heights) - 1
	cMax := len(heights[0]) - 1

	queue := []QueueEntry{{0, start, 0}}
	visited := make(map[[2]int]struct{})
	visited[start] = struct{}{}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr.pos == end {
			return curr.steps
		}
		for _, d := range directions {
			pos2 := [2]int{curr.pos[0] + d[0], curr.pos[1] + d[1]}
			if pos2[0] < 0 || pos2[1] < 0 || pos2[0] > rMax || pos2[1] > cMax {
				continue
			}
			ht := heights[pos2[0]][pos2[1]]
			if ht > (curr.ht + 1) {
				continue
			}
			if _, exists := visited[pos2]; exists {
				continue
			}
			visited[pos2] = struct{}{}
			queue = append(queue, QueueEntry{curr.steps + 1, pos2, ht})
		}
	}

	return math.MaxInt
}

func part2(heights [][]rune, end [2]int) int {
	var starts [][2]int
	for r, row := range heights {
		for c, ht := range row {
			if ht == 0 {
				starts = append(starts, [2]int{r, c})
			}
		}
	}

	minSteps := math.MaxInt
	for _, start := range starts {
		res := part1(heights, start, end)
		if res < minSteps {
			minSteps = res
		}
	}

	return minSteps
}
