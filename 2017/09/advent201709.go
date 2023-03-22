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
	fmt.Println(part1(input))
}

func cleanGarbage(input string) string {
	var cleaned strings.Builder
	i := 0
	for i < len(input) {
		if string(input[i]) == "<" {
			for string(input[i]) != ">" {
				if string(input[i]) == "!" {
					i++
				}
				i++
			}
		} else {
			cleaned.WriteByte(input[i])
		}
		i++
	}
	return cleaned.String()
}

func part1(input string) int {
	cleaned := cleanGarbage(input)
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
	return totalScore
}
