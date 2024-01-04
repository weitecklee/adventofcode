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
	instructions := parseInput(input)
	fmt.Println(part1(instructions))
	fmt.Println(part2(instructions))
}

type Instruction struct {
	instruct string
	coords   [4]int
}

func parseInput(input []string) []Instruction {
	instructions := []Instruction{}
	re := regexp.MustCompile(`\d+`)
	for _, line := range input {
		matches := re.FindAllString(line, -1)
		parts := strings.Split(line, " ")
		coords := [4]int{}
		for i, match := range matches {
			n, _ := strconv.Atoi(match)
			coords[i] = n
		}
		if coords[2] < coords[0] || coords[3] < coords[1] {
			panic(coords)
		}
		instruct := parts[0]
		if instruct != "toggle" {
			instruct = parts[1]
		}
		instructions = append(instructions, Instruction{
			instruct: instruct,
			coords:   coords,
		})
	}
	return instructions
}

func part1(instructions []Instruction) int {
	lights := [1000][1000]bool{}
	for _, step := range instructions {
		switch step.instruct {
		case "on":
			for i := step.coords[0]; i <= step.coords[2]; i++ {
				for j := step.coords[1]; j <= step.coords[3]; j++ {
					lights[i][j] = true
				}
			}
		case "off":
			for i := step.coords[0]; i <= step.coords[2]; i++ {
				for j := step.coords[1]; j <= step.coords[3]; j++ {
					lights[i][j] = false
				}
			}
		case "toggle":
			for i := step.coords[0]; i <= step.coords[2]; i++ {
				for j := step.coords[1]; j <= step.coords[3]; j++ {
					lights[i][j] = !lights[i][j]
				}
			}
		}
	}
	count := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if lights[i][j] {
				count++
			}
		}
	}
	return count
}

func part2(instructions []Instruction) int {
	lights := [1000][1000]int{}
	for _, step := range instructions {
		switch step.instruct {
		case "on":
			for i := step.coords[0]; i <= step.coords[2]; i++ {
				for j := step.coords[1]; j <= step.coords[3]; j++ {
					lights[i][j]++
				}
			}
		case "off":
			for i := step.coords[0]; i <= step.coords[2]; i++ {
				for j := step.coords[1]; j <= step.coords[3]; j++ {
					if lights[i][j] > 0 {
						lights[i][j]--
					}
				}
			}
		case "toggle":
			for i := step.coords[0]; i <= step.coords[2]; i++ {
				for j := step.coords[1]; j <= step.coords[3]; j++ {
					lights[i][j] += 2
				}
			}
		}
	}
	total := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			total += lights[i][j]
		}
	}
	return total
}
