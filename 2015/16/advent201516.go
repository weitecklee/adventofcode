package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n")
	fmt.Println(part1(input))
}

func part1(input []string) int {
	re := regexp.MustCompile(`[a-z]+: \d+`)
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
	matches := re.FindAllString(tickerTape, -1)
	auntSue := map[string]bool{}
	for _, match := range matches {
		auntSue[match] = true
	}
loop:
	for i, line := range input {
		matches := re.FindAllString(line, -1)
		for _, match := range matches {
			if !auntSue[match] {
				continue loop
			}
		}
		return i + 1
	}
	return -1
}
