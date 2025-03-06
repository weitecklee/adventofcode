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

func parseInput(data []string) [][]int {
	puzzleInput := make([][]int, len(data))
	for i, line := range data {
		s := strings.Split(line, " ")
		nums := make([]int, len(s))
		for j, n := range s {
			if num, err := strconv.Atoi(n); err != nil {
				panic(err)
			} else {
				nums[j] = num
			}
		}
		puzzleInput[i] = nums
	}
	return puzzleInput
}

func extrapolate(nums []int, nextValue bool) int {
	pyramid := [][]int{nums}
	for {
		curr := pyramid[len(pyramid)-1]
		next := make([]int, len(curr)-1)
		allZeros := true
		for i := range len(curr) - 1 {
			next[i] = curr[i+1] - curr[i]
			if next[i] != 0 {
				allZeros = false
			}
		}
		if allZeros {
			break
		}
		pyramid = append(pyramid, next)
	}
	res := 0
	for i := len(pyramid) - 1; i >= 0; i-- {
		if nextValue {
			res += pyramid[i][len(pyramid[i])-1]
		} else {
			res = pyramid[i][0] - res
		}
	}
	return res
}

func part1(puzzleInput [][]int) int {
	res := 0
	for _, nums := range puzzleInput {
		res += extrapolate(nums, true)
	}
	return res
}

func part2(puzzleInput [][]int) int {
	res := 0
	for _, nums := range puzzleInput {
		res += extrapolate(nums, false)
	}
	return res
}
