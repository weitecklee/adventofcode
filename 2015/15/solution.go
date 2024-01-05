package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
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
	ingredients := parseInput(input)
	fmt.Println(part1(ingredients))
	fmt.Println(part2(ingredients))
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

func recur(ingredients *[][]int, maxTotal *int, n int, lenIngredients int, amounts *[]int, runningAmount int, part2 bool) {
	if n == lenIngredients-1 {
		(*amounts)[n] = 100 - runningAmount
		if part2 {
			calories := 0
			for j := 0; j < lenIngredients; j++ {
				calories += (*ingredients)[j][4] * (*amounts)[j]
			}
			if calories != 500 {
				return
			}
		}
		total := 1
		for i := 0; i < 4; i++ {
			curr := 0
			for j := 0; j < lenIngredients; j++ {
				curr += (*ingredients)[j][i] * (*amounts)[j]
			}
			if curr > 0 {
				total *= curr
			}
		}
		if total > *maxTotal {
			*maxTotal = total
		}
		return
	}
	for i := 0; i+runningAmount <= 100; i++ {
		(*amounts)[n] = i
		recur(ingredients, maxTotal, n+1, lenIngredients, amounts, runningAmount+i, part2)
	}
}

func part1(ingredients [][]int) int {
	maxTotal := 0
	amounts := make([]int, len(ingredients))
	recur(&ingredients, &maxTotal, 0, len(ingredients), &amounts, 0, false)
	return maxTotal
}

func part2(ingredients [][]int) int {
	maxTotal := 0
	amounts := make([]int, len(ingredients))
	recur(&ingredients, &maxTotal, 0, len(ingredients), &amounts, 0, true)
	return maxTotal
}
