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
	fmt.Println(part2(puzzleInput))
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

func generateSecretNumbers(num, n int) []int {
	secretNumbers := make([]int, n)
	for i := 0; i < n; i++ {
		num = prune(mix(num, num*64))
		num = prune(mix(num, num/32))
		num = prune(mix(num, num*2048))
		secretNumbers[i] = num
	}
	return secretNumbers
}

func part1(puzzleInput []int) int {
	res := 0
	for _, n := range puzzleInput {
		res += generateSecretNumbers(n, 2000)[1999]
	}
	return res
}

func part2(puzzleInput []int) int {
	sequenceScoresAll := make(map[[4]int]int)
	var sequence [4]int
	for _, n := range puzzleInput {
		secretNumbers := generateSecretNumbers(n, 2000)
		prices := make([]int, len(secretNumbers))
		for i, secretNum := range secretNumbers {
			prices[i] = secretNum % 10
		}
		sequenceScores := make(map[[4]int]int)
		var currSequence []int
		for i := 1; i < 5; i++ {
			currSequence = append(currSequence, prices[i]-prices[i-1])
		}
		copy(sequence[:], currSequence)
		sequenceScores[sequence] = prices[4]
		for i := 5; i < len(secretNumbers); i++ {
			currSequence = currSequence[1:]
			currSequence = append(currSequence, prices[i]-prices[i-1])
			copy(sequence[:], currSequence)
			if _, ok := sequenceScores[sequence]; !ok {
				sequenceScores[sequence] = prices[i]
			}
		}
		for seq, score := range sequenceScores {
			sequenceScoresAll[seq] += score
		}
	}
	res := 0
	for _, score := range sequenceScoresAll {
		if score > res {
			res = score
		}
	}
	return res
}
