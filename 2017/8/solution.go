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
	input := strings.Split(string(data), "\n")
	part1, part2 := solve(input)
	fmt.Println(part1)
	fmt.Println(part2)
}

func solve(input []string) (int, int) {
	regMap := map[string]int{}
	runningMax := 0
	for _, line := range input {
		parts := strings.Split(line, " ")
		reg := parts[0]
		incN, _ := strconv.Atoi(parts[2])
		if parts[1] == "dec" {
			incN *= -1
		}
		compN, _ := strconv.Atoi(parts[6])
		check := false
		regIf := parts[4]
		switch parts[5] {
		case "!=":
			check = regMap[regIf] != compN
		case "==":
			check = regMap[regIf] == compN
		case ">":
			check = regMap[regIf] > compN
		case ">=":
			check = regMap[regIf] >= compN
		case "<":
			check = regMap[regIf] < compN
		case "<=":
			check = regMap[regIf] <= compN
		default:
			panic(parts[5])
		}
		if check {
			regMap[reg] += incN
			if regMap[reg] > runningMax {
				runningMax = regMap[reg]
			}
		}
	}
	max := 0
	for _, v := range regMap {
		if v > max {
			max = v
		}
	}
	return max, runningMax
}
