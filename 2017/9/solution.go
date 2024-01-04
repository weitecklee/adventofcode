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
	input := string(data)
	part1, part2 := solve(input)
	fmt.Println(part1)
	fmt.Println(part2)
}

func cleanGarbage(input string) (string, int) {
	var cleaned strings.Builder
	i := 0
	count := 0
	for i < len(input) {
		if string(input[i]) == "<" {
			for string(input[i]) != ">" {
				if string(input[i]) == "!" {
					i++
					count--
				}
				count++
				i++
			}
			count--
		} else {
			cleaned.WriteByte(input[i])
		}
		i++
	}
	return cleaned.String(), count
}

func solve(input string) (int, int) {
	cleaned, garbageCount := cleanGarbage(input)
	totalScore := 0
	currScore := 0
	for _, c := range cleaned {
		if string(c) == "{" {
			currScore++
		} else if string(c) == "}" {
			totalScore += currScore
			currScore--
		}
	}
	return totalScore, garbageCount
}
