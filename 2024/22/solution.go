package main

import (
	"fmt"
	"os"
	"path/filepath"
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
	puzzleInput := parseInput(strings.Split(string(data), "\n"))
	fmt.Println(part1(puzzleInput))
}

func parseInput(data []string) []int {
	puzzleInput := make([]int, len(data))
	for i, s := range data {
		if n, err := strconv.Atoi(s); err != nil {
			panic(err)
		} else {
			puzzleInput[i] = n
		}
	}
	return puzzleInput
}

func mix(secretNum, num int) int {
	return secretNum ^ num
}

func prune(secretNum int) int {
	return secretNum % 16777216
}

func secretNumberN(num, n int) int {
	secretNumber := num
	for i := 0; i < n; i++ {
		secretNumber = prune(mix(secretNumber, secretNumber*64))
		secretNumber = prune(mix(secretNumber, secretNumber/32))
		secretNumber = prune(mix(secretNumber, secretNumber*2048))
	}
	return secretNumber
}

func part1(puzzleInput []int) int {
	res := 0
	for _, n := range puzzleInput {
		res += secretNumberN(n, 2000)
	}
	return res
}
