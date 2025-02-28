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
	numbers := make([]int, 0, len(data))
	for _, s := range data {
		if n, err := strconv.Atoi(s); err != nil {
			panic(err)
		} else {
			numbers = append(numbers, n)
		}
	}
	return numbers
}

func calculateFuel(mass int) int {
	return mass/3 - 2
}

func part1(puzzleInput []int) int {
	res := 0
	for _, mass := range puzzleInput {
		res += calculateFuel(mass)
	}
	return res
}

func part2(puzzleInput []int) int {
	res := 0
	for _, mass := range puzzleInput {
		fuel := calculateFuel(mass)
		for fuel > 0 {
			res += fuel
			fuel = calculateFuel(fuel)
		}
	}
	return res
}
