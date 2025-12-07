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
	puzzleInput := strings.Split(string(data), "\n")
	fmt.Println(solve(puzzleInput))
}

func solve(puzzleInput []string) (int, int) {
	var start int
	for col, chr := range puzzleInput[0] {
		if chr == 'S' {
			start = col
			break
		}
	}

	var part1 int
	currRow := make([]int, len(puzzleInput[0]))
	currRow[start] = 1
	for _, row := range puzzleInput[1:] {
		nextRow := make([]int, len(currRow))
		for i, n := range currRow {
			if n == 0 {
				continue
			}
			if row[i] == '^' {
				nextRow[i-1] += n
				nextRow[i+1] += n
				part1 += 1
			} else {
				nextRow[i] += n
			}
		}
		currRow = nextRow
	}

	var part2 int
	for _, n := range currRow {
		part2 += n
	}

	return part1, part2
}
