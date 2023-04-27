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
	fmt.Println(part1(parseInput(input)))
}

type Instruction struct {
	category string
	nums     []int
}

func parseInput(input []string) (int, []Instruction) {
	instructions := []Instruction{}
	firstLine := strings.Split(input[0], " ")
	ip, _ := strconv.Atoi(firstLine[1])
	for _, line := range input[1:] {
		parts := strings.Split(line, " ")
		tmp := Instruction{}
		tmp.category = parts[0]
		for _, part := range parts[1:] {
			n, _ := strconv.Atoi(part)
			tmp.nums = append(tmp.nums, n)
		}
		instructions = append(instructions, tmp)
	}
	return ip, instructions
}

func execute(instruction Instruction, register *[6]int) {
	switch instruction.category {
	case "addr":
		register[instruction.nums[2]] = register[instruction.nums[0]] + register[instruction.nums[1]]
	case "addi":
		register[instruction.nums[2]] = register[instruction.nums[0]] + instruction.nums[1]
	case "mulr":
		register[instruction.nums[2]] = register[instruction.nums[0]] * register[instruction.nums[1]]
	case "muli":
		register[instruction.nums[2]] = register[instruction.nums[0]] * instruction.nums[1]
	case "banr":
		register[instruction.nums[2]] = register[instruction.nums[0]] & register[instruction.nums[1]]
	case "bani":
		register[instruction.nums[2]] = register[instruction.nums[0]] & instruction.nums[1]
	case "borr":
		register[instruction.nums[2]] = register[instruction.nums[0]] | register[instruction.nums[1]]
	case "bori":
		register[instruction.nums[2]] = register[instruction.nums[0]] | instruction.nums[1]
	case "setr":
		register[instruction.nums[2]] = register[instruction.nums[0]]
	case "seti":
		register[instruction.nums[2]] = instruction.nums[0]
	case "gtir":
		if instruction.nums[0] > register[instruction.nums[1]] {
			register[instruction.nums[2]] = 1
		} else {
			register[instruction.nums[2]] = 0
		}
	case "gtri":
		if register[instruction.nums[0]] > instruction.nums[1] {
			register[instruction.nums[2]] = 1
		} else {
			register[instruction.nums[2]] = 0
		}
	case "gtrr":
		if register[instruction.nums[0]] > register[instruction.nums[1]] {
			register[instruction.nums[2]] = 1
		} else {
			register[instruction.nums[2]] = 0
		}
	case "eqir":
		if instruction.nums[0] == register[instruction.nums[1]] {
			register[instruction.nums[2]] = 1
		} else {
			register[instruction.nums[2]] = 0
		}
	case "eqri":
		if register[instruction.nums[0]] == instruction.nums[1] {
			register[instruction.nums[2]] = 1
		} else {
			register[instruction.nums[2]] = 0
		}
	case "eqrr":
		if register[instruction.nums[0]] == register[instruction.nums[1]] {
			register[instruction.nums[2]] = 1
		} else {
			register[instruction.nums[2]] = 0
		}
	}
}

func part1(ip int, instructions []Instruction) int {
	register := [6]int{}
	for {
		if register[ip] > len(instructions)-1 {
			break
		}
		execute(instructions[register[ip]], &register)
		register[ip]++
	}
	return register[0]
}
