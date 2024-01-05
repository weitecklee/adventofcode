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

func part2(input *[][]int) int {
	delay := 0
	for {
		delay++
		caught := false
		for _, layer := range *input {
			if (layer[0]+delay)%((layer[1]-1)*2) == 0 {
				caught = true
				break
			}
		}
		if !caught {
			break
		}
	}
	return delay
}
