package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
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
	puzzleInput := parseInput(strings.Split(string(data), ", "))
	fmt.Println(part1(puzzleInput))
	fmt.Println(part2(puzzleInput))
}

type Instruction struct {
	turn     byte
	distance int
}

func parseInput(data []string) []*Instruction {
	instructions := make([]*Instruction, len(data))
	for i, step := range data {
		turn := step[0]
		distance, err := strconv.Atoi(step[1:])
		if err != nil {
			panic(err)
		}
		instructions[i] = &Instruction{turn, distance}
	}
	return instructions
}

func part1(puzzleInput []*Instruction) int {
	pos := [2]int{0, 0}
	dir := [2]int{0, 1}
	for _, instruction := range puzzleInput {
		if instruction.turn == 'L' {
			dir[0], dir[1] = -dir[1], dir[0]
		} else {
			dir[0], dir[1] = dir[1], -dir[0]
		}
		pos[0] += dir[0] * instruction.distance
		pos[1] += dir[1] * instruction.distance
	}
	return utils.AbsInt(pos[0]) + utils.AbsInt(pos[1])
}

func part2(puzzleInput []*Instruction) int {
	pos := [2]int{0, 0}
	dir := [2]int{0, 1}
	history := make(map[[2]int]struct{})
	history[[2]int{0, 0}] = struct{}{}
	for _, instruction := range puzzleInput {
		if instruction.turn == 'L' {
			dir[0], dir[1] = -dir[1], dir[0]
		} else {
			dir[0], dir[1] = dir[1], -dir[0]
		}
		for range instruction.distance {
			pos[0] += dir[0]
			pos[1] += dir[1]
			if _, exists := history[pos]; exists {
				return utils.AbsInt(pos[0]) + utils.AbsInt(pos[1])
			}
			history[pos] = struct{}{}
		}
	}
	return -1
}
