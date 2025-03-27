package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/weitecklee/adventofcode/utils"
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
	fmt.Println(solve(puzzleInput))
}

func parseInput(data []string) [][]int {
	puzzleInput := make([][]int, len(data))
	for i, line := range data {
		puzzleInput[i] = utils.ExtractInts(line)
	}
	return puzzleInput
}

func solve(puzzleInput [][]int) (int, int) {
	areaMap := make(map[[2]int]map[int]struct{})
	claimMap := make(map[int]bool, len(puzzleInput))
	for _, claim := range puzzleInput {
		id := claim[0]
		c := claim[1]
		r := claim[2]
		w := claim[3]
		h := claim[4]
		claimMap[id] = true
		for dr := range h {
			for dc := range w {
				pos := [2]int{r + dr, c + dc}
				if _, exists := areaMap[pos]; !exists {
					areaMap[pos] = make(map[int]struct{})
				}
				curr := areaMap[pos]
				if len(curr) > 0 {
					for claimId := range curr {
						claimMap[claimId] = false
					}
					claimMap[id] = false
				}
				curr[id] = struct{}{}
			}
		}
	}

	part1 := 0
	for _, claims := range areaMap {
		if len(claims) >= 2 {
			part1++
		}
	}

	part2 := -1
	for id, noOverlap := range claimMap {
		if noOverlap {
			part2 = id
			break
		}
	}

	return part1, part2
}
