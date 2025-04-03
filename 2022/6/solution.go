package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
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

func findMarker(puzzleInput string, nDistinct int) int {
	letters := make(map[byte]int, nDistinct)
	for i := range puzzleInput {
		letters[puzzleInput[i]]++
		if i >= nDistinct {
			letters[puzzleInput[i-nDistinct]]--
			if letters[puzzleInput[i-nDistinct]] == 0 {
				delete(letters, puzzleInput[i-nDistinct])
			}
		}
		if len(letters) == nDistinct {
			return i + 1
		}
	}
	return -1
}

func part1(puzzleInput string) int {
	return findMarker(puzzleInput, 4)
}

func part2(puzzleInput string) int {
	return findMarker(puzzleInput, 14)
}
