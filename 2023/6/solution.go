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
	part1Input, part2Input := parseInput(strings.Split(string(data), "\n"))
	fmt.Println(part1(part1Input))
	fmt.Println(part2(part2Input))
}

func parseInput(data []string) ([][]int, []int) {
	numRegex := regexp.MustCompile(`\d+`)
	match0 := numRegex.FindAllString(data[0], -1)
	match1 := numRegex.FindAllString(data[1], -1)
	times := make([]int, len(match0))
	distances := make([]int, len(match1))
	for i, s := range match0 {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		times[i] = n
	}
	for i, s := range match1 {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		distances[i] = n
	}
	time2, err := strconv.Atoi(strings.Join(match0, ""))
	if err != nil {
		panic(err)
	}
	distance2, err := strconv.Atoi(strings.Join(match1, ""))
	if err != nil {
		panic(err)
	}
	return [][]int{times, distances}, []int{time2, distance2}
}

func waysToWin(time, distance int) int {
	var lo, hi int
	for (time-lo)*lo <= distance {
		lo++
	}
	hi = time
	for (time-hi)*hi <= distance {
		hi--
	}
	return hi - lo + 1
}

func part1(puzzleInput [][]int) int {
	res := 1
	for i := range puzzleInput[0] {
		res *= waysToWin(puzzleInput[0][i], puzzleInput[1][i])
	}
	return res
}

func part2(puzzleInput []int) int {
	return waysToWin(puzzleInput[0], puzzleInput[1])
}
