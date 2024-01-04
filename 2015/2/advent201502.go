package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
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
