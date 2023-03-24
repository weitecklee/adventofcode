package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(part1(parseInput(strings.Split(string(data), "\n"))))
}

func parseInput(input []string) []int {
	parsed := []int{}
	for _, line := range input {
		s := strings.Split(line, " ")
		n, _ := strconv.Atoi(s[len(s)-1])
		parsed = append(parsed, n)
	}
	return parsed
}

func part1(input []int) int {
	count := 0
	factorA := 16807
	factorB := 48271
	genA := input[0]
	genB := input[1]
	for i := 0; i < 40000000; i++ {
		genA = (genA * factorA) % 2147483647
		genB = (genB * factorB) % 2147483647
		if genA%65536 == genB%65536 {
			count++
		}
	}
	return count
}
