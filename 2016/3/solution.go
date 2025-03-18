package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
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

func parseInput(puzzleInput []string) [][]int {
	numRegex := regexp.MustCompile(`\d+`)
	res := make([][]int, len(puzzleInput))
	for i, line := range puzzleInput {
		match := numRegex.FindAllString(line, -1)
		nums := make([]int, len(match))
		for j, s := range match {
			n, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			nums[j] = n
		}
		res[i] = nums
	}
	return res
}

func isPossibleTriangle(n1, n2, n3 int) bool {
	if n1+n2 <= n3 {
		return false
	}
	if n1+n3 <= n2 {
		return false
	}
	if n2+n3 <= n1 {
		return false
	}
	return true
}

func part1(puzzleInput [][]int) int {
	res := 0
	for _, nums := range puzzleInput {
		if isPossibleTriangle(nums[0], nums[1], nums[2]) {
			res++
		}
	}
	return res
}

func part2(puzzleInput [][]int) int {
	res := 0
	for i := 0; i < len(puzzleInput)-2; i += 3 {
		for j := range len(puzzleInput[0]) {
			if isPossibleTriangle(puzzleInput[i][j], puzzleInput[i+1][j], puzzleInput[i+2][j]) {
				res++
			}
		}
	}
	return res
}
