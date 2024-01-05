package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
)

func main() {
	_, filename, _, _ := runtime.Caller(0)
	dirname := filepath.Dir(filename)
	inputFilePath := filepath.Join(dirname, "input.txt")
	data, err := os.ReadFile(inputFilePath)
	if err != nil {
		panic(err)
	}
	input := string(data)
	parsedInput := parseInput(input)
	fmt.Println(part1(parsedInput))
	fmt.Println(part2(parsedInput))
}

func parseInput(input string) []int {
	res := []int{}
	for _, c := range input {
		n, _ := strconv.Atoi(string(c))
		res = append(res, n)
	}
	return res
}

func part1(input []int) int {
	res := 0
	if input[0] == input[len(input)-1] {
		res = input[0]
	}
	for i := 0; i < len(input)-1; i++ {
		if input[i] == input[i+1] {
			res += input[i]
		}
	}
	return res
}

func part2(input []int) int {
	res := 0
	n := len(input)
	input = append(input, input...)
	for i := 0; i < n; i++ {
		if input[i] == input[i+n/2] {
			res += input[i]
		}
	}
	return res
}
