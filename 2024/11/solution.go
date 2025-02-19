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
	puzzleInput := strings.Split(string(data), " ")
	stoneMap := parseInput(&puzzleInput)
	fmt.Println(solve(stoneMap))
}

func parseInput(puzzleInput *[]string) *map[int]int {
	stoneMap := make(map[int]int)
	for _, s := range *puzzleInput {
		if n, err := strconv.Atoi(s); err == nil {
			stoneMap[n]++
		}
	}
	return &stoneMap
}

func blink(stoneMap *map[int]int) {
	newStoneMap := make(map[int]int)
	for n, k := range *stoneMap {
		if n == 0 {
			newStoneMap[1] += k
		} else if s := strconv.Itoa(n); len(s)%2 == 0 {
			if n1, err := strconv.Atoi(s[0 : len(s)/2]); err == nil {
				newStoneMap[n1] += k
			} else {
				panic(err)
			}
			if n2, err := strconv.Atoi(s[len(s)/2:]); err == nil {
				newStoneMap[n2] += k
			} else {
				panic(err)
			}
		} else {
			newStoneMap[n*2024] += k
		}
	}
	*stoneMap = newStoneMap
}

func nStones(stoneMap *map[int]int) int {
	res := 0
	for _, k := range *stoneMap {
		res += k
	}
	return res
}

func solve(stoneMap *map[int]int) (int, int) {
	for i := 0; i < 25; i++ {
		blink(stoneMap)
	}
	part1 := nStones(stoneMap)
	for i := 25; i < 75; i++ {
		blink(stoneMap)
	}
	part2 := nStones(stoneMap)
	return part1, part2
}
