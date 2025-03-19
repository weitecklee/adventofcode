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
	fileBlocks := parseInput(string(data))
	fileBlocks2 := slices.Clone(fileBlocks)
	fmt.Println(part1(fileBlocks))
	fmt.Println(part2(fileBlocks2))
}

func parseInput(data string) []int {
	puzzleInput := make([]int, len(data))
	for i, ch := range data {
		puzzleInput[i] = int(ch - '0')
	}
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
	return fileBlocks
}

func computeChecksum(fileBlocks []int) int {
	res := 0
	for i, n := range fileBlocks {
		if n >= 0 {
			res += i * n
		}
	}
	return res
}

func part1(fileBlocks []int) int {
	left := 0
	right := len(fileBlocks) - 1
	for {
		for left < len(fileBlocks) && fileBlocks[left] >= 0 {
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
	return computeChecksum(fileBlocks)
}

func part2(fileBlocks []int) int {
	right := len(fileBlocks) - 1
	for right >= 0 {
		left := right
		for fileBlocks[left] == fileBlocks[right] {
			left--
		}
		size := right - left
		i := 0
		for i < len(fileBlocks) {
			for i < len(fileBlocks) && fileBlocks[i] >= 0 {
				i++
			}
			j := i
			for j < len(fileBlocks) && fileBlocks[j] < 0 {
				j++
			}
			if j-i >= size {
				break
			}
			i = j + 1
		}
		if i < right {
			for k := range size {
				fileBlocks[i+k], fileBlocks[right-k] = fileBlocks[right-k], fileBlocks[i+k]
			}
		}
		right = left
		for right >= 0 && fileBlocks[right] <= 0 {
			right--
		}
	}
	return computeChecksum(fileBlocks)
}
