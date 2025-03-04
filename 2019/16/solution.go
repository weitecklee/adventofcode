package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
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
	puzzleInput := parseInput(string(data))
	fmt.Println(part1(puzzleInput))
}

func parseInput(s string) []int {
	nums := make([]int, 0, len(s))
	for _, ch := range s {
		if n, err := strconv.Atoi(string(ch)); err != nil {
			panic(err)
		} else {
			nums = append(nums, n)
		}
	}
	return nums
}

func calcPatterns(n int) [][]int {
	patterns := make([][]int, n)
	for i := range n {
		pattern := make([]int, n)
		for k := 0; 4*k*(i+1)+i < n; k++ {
			for j := 4*k*(i+1) + i; j < (4*k+1)*(i+1)+i && j < n; j++ {
				pattern[j] = 1
			}
			for j := (4*k+2)*(i+1) + i; j < (4*k+3)*(i+1)+i && j < n; j++ {
				pattern[j] = -1
			}
		}
		patterns[i] = pattern
	}
	return patterns
}

func fft(nums []int, patterns [][]int) []int {
	res := make([]int, len(nums))
	var curr int
	for i := range nums {
		curr = 0
		for j, num := range nums {
			curr += patterns[i][j] * num
		}
		res[i] = utils.AbsInt(curr % 10)
	}
	return res
}

func part1(puzzleInput []int) string {
	patterns := calcPatterns(len(puzzleInput))
	for range 100 {
		puzzleInput = fft(puzzleInput, patterns)
	}
	var sb strings.Builder
	for i := range 8 {
		sb.WriteByte(byte(puzzleInput[i]) + '0')
	}
	return sb.String()
}
