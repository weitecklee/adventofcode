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
	input := string(data)
	part1, part2 := solve(parseInput(input))
	fmt.Println(part1)
	fmt.Println(part2)
}

func parseInput(input string) []int {
	re := regexp.MustCompile(`\d+`)
	parsed := []int{}
	matches := re.FindAllString(input, -1)
	for _, s := range matches {
		n, _ := strconv.Atoi(s)
		parsed = append(parsed, n)
	}
	return parsed
}

func recordHistory(input []int) string {
	str := []string{}
	for _, n := range input {
		s := strconv.Itoa(n)
		str = append(str, s)
	}
	return strings.Join(str, ",")
}

func solve(input []int) (int, int) {
	n := len(input)
	history := map[string]int{}
	history[recordHistory(input)] = 1
	steps := 1
	for {
		steps++
		most := input[0]
		iMost := 0
		for i := 1; i < n; i++ {
			if input[i] > most {
				most = input[i]
				iMost = i
			}
		}
		input[iMost] = 0
		for most > 0 {
			iMost++
			if iMost >= n {
				iMost = 0
			}
			input[iMost]++
			most--
		}
		hist := recordHistory(input)
		if history[hist] > 0 {
			return steps - 1, steps - history[hist]
		}
		history[hist] = steps
	}
}
