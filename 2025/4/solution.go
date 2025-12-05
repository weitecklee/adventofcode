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
	puzzleInput := parseInput(strings.Split(string(data), "\n"))
	fmt.Println(part1(puzzleInput))
	fmt.Println(part2(puzzleInput))
}

func parseInput(data []string) map[[2]int]int {
	papers := make(map[[2]int]int)
	for i, row := range data {
		for j, ch := range row {
			if ch == '@' {
				papers[[2]int{i, j}] = 0
			}
		}
	}
	for coord := range papers {
		a, b := coord[0], coord[1]
		for r := a - 1; r <= a+1; r++ {
			for c := b - 1; c <= b+1; c++ {
				if r == a && c == b {
					continue
				}
				tmp := [2]int{r, c}
				if _, ok := papers[tmp]; ok {
					papers[tmp] += 1
				}
			}
		}
	}
	return papers
}

func removePaper(papers map[[2]int]int) int {
	removed := 0
	for _, ct := range papers {
		if ct < 4 {
			removed += 1
		}
	}
	return removed
}

func removePaperContinuous(papers map[[2]int]int) int {
	res := 0
	for {
		removed := 0
		for coord, ct := range papers {
			if ct < 4 {
				removed += 1
				delete(papers, coord)
				a, b := coord[0], coord[1]
				for r := a - 1; r <= a+1; r++ {
					for c := b - 1; c <= b+1; c++ {
						if r == a && c == b {
							continue
						}
						tmp := [2]int{r, c}
						if _, ok := papers[tmp]; ok {
							papers[tmp] -= 1
						}
					}
				}
			}
		}
		if removed == 0 {
			break
		}
		res += removed
	}
	return res
}

func part1(puzzleInput map[[2]int]int) int {
	return removePaper(puzzleInput)
}

func part2(puzzleInput map[[2]int]int) int {
	return removePaperContinuous(puzzleInput)
}
