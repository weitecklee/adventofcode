package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	row, col := parseInput(string(data))
	fmt.Println(part1(row, col))
}

func parseInput(input string) (int, int) {
	re := regexp.MustCompile(`\d+`)
	parts := re.FindAllString(input, -1)
	row, _ := strconv.Atoi(parts[0])
	col, _ := strconv.Atoi(parts[1])
	return row, col
}

func part1(row int, col int) int {
	pyramidRow := row + col - 1
	n := (pyramidRow-1)*pyramidRow/2 + col
	num := 20151125
	for i := 1; i < n; i++ {
		num = (num * 252533) % 33554393
	}
	return num
}
