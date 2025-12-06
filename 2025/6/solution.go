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
	numbers, ops := parseInput(strings.Split(string(data), "\n"))
	fmt.Println(part1(numbers, ops))
	fmt.Println(part2(strings.Split(string(data), "\n")))
}

func parseInput(data []string) ([][]int, string) {
	numbers := make([][]int, len(data)-1)
	for i := range len(data) - 1 {
		numbers[i] = utils.ExtractInts(data[i])
	}
	line := data[len(data)-1]
	spaceRegex := regexp.MustCompile(`\s+`)
	line = spaceRegex.ReplaceAllString(line, "")
	return numbers, line
}

func part1(numbers [][]int, ops string) int {
	res := 0
	for col := range len(numbers[0]) {
		var curr int
		if ops[col] == '+' {
			for row := range numbers {
				curr += numbers[row][col]
			}
		} else {
			curr = 1
			for row := range numbers {
				curr *= numbers[row][col]
			}
		}
		res += curr
	}
	return res
}

func part2(puzzleInput []string) int {
	res := 0
	ops := puzzleInput[len(puzzleInput)-1]
	spaceRegex := regexp.MustCompile(`\s+`)
	ops = spaceRegex.ReplaceAllString(ops, "")
	puzzleInput = puzzleInput[:len(puzzleInput)-1]
	rowLen := len(puzzleInput[0])
	for _, row := range puzzleInput {
		if len(row) > rowLen {
			rowLen = len(row)
		}
	}
	transposed := make([]string, rowLen)
	for i := range rowLen {
		var s strings.Builder
		for _, line := range puzzleInput {
			if i >= len(line) {
				s.WriteByte(' ')
			} else {
				s.WriteByte(line[i])
			}
		}
		transposed[i] = s.String()
	}
	var numbers []int
	var j int
	for _, s := range transposed {
		n, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			var tmp int
			if ops[j] == '+' {
				for _, n := range numbers {
					tmp += n
				}
			} else {
				tmp = 1
				for _, n := range numbers {
					tmp *= n
				}
			}
			res += tmp
			numbers = make([]int, 0)
			j += 1
		} else {
			numbers = append(numbers, n)
		}
	}
	var tmp int
	if ops[j] == '+' {
		for _, n := range numbers {
			tmp += n
		}
	} else {
		tmp = 1
		for _, n := range numbers {
			tmp *= n
		}
	}
	res += tmp
	return res
}
