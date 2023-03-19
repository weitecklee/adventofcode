package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n")
	parsedInput := parseInput(input)
	fmt.Println(part1(parsedInput))
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
