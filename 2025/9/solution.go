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

func doesIntersect(p1, p2 [2]int, edges [][4]int) bool {
	x1, y1, x2, y2 := p1[0], p1[1], p2[0], p2[1]
	if x1 > x2 {
		x1, x2 = x2, x1
	}
	if y1 > y2 {
		y1, y2 = y2, y1
	}
	for _, edge := range edges {
		if x1 < edge[2] && x2 > edge[0] && y1 < edge[3] && y2 > edge[1] {
			return true
		}
	}
	return false
}

func part2(puzzleInput [][2]int) int {
	l := len(puzzleInput)
	edges := make([][4]int, l)
	for i := range l - 1 {
		x1, y1, x2, y2 := puzzleInput[i][0], puzzleInput[i][1], puzzleInput[i+1][0], puzzleInput[i+1][1]
		if x1 > x2 {
			x1, x2 = x2, x1
		}
		if y1 > y2 {
			y1, y2 = y2, y1
		}
		edges[i] = [4]int{x1, y1, x2, y2}
	}
	x1, y1, x2, y2 := puzzleInput[l-1][0], puzzleInput[l-1][1], puzzleInput[0][0], puzzleInput[0][1]
	if x1 > x2 {
		x1, x2 = x2, x1
	}
	if y1 > y2 {
		y1, y2 = y2, y1
	}
	edges[l-1] = [4]int{x1, y1, x2, y2}
	res := 0
	for i, p1 := range puzzleInput {
		for _, p2 := range puzzleInput[i+1:] {
			size := calcRectangle(p1, p2)
			if size > res && !doesIntersect(p1, p2, edges) {
				res = size
			}
		}
	}
	return res
}
