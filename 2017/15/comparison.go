package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	parsedInput := parseInput(strings.Split(string(data), "\n"))
	fmt.Println(part1(parsedInput))
	fmt.Println(part2(parsedInput))
	fmt.Println(part1v2(parsedInput))
	fmt.Println(part2v2(parsedInput))
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	fmt.Printf("%v: %v\n", msg, time.Since(start))
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
	defer duration(track("part1"))
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

func part1v2(input []int) int {
	defer duration(track("part1 v2"))
	count := 0
	factorA := 16807
	factorB := 48271
	genA := input[0]
	genB := input[1]
	for i := 0; i < 40000000; i++ {
		genA = (genA * factorA) % 2147483647
		genB = (genB * factorB) % 2147483647
		if genA&0xFFFF == genB&0xFFFF {
			count++
		}
	}
	return count
}

func generator(val int, f int, m int) int {
	for {
		val = (val * f) % 2147483647
		if val%m == 0 {
			break
		}
	}
	return val
}

func part2(input []int) int {
	defer duration(track("part2"))
	count := 0
	factorA := 16807
	factorB := 48271
	genA := input[0]
	genB := input[1]
	for i := 0; i < 5000000; i++ {
		genA = generator(genA, factorA, 4)
		genB = generator(genB, factorB, 8)
		if genA%65536 == genB%65536 {
			count++
		}
	}
	return count
}

func part2v2(input []int) int {
	defer duration(track("part2 v2"))
	count := 0
	factorA := 16807
	factorB := 48271
	genA := input[0]
	genB := input[1]
	for i := 0; i < 5000000; i++ {
		genA = generator(genA, factorA, 4)
		genB = generator(genB, factorB, 8)
		if genA&0xFFFF == genB&0xFFFF {
			count++
		}
	}
	return count
}
