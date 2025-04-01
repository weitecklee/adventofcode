package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
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
	puzzleInput := parseInput(strings.Split(string(data), "\n"))
	fmt.Println(part1(puzzleInput))
	fmt.Println(part2(puzzleInput))
}

func parseInput(data []string) []int {
	puzzleInput := make([]int, len(data))
	for i, s := range data {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		puzzleInput[i] = n
	}
	return puzzleInput
}

func part1(puzzleInput []int) int {
	res := 0
	for _, n := range puzzleInput {
		res += n
	}
	return res
}

func part2(puzzleInput []int) int {
	res := 0
	history := make(map[int]struct{})
	history[0] = struct{}{}
	for {
		for _, n := range puzzleInput {
			res += n
			if _, exists := history[res]; exists {
				return res
			}
			history[res] = struct{}{}
		}
	}
}
