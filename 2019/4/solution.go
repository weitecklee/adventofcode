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
	puzzleInput := parseInput(string(data))
	fmt.Println(part1(puzzleInput))
	fmt.Println(part2(puzzleInput))
}

func parseInput(data string) [2]int {
	parts := strings.Split(data, "-")
	var nums [2]int
	if n, err := strconv.Atoi(parts[0]); err != nil {
		panic(err)
	} else {
		nums[0] = n
	}
	if n, err := strconv.Atoi(parts[1]); err != nil {
		panic(err)
	} else {
		nums[1] = n
	}
	return nums
}

func isValid(n int, stricter bool) bool {
	a, b := n/10, n%10
	twins := make(map[int]int)
	for a > 0 {
		c, d := a/10, a%10
		if d > b {
			return false
		}
		if d == b {
			twins[d]++
		}
		a, b = c, d
	}
	if !stricter {
		return len(twins) > 0
	}
	for _, n := range twins {
		if n == 1 {
			return true
		}
	}
	return false
}

func part1(puzzleInput [2]int) int {
	res := 0
	for n := puzzleInput[0]; n <= puzzleInput[1]; n++ {
		if isValid(n, false) {
			res++
		}
	}
	return res
}

func part2(puzzleInput [2]int) int {
	res := 0
	for n := puzzleInput[0]; n <= puzzleInput[1]; n++ {
		if isValid(n, true) {
			res++
		}
	}
	return res
}
