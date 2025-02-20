package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
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
	puzzleInput := parseInput(strings.Split(string(data), "\n"))
	fmt.Println(part1(puzzleInput))
	fmt.Println(part2(puzzleInput))
}

func parseInput(data []string) [][]int {
	numbers := make([][]int, len(data))
	numRegex := regexp.MustCompile(`\d+`)
	for i, s := range data {
		matches := numRegex.FindAllString(s, -1)
		nums := make([]int, len(matches))
		for j, match := range matches {
			if n, err := strconv.Atoi(match); err != nil {
				panic(err)
			} else {
				nums[j] = n
			}
		}
		numbers[i] = nums
	}
	return numbers
}

func isSafe(nums []int) bool {
	increasing := nums[1] > nums[0]
	for i := 1; i < len(nums); i++ {
		diff := utils.AbsInt((nums[i] - nums[i-1]))
		if diff < 1 || diff > 3 {
			return false
		}
		if (nums[i] > nums[i-1]) != increasing {
			return false
		}
	}
	return true
}

func isSafeWithTolerance(nums []int) bool {
	if isSafe(nums) {
		return true
	}
	for i := range nums {
		nums2 := make([]int, 0, len(nums)-1)
		nums2 = append(nums2, (nums)[:i]...)
		nums2 = append(nums2, (nums)[i+1:]...)
		if isSafe(nums2) {
			return true
		}
	}
	return false
}

func part1(puzzleInput [][]int) int {
	res := 0
	for _, nums := range puzzleInput {
		if isSafe(nums) {
			res++
		}
	}
	return res
}

func part2(puzzleInput [][]int) int {
	res := 0
	for _, nums := range puzzleInput {
		if isSafeWithTolerance(nums) {
			res++
		}
	}
	return res
}
