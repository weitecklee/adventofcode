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
	input := string(data)
	fmt.Println(part1(parseInput(input)))
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

func part1(input []int) int {
	n := len(input)
	history := map[string]bool{}
	history[recordHistory(input)] = true
	steps := 0
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
		if history[hist] {
			break
		}
		history[hist] = true
	}
	return steps
}
