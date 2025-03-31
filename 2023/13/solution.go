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
	puzzleInput := parseInput(strings.Split(string(data), "\n\n"))
	fmt.Println(part1(puzzleInput))
	fmt.Println(part2(puzzleInput))
}

type Pattern []string

func (p Pattern) FindReflection(tolerance int) (int, int) {
	var smudges int
	for r := 1; r < len(p); r++ {
		smudges = 0
		for dr := 1; r-dr >= 0 && r+dr-1 < len(p); dr++ {
			for c := range p[0] {
				if p[r-dr][c] != p[r+dr-1][c] {
					smudges++
					if smudges > tolerance {
						break
					}
				}
			}
			if smudges > tolerance {
				break
			}
		}
		if smudges == tolerance {
			return r, 0
		}
	}
	for c := 1; c < len(p[0]); c++ {
		smudges = 0
		for dc := 1; c-dc >= 0 && c+dc-1 < len(p[0]); dc++ {
			for r := range p {
				if p[r][c-dc] != p[r][c+dc-1] {
					smudges++
					if smudges > tolerance {
						break
					}
				}
			}
			if smudges > tolerance {
				break
			}
		}
		if smudges == tolerance {
			return 0, c
		}
	}

	return -1, -1
}

func parseInput(data []string) []Pattern {
	patterns := make([]Pattern, len(data))
	for i, pattern := range data {
		patterns[i] = Pattern(strings.Split(pattern, "\n"))
	}
	return patterns
}

func part1(patterns []Pattern) int {
	res := 0
	for _, pattern := range patterns {
		r, c := pattern.FindReflection(0)
		if r == -1 {
			panic(fmt.Sprintf("Reflection not found for pattern:\n%s", strings.Join(pattern, "\n")))
		}
		res += 100*r + c
	}
	return res
}

func part2(patterns []Pattern) int {
	res := 0
	for _, pattern := range patterns {
		r, c := pattern.FindReflection(1)
		if r == -1 {
			panic(fmt.Sprintf("Smudged reflection not found for pattern:\n%s", strings.Join(pattern, "\n")))
		}
		res += 100*r + c
	}
	return res
}
