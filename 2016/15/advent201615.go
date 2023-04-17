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
	fmt.Println(part2(parsedInput))
}

func parseInput(input []string) [][]int {
	re := regexp.MustCompile(`\d+`)
	discs := [][]int{}
	for _, line := range input {
		disc := []int{}
		matches := re.FindAllString(line, -1)
		for _, match := range matches {
			n, _ := strconv.Atoi(match)
			disc = append(disc, n)
		}
		discs = append(discs, disc)
	}
	return discs
}

func part1(discs [][]int) int {
	pressTime := 0
	pass := false
	for !pass {
		pressTime++
		pass = true
		for _, disc := range discs {
			if (pressTime+disc[0]+disc[3])%disc[1] != 0 {
				pass = false
				break
			}
		}
	}
	return pressTime
}

func part2(discs [][]int) int {
	pressTime := 0
	pass := false
	discs = append(discs, []int{7, 11, 0, 0})
	for !pass {
		pressTime++
		pass = true
		for _, disc := range discs {
			if (pressTime+disc[0]+disc[3])%disc[1] != 0 {
				pass = false
				break
			}
		}
	}
	return pressTime
}
