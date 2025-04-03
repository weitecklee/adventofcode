package main

import (
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

func parseInput(data []string) [][2][2]int {
	puzzleInput := make([][2][2]int, len(data))
	for i, line := range data {
		nums := utils.ExtractInts(line)
		puzzleInput[i] = [2][2]int{{nums[0], nums[1]}, {nums[2], nums[3]}}
	}
	return puzzleInput
}

func fullyContained(pairs [2][2]int) bool {
	a, b := pairs[0][0], pairs[0][1]
	c, d := pairs[1][0], pairs[1][1]
	return (c <= a && c <= b && d >= a && d >= b) || (a <= c && a <= d && b >= c && b >= d)
}

func hasOverlap(pairs [2][2]int) bool {
	a, b := pairs[0][0], pairs[0][1]
	c, d := pairs[1][0], pairs[1][1]
	return (a >= c && a <= d) || (b >= c && b <= d) || (c >= a && c <= b) || (d >= a && d <= b)
}

func part1(puzzleInput [][2][2]int) int {
	res := 0
	for _, pairs := range puzzleInput {
		if fullyContained(pairs) {
			res++
		}
	}
	return res
}

func part2(puzzleInput [][2][2]int) int {
	res := 0
	for _, pairs := range puzzleInput {
		if hasOverlap(pairs) {
			res++
		}
	}
	return res
}
