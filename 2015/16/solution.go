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
	auntSue, candidates := parseInput(input)
	fmt.Println(part1(auntSue, candidates))
	fmt.Println(part2(auntSue, candidates))
}

func parseInput(input []string) (map[string]int, []map[string]int) {
	re := regexp.MustCompile(`([a-z]+): (\d+)`)
	tickerTape := `children: 3
	cats: 7
	samoyeds: 2
	pomeranians: 3
	akitas: 0
	vizslas: 0
	goldfish: 5
	trees: 3
	cars: 2
	perfumes: 1`
	matches := re.FindAllStringSubmatch(tickerTape, -1)
	auntSue := map[string]int{}
	for _, match := range matches {
		n, _ := strconv.Atoi(match[2])
		auntSue[match[1]] = n
	}
	candidates := []map[string]int{}
	for _, line := range input {
		matches := re.FindAllStringSubmatch(line, -1)
		candidate := map[string]int{}
		for _, match := range matches {
			n, _ := strconv.Atoi(match[2])
			candidate[match[1]] = n
		}
		candidates = append(candidates, candidate)
	}
	return auntSue, candidates
}

func part1(auntSue map[string]int, candidates []map[string]int) int {
loop:
	for i, candidate := range candidates {
		for trait, n := range candidate {
			if auntSue[trait] != n {
				continue loop
			}
		}
		return i + 1
	}
	return -1
}

func part2(auntSue map[string]int, candidates []map[string]int) int {
loop:
	for i, candidate := range candidates {
		for trait, n := range candidate {
			if trait == "cats" || trait == "trees" {
				if auntSue[trait] >= n {
					continue loop
				}
			} else if trait == "pomeranians" || trait == "goldfish" {
				if auntSue[trait] <= n {
					continue loop
				}

			} else if auntSue[trait] != n {
				continue loop
			}
		}
		return i + 1
	}
	return -1
}
