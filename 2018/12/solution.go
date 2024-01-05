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
	rules, startState, potLen := parseInput(input)
	fmt.Println(part1(rules, startState, potLen))
	fmt.Println(part2(rules, startState, potLen))
}

func parseInput(input []string) (map[string]string, map[int]string, int) {
	rules := map[string]string{}
	parts := strings.Split(input[0], " ")
	pots := map[int]string{}
	for i, c := range parts[2] {
		if c == '#' {
			pots[i] = "#"
		}
	}
	potLen := len(parts[2])
	for _, line := range input[2:] {
		parts := strings.Split(line, " ")
		rules[parts[0]] = parts[2]
	}
	return rules, pots, potLen
}

func newGeneration(rules *map[string]string, pots *map[int]string, potRange *[2]int) map[int]string {
	pots2 := map[int]string{}
	(*potRange)[0]--
	(*potRange)[1]++
	for i := (*potRange)[0]; i <= (*potRange)[1]; i++ {
		potState := ""
		for j := -2; j <= 2; j++ {
			if (*pots)[i+j] == "" || (*pots)[i+j] == "." {
				potState += "."
			} else {
				potState += "#"
			}
		}
		pots2[i] = (*rules)[potState]
	}
	return pots2
}

func countPots(pots *map[int]string) int {
	count := 0
	for i, s := range *pots {
		if s == "#" {
			count += i
		}
	}
	return count
}

func part1(rules map[string]string, pots map[int]string, potLen int) int {
	potRange := [2]int{0, potLen - 1}
	for generations := 0; generations < 20; generations++ {
		pots = newGeneration(&rules, &pots, &potRange)
	}
	return countPots(&pots)
}

func part2(rules map[string]string, pots map[int]string, potLen int) int {
	/*
		Analysis of the growth pattern of `count` reveals that it reaches constant growth
		after a certain number of generations. (For my input, it keeps growing by 8 after
		153 generations)
		Iterate until it reaches this constant growth and calculate the remaining growth
		necessary (Just comparing the current difference and the previous difference is
		enough, for mine at least)
	*/
	potRange := [2]int{0, potLen - 1}
	count := countPots(&pots)
	diff := 0
	for generations := 1; generations < 1000; generations++ {
		pots = newGeneration(&rules, &pots, &potRange)
		count2 := countPots(&pots)
		diff2 := count2 - count
		if diff2 == diff {
			return (50000000000-generations)*diff + count2
		}
		diff = diff2
		count = count2
	}
	return -1
}
