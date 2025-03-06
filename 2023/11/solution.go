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
	puzzleInput := strings.Split(string(data), "\n")
	fmt.Println(solve(puzzleInput, 2))
	fmt.Println(solve(puzzleInput, 1000000))
}

func solve(puzzleInput []string, expansionFactor int) int {
	var galaxies [][2]int
	var rowsToExpand, colsToExpand []int
	var expandRow, expandCol bool
	for r, row := range puzzleInput {
		expandRow = true
		for c, ch := range row {
			if ch == '#' {
				galaxies = append(galaxies, [2]int{r, c})
				expandRow = false
			}
		}
		if expandRow {
			rowsToExpand = append(rowsToExpand, r)
		}
	}
	for c := range puzzleInput[0] {
		expandCol = true
		for r := range puzzleInput {
			if puzzleInput[r][c] == '#' {
				expandCol = false
				break
			}
		}
		if expandCol {
			colsToExpand = append(colsToExpand, c)
		}
	}
	res := 0
	for i, galaxy1 := range galaxies {
		for _, galaxy2 := range galaxies[i+1:] {
			d := utils.AbsInt(galaxy1[0]-galaxy2[0]) + utils.AbsInt(galaxy1[1]-galaxy2[1])
			for _, r := range rowsToExpand {
				if (r > galaxy1[0] && r < galaxy2[0]) || (r < galaxy1[0] && r > galaxy2[0]) {
					d += expansionFactor - 1
				}
			}
			for _, c := range colsToExpand {
				if (c > galaxy1[1] && c < galaxy2[1]) || (c < galaxy1[1] && c > galaxy2[1]) {
					d += expansionFactor - 1
				}
			}
			res += d
		}
	}
	return res
}
