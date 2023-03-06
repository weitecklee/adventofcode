package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n")
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func findAnswers(data []string) map[rune]int {
	answers := map[rune]int{}
	for _, line := range data {
		for _, char := range line {
			answers[char]++
		}
	}
	return answers
}

func part1(input []string) int {
	row := 0
	sum := 0
	for row < len(input) {
		data := []string{}
		for row < len(input) && len(input[row]) > 0 {
			data = append(data, input[row])
			row++
		}
		sum += len(findAnswers(data))
		row++
	}
	return sum
}

func part2(input []string) int {
	row := 0
	sum := 0
	for row < len(input) {
		data := []string{}
		for row < len(input) && len(input[row]) > 0 {
			data = append(data, input[row])
			row++
		}
		answers := findAnswers(data)
		for _, n := range answers {
			if n == len(data) {
				sum++
			}
		}
		row++
	}
	return sum
}
