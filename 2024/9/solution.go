package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"slices"
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
}

func parseInput(data string) []int {
	puzzleInput := make([]int, len(data))
	for i, ch := range data {
		puzzleInput[i] = int(ch - '0')
	}
	return puzzleInput
}

func part1(puzzleInput []int) int {
	totalBlocks := 0
	for _, n := range puzzleInput {
		totalBlocks += n
	}
	fileBlocks := slices.Repeat([]int{-1}, totalBlocks)
	j := 0
	for i, n := range puzzleInput {
		if i%2 == 0 {
			for k := range n {
				fileBlocks[j+k] = i / 2
			}
		}
		j += n
	}
	left := 0
	right := totalBlocks - 1
	for {
		for left < totalBlocks && fileBlocks[left] >= 0 {
			left++
		}
		for fileBlocks[right] < 0 {
			right--
		}
		if left >= right {
			break
		}
		fileBlocks[left], fileBlocks[right] = fileBlocks[right], fileBlocks[left]
	}
	res := 0
	for i, n := range fileBlocks {
		if n < 0 {
			break
		}
		res += i * n
	}
	return res
}
