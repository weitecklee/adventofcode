package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n")
	slope := []int{1, 3}
	fmt.Println(part1(slope, input))
	slopes := [][]int{
		{1, 1},
		{1, 3},
		{1, 5},
		{1, 7},
		{2, 1},
	}
	fmt.Println(part2(slopes, input))
}

func sloper(slope []int, input []string) int {
	pos := [2]int{0, 0}
	h, w := len(input), len(input[0])
	trees := 0
	for pos[0] < h {
		if input[pos[0]][pos[1]] == "#"[0] {
			trees++
		}
		pos[0] += slope[0]
		pos[1] += slope[1]
		if pos[1] >= w {
			pos[1] -= w
		}
	}
	return trees
}

func part1(slope []int, input []string) int {
	return sloper(slope, input)
}

func part2(slopes [][]int, input []string) int {
	trees := 1
	for _, slope := range slopes {
		trees *= sloper(slope, input)
	}
	return trees
}
