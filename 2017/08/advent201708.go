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
	input := strings.Split(string(data), "\n")
	fmt.Println(part1(input))
}

func part1(input []string) int {
	regMap := map[string]int{}
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
		}
	}
	max := 0
	for _, v := range regMap {
		if v > max {
			max = v
		}
	}
	return max
}
