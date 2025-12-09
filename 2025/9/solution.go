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
}

func parseInput(data []string) [][2]int {
	points := make([][2]int, len(data))
	for i, line := range data {
		points[i] = [2]int(utils.ExtractInts(line))
	}
	return points
}

func calcRectangle(p1, p2 [2]int) int {
	return (utils.AbsInt(p1[0]-p2[0]) + 1) * (utils.AbsInt(p1[1]-p2[1]) + 1)
}

func part1(puzzleInput [][2]int) int {
	res := 0
	for i, p1 := range puzzleInput {
		for _, p2 := range puzzleInput[i+1:] {
			size := calcRectangle(p1, p2)
			if size > res {
				res = size
			}
		}
	}
	return res
}
