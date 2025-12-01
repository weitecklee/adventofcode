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
	res := make([]int, len(data))
	for i, line := range data {
		n, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		if line[0] == 'L' {
			n *= -1
		}
		res[i] = n
	}
	return res
}

func part1(puzzleInput []int) int {
	res := 0
	dial := 50
	for _, n := range puzzleInput {
		dial += n
		dial %= 100
		if dial == 0 {
			res++
		}
	}
	return res
}

func part2(puzzleInput []int) int {
	dial := 50
	res := 0
	for _, n := range puzzleInput {
		inc := 1
		if n < 0 {
			n *= -1
			inc = -1
		}
		for range n {
			dial += inc
			dial %= 100
			if dial == 0 {
				res++
			}
		}
	}
	return res
}
