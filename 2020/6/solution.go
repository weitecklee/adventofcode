package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
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
