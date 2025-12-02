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
	puzzleInput := parseInput(strings.Split(string(data), ","))
	fmt.Println(part1(puzzleInput))
	fmt.Println(part2(puzzleInput))
}

func parseInput(data []string) [][2]int {
	ranges := make([][2]int, len(data))
	for i, pair := range data {
		numbers := strings.Split(pair, "-")
		n1, _ := strconv.Atoi(numbers[0])
		n2, _ := strconv.Atoi(numbers[1])
		ranges[i] = [2]int{n1, n2}
	}
	return ranges
}

func part1(puzzleInput [][2]int) int {
	res := 0
	for _, pair := range puzzleInput {
		for n := pair[0]; n <= pair[1]; n++ {
			if isInvalidId(n) {
				res += n
			}
		}
	}
	return res
}

func part2(puzzleInput [][2]int) int {
	res := 0
	for _, pair := range puzzleInput {
		for n := pair[0]; n <= pair[1]; n++ {
			if isInvalidId2(n) {
				res += n
			}
		}
	}
	return res
}

func isInvalidId(n int) bool {
	s := strconv.Itoa(n)
	return s[0:len(s)/2] == s[len(s)/2:]
}

func isInvalidId2(n int) bool {
	s := strconv.Itoa(n)
	l := len(s)
	for i := 1; i < l; i++ {
		if l%i != 0 {
			continue
		}
		parts := make([]string, l/i)
		for j := 0; j < l/i; j++ {
			parts[j] = s[i*j : i*(j+1)]
		}
		if areAllTheSameString(parts) {
			return true
		}
	}
	return false
}

func areAllTheSameString(ss []string) bool {
	for i := 0; i < len(ss)-1; i++ {
		if ss[i] != ss[i+1] {
			return false
		}
	}
	return true
}
