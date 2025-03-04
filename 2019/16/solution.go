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
	fmt.Println(part2(puzzleInput))
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

/*
	Trick to Part 2 is that the offset is in the 2nd half of the list.
	There, the output digits can be calculated by starting from the
	end and working backwards, adding another digit each time.
	The pattern for the 2nd half of the list will always look like this:
	0 0 0 ... 1 1 1 1 1 1
	0 0 0 ... 0 1 1 1 1 1
	0 0 0 ... 0 0 1 1 1 1
	0 0 0 ... 0 0 0 1 1 1
	0 0 0 ... 0 0 0 0 1 1
	0 0 0 ... 0 0 0 0 0 1
	Therefore, you only need the digits after the offset. Everything before
	will not contribute anything to what we're looking for.
*/

func fft2(nums []int) []int {
	res := make([]int, len(nums))
	res[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		res[i] = (res[i+1] + nums[i]) % 10
	}
	return res
}

func part2(puzzleInput []int) string {
	offset := 0
	for i := range 7 {
		offset = offset*10 + puzzleInput[i]
	}
	extendedInput := make([]int, len(puzzleInput)*10000)
	for i := range 10000 {
		copy(extendedInput[i*len(puzzleInput):], puzzleInput)
	}
	extendedInput = extendedInput[offset:]
	for range 100 {
		extendedInput = fft2(extendedInput)
	}
	var sb strings.Builder
	for i := range 8 {
		sb.WriteByte(byte(extendedInput[i]) + '0')
	}
	return sb.String()
}
