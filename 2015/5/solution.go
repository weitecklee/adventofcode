package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
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

var (
	vowelRegex  = regexp.MustCompile(`[aeiou].*[aeiou].*[aeiou]`)
	bannedRegex = regexp.MustCompile(`ab|cd|pq|xy`)
)

func hasDouble(s string) bool {
	for i := range len(s) - 1 {
		if s[i] == s[i+1] {
			return true
		}
	}
	return false
}

func isNice(s string) bool {
	return vowelRegex.MatchString(s) && hasDouble(s) && !bannedRegex.MatchString(s)
}

func part1(puzzleInput []string) int {
	res := 0
	for _, line := range puzzleInput {
		if isNice(line) {
			res++
		}
	}
	return res
}

func hasNonOverlappingPairs(s string) bool {
	pairs := make(map[string]int, len(s)-1)
	for i := range len(s) - 1 {
		pair := s[i : i+2]
		if n, exists := pairs[pair]; exists {
			if i-n > 1 {
				return true
			}
		} else {
			pairs[pair] = i
		}
	}
	return false
}

func hasDoubleWithBetween(s string) bool {
	for i := range len(s) - 2 {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}

func isNice2(s string) bool {
	return hasNonOverlappingPairs(s) && hasDoubleWithBetween(s)
}

func part2(puzzleInput []string) int {
	res := 0
	for _, line := range puzzleInput {
		if isNice2(line) {
			res++
		}
	}
	return res
}
