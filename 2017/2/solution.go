package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
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
	input := strings.Split(string(data), "\n")
	parsedInput := parseInput(input)
	fmt.Println(part1(parsedInput))
	fmt.Println(part2(parsedInput))
}

func parseInput(input []string) [][]int {
	res := [][]int{}
	re := regexp.MustCompile(`\d+`)
	for _, row := range input {
		nrow := []int{}
		matches := re.FindAllString(row, -1)
		for _, match := range matches {
			n, _ := strconv.Atoi(match)
			nrow = append(nrow, n)
		}
		res = append(res, nrow)
	}
	return res
}

func part1(input [][]int) int {
	res := 0
	for _, row := range input {
		max := row[0]
		min := row[0]
		for _, n := range row {
			if n > max {
				max = n
			}
			if n < min {
				min = n
			}
		}
		res += max - min
	}
	return res
}

func part2(input [][]int) int {
	res := 0
	for _, row := range input {
		found := false
		for i := 0; i < len(row); i++ {
			for j := i + 1; j < len(row); j++ {
				if row[i]%row[j] == 0 {
					res += row[i] / row[j]
					found = true
					break
				}
				if row[j]%row[i] == 0 {
					res += row[j] / row[i]
					found = true
					break
				}
			}
			if found {
				break
			}
		}
	}
	return res
}
