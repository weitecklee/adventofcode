package main

import (
	"fmt"
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
	puzzleInput := strings.Split((string(data)), "\n")
	walls, start, h, w := parseInput(puzzleInput)
	visited := part1(walls, start, h, w)
	fmt.Println(len(visited))
	fmt.Println(part2(visited, walls, start, h, w))
}

func parseInput(puzzleInput []string) (map[[2]int]struct{}, [2]int, int, int) {
	start := [2]int{0, 0}
	walls := make(map[[2]int]struct{})
	for i, row := range puzzleInput {
		for j := range row {
			if row[j] == '^' {
				start = [2]int{i, j}
			} else if row[j] == '#' {
				walls[[2]int{i, j}] = struct{}{}
			}
		}
	}
	h := len(puzzleInput)
	w := len(puzzleInput[0])
	return walls, start, h, w
}

func part1(walls map[[2]int]struct{}, start [2]int, h, w int) map[[2]int]struct{} {
	visited := make(map[[2]int]struct{})
	pos := start
	direction := [2]int{-1, 0}
	for {
		visited[pos] = struct{}{}
		pos2 := [2]int{pos[0] + direction[0], pos[1] + direction[1]}
		if pos2[0] < 0 || pos2[1] < 0 || pos2[0] > h-1 || pos2[1] > w-1 {
			break
		}
		if _, ok := walls[pos2]; ok {
			direction[0], direction[1] = direction[1], -direction[0]
		} else {
			pos = pos2
		}
	}
	return visited
}

func isStuckInALoop(walls map[[2]int]struct{}, start [2]int, h, w int) bool {
	visited := make(map[[4]int]struct{})
	pos := start
	direction := [2]int{-1, 0}
	for {
		state := [4]int{pos[0], pos[1], direction[0], direction[1]}
		if _, ok := visited[state]; ok {
			return true
		}
		visited[state] = struct{}{}
		pos2 := [2]int{pos[0] + direction[0], pos[1] + direction[1]}
		if pos2[0] < 0 || pos2[1] < 0 || pos2[0] > h-1 || pos2[1] > w-1 {
			break
		}
		if _, ok := walls[pos2]; ok {
			direction[0], direction[1] = direction[1], -direction[0]
		} else {
			pos = pos2
		}
	}
	return false
}

func part2(visited, walls map[[2]int]struct{}, start [2]int, h, w int) int {
	res := 0
	for pos := range visited {
		if pos == start {
			continue
		}
		walls[pos] = struct{}{}
		if isStuckInALoop(walls, start, h, w) {
			res++
		}
		delete(walls, pos)
	}
	return res
}
