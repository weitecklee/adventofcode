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
	fmt.Println(solve(sizes))
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

func recur(sizes []int, total int, nContainers int, ways *map[int]int) int {
	n := 0
	for i := 0; i < len(sizes) && sizes[i] <= total; i++ {
		if total == sizes[i] {
			(*ways)[nContainers+1]++
			n++
		} else {
			n += recur(sizes[i+1:], total-sizes[i], nContainers+1, ways)
		}
	}
	return n
}

func solve(sizes []int) (int, int) {
	ways := map[int]int{}
	n := recur(sizes, 150, 0, &ways)
	min := len(sizes)
	res := 0
	for i, j := range ways {
		if i < min {
			min = i
			res = j
		}
	}
	return n, res
}
