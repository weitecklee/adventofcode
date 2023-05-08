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
	ingredients := parseInput(input)
	fmt.Println(part1(ingredients))
}

func parseInput(input []string) [][]int {
	ingredients := [][]int{}
	re := regexp.MustCompile(`-?\d+`)
	for _, line := range input {
		matches := re.FindAllString(line, -1)
		nums := []int{}
		for _, match := range matches {
			n, _ := strconv.Atoi(match)
			nums = append(nums, n)
		}
		ingredients = append(ingredients, nums)
	}
	return ingredients
}

func part1(ingredients [][]int) int {
	maxTotal := 0
	for i := 0; i <= 100; i++ {
		for j := 0; i+j <= 100; j++ {
			for k := 0; i+j+k <= 100; k++ {
				l := 100 - i - j - k
				total := 1
				for m := 0; m < 4; m++ {
					curr := ingredients[0][m]*i + ingredients[1][m]*j + ingredients[2][m]*k + ingredients[3][m]*l
					if curr > 0 {
						total *= curr
					}
				}
				if total > maxTotal {
					maxTotal = total
				}
			}
		}
	}
	return maxTotal
}
