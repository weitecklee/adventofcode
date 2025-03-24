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

var directions = [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func parseInput(data []string) [][]int {
	puzzleInput := make([][]int, len(data))
	for i, line := range data {
		nums := make([]int, len(line))
		for j, s := range line {
			nums[j] = int(s - '0')
		}
		puzzleInput[i] = nums
	}
	return puzzleInput
}

func part1(puzzleInput [][]int) int {
	visibleTrees := make(map[[2]int]struct{})
	var curr int
	for r, row := range puzzleInput {
		visibleTrees[[2]int{r, 0}] = struct{}{}
		curr = puzzleInput[r][0]
		for c := 1; c < len(row); c++ {
			if row[c] > curr {
				visibleTrees[[2]int{r, c}] = struct{}{}
				curr = row[c]
			}
			if curr == 9 {
				break
			}
		}
		visibleTrees[[2]int{r, len(row) - 1}] = struct{}{}
		curr = puzzleInput[r][len(row)-1]
		for c := len(row) - 2; c >= 0; c-- {
			if row[c] > curr {
				visibleTrees[[2]int{r, c}] = struct{}{}
				curr = row[c]
			}
			if curr == 9 {
				break
			}
		}
	}
	for c := range puzzleInput[0] {
		visibleTrees[[2]int{0, c}] = struct{}{}
		curr = puzzleInput[0][c]
		for r := 1; r < len(puzzleInput); r++ {
			if puzzleInput[r][c] > curr {
				visibleTrees[[2]int{r, c}] = struct{}{}
				curr = puzzleInput[r][c]
			}
			if curr == 9 {
				break
			}
		}
		visibleTrees[[2]int{len(puzzleInput) - 1, c}] = struct{}{}
		curr = puzzleInput[len(puzzleInput)-1][c]
		for r := len(puzzleInput) - 2; r >= 0; r-- {
			if puzzleInput[r][c] > curr {
				visibleTrees[[2]int{r, c}] = struct{}{}
				curr = puzzleInput[r][c]
			}
			if curr == 9 {
				break
			}
		}
	}
	return len(visibleTrees)
}

func part2(puzzleInput [][]int) int {
	res := 0
	for r, row := range puzzleInput {
		for c, tree := range row {
			curr := 1
			for _, d := range directions {
				tmp := 0
				r2, c2 := r+d[0], c+d[1]
				for r2 >= 0 && c2 >= 0 && r2 < len(puzzleInput) && c2 < len(row) {
					tmp++
					if puzzleInput[r2][c2] >= tree {
						break
					}
					r2 += d[0]
					c2 += d[1]
				}
				curr *= tmp
			}
			if curr > res {
				res = curr
			}
		}
	}
	return res
}
