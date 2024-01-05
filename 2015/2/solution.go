package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
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
	boxes := parseInput(input)
	fmt.Println(part1(boxes))
	fmt.Println(part2(boxes))
}

func parseInput(input []string) [][]int {
	numbers := [][]int{}
	for _, line := range input {
		parts := strings.Split(line, "x")
		nums := []int{}
		for _, part := range parts {
			n, _ := strconv.Atoi(part)
			nums = append(nums, n)
		}
		sort.Ints(nums)
		numbers = append(numbers, nums)
	}
	return numbers
}

func part1(boxes [][]int) int {
	area := 0
	for _, box := range boxes {
		area += 3*box[0]*box[1] + 2*box[0]*box[2] + 2*box[1]*box[2]
	}
	return area
}

func part2(boxes [][]int) int {
	feet := 0
	for _, box := range boxes {
		feet += 2*box[0] + 2*box[1] + box[0]*box[1]*box[2]
	}
	return feet
}
