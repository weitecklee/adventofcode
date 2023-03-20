package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n")
	fmt.Println(part1(parseInput(input)))
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
