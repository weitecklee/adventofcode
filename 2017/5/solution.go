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
	input := strings.Split(string(data), "\n")
	fmt.Println(part1(parseInput(input)))
	fmt.Println(part2(parseInput(input)))
}

func parseInput(input []string) []int {
	parsed := []int{}
	for _, line := range input {
		n, _ := strconv.Atoi(line)
		parsed = append(parsed, n)
	}
	return parsed
}

func part1(input []int) int {
	step := 0
	row := 0
	for row >= 0 && row < len(input) {
		step++
		input[row]++
		row += input[row] - 1
	}
	return step
}

func part2(input []int) int {
	step := 0
	row := 0
	for row >= 0 && row < len(input) {
		step++
		oldRow := row
		row += input[row]
		if input[oldRow] >= 3 {
			input[oldRow]--
		} else {
			input[oldRow]++
		}
	}
	return step
}
