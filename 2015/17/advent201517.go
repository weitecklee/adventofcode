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
	sizes := parseInput(input)
	fmt.Println(part1(sizes))
}

func parseInput(input []string) []int {
	sizes := []int{}
	for _, line := range input {
		n, _ := strconv.Atoi(line)
		sizes = append(sizes, n)
	}
	sort.Ints(sizes)
	return sizes
}

func recur(sizes []int, total int) int {
	n := 0
	for i := 0; i < len(sizes) && sizes[i] <= total; i++ {
		if total == sizes[i] {
			n++
		} else {
			n += recur(sizes[i+1:], total-sizes[i])
		}
	}
	return n
}

func part1(sizes []int) int {
	return recur(sizes, 150)
}
