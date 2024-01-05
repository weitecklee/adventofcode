package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
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
	parsedInput := []int{}
	for _, s := range input {
		n, _ := strconv.Atoi(s)
		parsedInput = append(parsedInput, n)
	}
	parsedInput = append(parsedInput, 0)
	sort.Ints(parsedInput)
	parsedInput = append(parsedInput, parsedInput[len(parsedInput)-1]+3)
	fmt.Println(part1(parsedInput))
	fmt.Println(part2(parsedInput))
}

func part1(input []int) int {
	diff1 := 0
	diff3 := 0
	for i := 1; i < len(input); i++ {
		diff := input[i] - input[i-1]
		if diff == 1 {
			diff1++
		} else if diff == 3 {
			diff3++
		}
	}
	return diff1 * diff3
}

func part2(input []int) int {
	branches := map[int]int{}
	branches[input[len(input)-1]] = 1
	for i := len(input) - 2; i >= 0; i-- {
		branches[input[i]] = branches[input[i]+1] + branches[input[i]+2] + branches[input[i]+3]
	}
	return branches[0]
}
