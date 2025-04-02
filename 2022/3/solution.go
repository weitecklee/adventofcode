package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
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
	puzzleInput := strings.Split(string(data), "\n")
	fmt.Println(part1(puzzleInput))
	fmt.Println(part2(puzzleInput))
}

func makeSet(s string) map[rune]struct{} {
	set := make(map[rune]struct{}, len(s))
	for _, c := range s {
		set[c] = struct{}{}
	}
	return set
}

func calcPriority(r rune) int {
	if r > 'Z' {
		return int(r-'a') + 1
	}
	return int(r-'A') + 27
}

func part1(puzzleInput []string) int {
	res := 0
	for _, line := range puzzleInput {
		l := len(line)
		set1 := makeSet(line[:l/2])
		set2 := makeSet(line[l/2:])
		for c := range set1 {
			if _, exists := set2[c]; exists {
				res += calcPriority(c)
				break
			}
		}
	}
	return res
}

func part2(puzzleInput []string) int {
	res := 0
	for i := 0; i < len(puzzleInput); i += 3 {
		set1 := makeSet(puzzleInput[i])
		set2 := makeSet(puzzleInput[i+1])
		set3 := makeSet(puzzleInput[i+2])
		for c := range set1 {
			_, exists1 := set2[c]
			_, exists2 := set3[c]
			if exists1 && exists2 {
				res += calcPriority(c)
				break
			}
		}
	}
	return res
}
