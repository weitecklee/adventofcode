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
	weights := parseInput(input)
	fmt.Println(solve(weights, 3))
	fmt.Println(solve(weights, 4))
}

func parseInput(input []string) []int {
	weights := []int{}
	for _, line := range input {
		n, _ := strconv.Atoi(line)
		weights = append(weights, n)
	}
	return weights
}

func makeGroups(weights []int, target int, curr []int) [][]int {
	groups := [][]int{}
	if target == 0 {
		return [][]int{curr}
	}
	for i := range weights {
		if target < weights[i] {
			break
		}
		groups2 := makeGroups(weights[i+1:], target-weights[i], append(curr, weights[i]))
		if groups2 != nil && len(groups2) > 0 {
			groups = append(groups, groups2...)
		}
	}
	return groups
}

/*
	Currently, makeGroups returns some wrong groups that repeat weights and exceed target.
	For now, check is included to make sure weight of group is equal to target.
*/

func solve(weights []int, parts int) int {
	totalWeight := 0
	for _, weight := range weights {
		totalWeight += weight
	}
	target := totalWeight / parts
	groups := makeGroups(weights, target, []int{})
	size := len(weights)
	minQE := 0
	for _, group := range groups {
		sum := 0
		prod := 1
		for _, weight := range group {
			sum += weight
			prod *= weight
		}
		if sum == target {
			if len(group) < size {
				size = len(group)
				minQE = prod
			} else if len(group) == size && prod < minQE {
				minQE = prod
			}
		}
	}
	return minQE
}
