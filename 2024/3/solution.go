package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
)

func main() {
	_, filename, _, _ := runtime.Caller(0)
	dirname := filepath.Dir(filename)
	inputFilePath := filepath.Join(dirname, "input.txt")
	data, err := os.ReadFile(inputFilePath)
	if err != nil {
		panic(err)
	}
	puzzleInput := string(data)
	fmt.Println(part1(puzzleInput))
	fmt.Println(part2(puzzleInput))
}

func part1(puzzleInput string) int {
	res := 0
	mulRegex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := mulRegex.FindAllStringSubmatch(puzzleInput, -1)
	for _, match := range matches {
		var n1, n2 int
		var err error
		if n1, err = strconv.Atoi(match[1]); err != nil {
			panic(err)
		}
		if n2, err = strconv.Atoi(match[2]); err != nil {
			panic(err)
		}
		res += n1 * n2
	}

	return res
}

func part2(puzzleInput string) int {
	removeRegex := regexp.MustCompile(`(?s)don\'t\(\).*?(?:do\(\)|$)`)
	puzzleInput2 := removeRegex.ReplaceAllString(puzzleInput, "")
	return part1(puzzleInput2)
}
