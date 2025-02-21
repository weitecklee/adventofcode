package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var directions = [8][2]int{
	{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {-1, -1}, {-1, 1}, {1, 1}, {1, -1},
}

func main() {
	_, filename, _, _ := runtime.Caller(0)
	dirname := filepath.Dir(filename)
	inputFilePath := filepath.Join(dirname, "input.txt")
	data, err := os.ReadFile(inputFilePath)
	if err != nil {
		panic(err)
	}
	puzzleInput := strings.Split((string(data)), "\n")
	fmt.Println(part1(puzzleInput))
	fmt.Println(part2(puzzleInput))
}

func stringsInAllDirections(puzzleInput []string, r, c int) []string {
	res := make([]string, 0, 8)
	var sb strings.Builder
	for _, dir := range directions {
		sb.Reset()
		r2 := r
		c2 := c
		for i := 0; i < 4 && r2 >= 0 && c2 >= 0 && r2 < len(puzzleInput) && c2 < len(puzzleInput[0]); i++ {
			sb.WriteByte(puzzleInput[r2][c2])
			r2 += dir[0]
			c2 += dir[1]
		}
		res = append(res, sb.String())
	}
	return res
}

func part1(puzzleInput []string) int {
	res := 0
	for r := range puzzleInput {
		for c, chr := range puzzleInput[r] {
			if chr == 'X' {
				words := stringsInAllDirections(puzzleInput, r, c)
				for _, word := range words {
					if word == "XMAS" {
						res++
					}
				}
			}
		}
	}
	return res
}

func part2(puzzleInput []string) int {
	res := 0
	masCombos := []string{"MMSS", "MSSM", "SMMS", "SSMM"}
	matches := make(map[string]struct{})
	for _, combo := range masCombos {
		matches[combo] = struct{}{}
	}
	var sb strings.Builder
	for r := 1; r < len(puzzleInput)-1; r++ {
		for c := 1; c < len(puzzleInput[0])-1; c++ {
			if puzzleInput[r][c] == 'A' {
				sb.Reset()
				for i := 4; i < 8; i++ {
					sb.WriteByte(puzzleInput[r+directions[i][0]][c+directions[i][1]])
				}
				if _, ok := matches[sb.String()]; ok {
					res++
				}
			}
		}
	}
	return res
}
