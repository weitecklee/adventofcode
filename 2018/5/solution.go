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

func react(polymerString string, throwout rune) int {
	polymer := make([]rune, 0, len(polymerString))
	for _, ch := range polymerString {
		if ch == throwout || ch == throwout+32 {
			continue
		}
		if len(polymer) > 0 &&
			(polymer[len(polymer)-1]-ch == 32 ||
				ch-polymer[len(polymer)-1] == 32) {
			polymer = polymer[:len(polymer)-1]
		} else {
			polymer = append(polymer, ch)
		}
	}
	return len(polymer)
}

func part1(puzzleInput string) int {
	return react(puzzleInput, 0)
}

func part2(puzzleInput string) int {
	min := len(puzzleInput)
	for throwout := rune('A'); throwout <= rune('Z'); throwout++ {
		res := react(puzzleInput, throwout)
		if res < min {
			min = res
		}
	}
	return min
}
