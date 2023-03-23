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

func parseInput(input []string) *[][]int {
	parsed := [][]int{}
	re := regexp.MustCompile(`\d+`)
	for _, line := range input {
		matches := re.FindAllString(line, -1)
		tmp := []int{}
		for _, match := range matches {
			n, _ := strconv.Atoi(match)
			tmp = append(tmp, n)
		}
		parsed = append(parsed, tmp)
	}
	return &parsed
}

func part1(input *[][]int) int {
	severity := 0
	for _, layer := range *input {
		if layer[0]%((layer[1]-1)*2) == 0 {
			severity += layer[0] * layer[1]
		}
	}
	return severity
}
