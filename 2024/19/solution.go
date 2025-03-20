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
	designs := parseInput(strings.Split(string(data), "\n\n"))
	fmt.Println(solve(designs))
}

var (
	patternMap = make(map[string]struct{})
	countMemo  = make(map[string]int)
)

func parseInput(data []string) []string {
	patterns := strings.Split(data[0], ", ")
	designs := strings.Split(data[1], "\n")
	patternMap[""] = struct{}{}
	for _, pattern := range patterns {
		patternMap[pattern] = struct{}{}
	}
	countMemo[""] = 1
	return designs
}

func countWays(design string) int {
	if n, ok := countMemo[design]; ok {
		return n
	}
	ways := 0
	for i := range design {
		if _, ok := patternMap[design[:i+1]]; ok {
			ways += countWays(design[i+1:])
		}
	}
	countMemo[design] = ways
	return ways
}

func solve(designs []string) (int, int) {
	part1 := 0
	part2 := 0
	for _, design := range designs {
		ways := countWays(design)
		if ways > 0 {
			part1++
			part2 += ways
		}
	}
	return part1, part2
}
