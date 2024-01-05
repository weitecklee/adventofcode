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
	input := strings.Split(string(data), "\n")
	fmt.Println(part1(parseInput(input)))
}

func parseInput(input []string) map[[2]int]bool {
	bugs := map[[2]int]bool{}
	for j, row := range input {
		for i, c := range row {
			if c == '#' {
				bugs[[2]int{i, j}] = true
			}
		}
	}
	return bugs
}

func state(bugs map[[2]int]bool) int {
	stateNum := 0.0
	for coord := range bugs {
		stateNum += math.Pow(2, float64(5*coord[1]+coord[0]))
	}
	return int(stateNum)
}

func aBugsLife(bugs map[[2]int]bool) map[[2]int]bool {
	bugs2 := map[[2]int]bool{}
	adjacents := map[[2]int]int{}
	checks := [][2]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}
	for coord := range bugs {
		for _, check := range checks {
			toCheck := [2]int{coord[0] + check[0], coord[1] + check[1]}
			if toCheck[0] < 0 || toCheck[1] < 0 || toCheck[0] > 4 || toCheck[1] > 4 {
				continue
			}
			adjacents[toCheck]++
		}
	}
	for coord := range bugs {
		if adjacents[coord] == 1 {
			bugs2[coord] = true
		}
	}
	for coord, n := range adjacents {
		if !bugs[coord] && (n == 1 || n == 2) {
			bugs2[coord] = true
		}
	}
	return bugs2
}

func part1(bugs map[[2]int]bool) int {
	stateHistory := map[int]bool{}
	stateHistory[state(bugs)] = true
	for {
		bugs = aBugsLife(bugs)
		currState := state(bugs)
		if stateHistory[currState] {
			return currState
		}
		stateHistory[currState] = true
	}
	return -1
}
