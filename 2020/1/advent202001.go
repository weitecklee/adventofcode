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
	parsedInput := []int{}
	for _, s := range input {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		parsedInput = append(parsedInput, n)
	}
	fmt.Println(part1(parsedInput))
	fmt.Println(part2(parsedInput))
}

func part1(input []int) int {
	seen := make(map[int]bool)
	for _, n := range input {
		if seen[2020-n] {
			return n * (2020 - n)
		}
		seen[n] = true
	}
	return -1
}

func part2(input []int) int {
	sort.Sort(sort.IntSlice(input))
	z := len(input)
	for i := 0; i < z-2; i++ {
		j := i + 1
		k := z - 1
		for j < k {
			sum := input[i] + input[j] + input[k]
			if sum == 2020 {
				return input[i] * input[j] * input[k]
			}
			if sum < 2020 {
				j++
			} else {
				k--
			}
		}
	}
	return -1
}
