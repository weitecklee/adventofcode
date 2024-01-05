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
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		parsedInput = append(parsedInput, n)
	}
	invalidNum := part1(parsedInput)
	fmt.Println(invalidNum)
	fmt.Println(part2(parsedInput, invalidNum))
}

func part1(input []int) int {
	preambleLen := 25
	preamble := map[int]bool{}
	for i := 0; i < preambleLen; i++ {
		preamble[input[i]] = true
	}
	for i := preambleLen; i < len(input); i++ {
		valid := false
		for n := range preamble {
			if preamble[input[i]-n] && n != input[i]-n {
				valid = true
				break
			}
		}
		if !valid {
			return input[i]
		}
		preamble[input[i]] = true
		delete(preamble, input[i-preambleLen])
	}
	return -1
}

func part2(input []int, invalidNum int) int {
	start := 0
	finish := 1
	sum := input[0] + input[1]
	for sum != invalidNum {
		finish++
		sum += input[finish]
		for sum > invalidNum {
			sum -= input[start]
			start++
		}
	}
	contiguousRange := input[start : finish+1]
	sort.Ints(contiguousRange)
	return contiguousRange[0] + contiguousRange[len(contiguousRange)-1]
}
