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

func part1(puzzleInput []string) int {
	var twins, triplets int
	for _, s := range puzzleInput {
		letterMap := make(map[rune]int, len(s))
		for _, r := range s {
			letterMap[r]++
		}
		var twinFound, tripletFound bool
		for _, n := range letterMap {
			if n == 2 {
				twinFound = true
			} else if n == 3 {
				tripletFound = true
			}
		}
		if twinFound {
			twins++
		}
		if tripletFound {
			triplets++
		}
	}
	return twins * triplets
}

func part2(puzzleInput []string) string {
	for i, s1 := range puzzleInput {
		for _, s2 := range puzzleInput[i+1:] {
			var diffs, idx int
			for k := range s1 {
				if s1[k] != s2[k] {
					diffs++
					idx = k
					if diffs > 1 {
						break
					}
				}
			}
			if diffs == 1 {
				return fmt.Sprintf("%s%s", s1[:idx], s1[idx+1:])
			}
		}
	}
	return ""
}
