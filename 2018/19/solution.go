package main

import (
	"fmt"
	"math"
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
	fmt.Println(part1(parseInput(input)))
	fmt.Println(part2(parseInput(input)))
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

/*
Main loop of process takes a number N and loops through pairs of numbers (1 <= i, j <= N)
and adds i to register 0 if i * j == N. Basically it is calculating the sum of divisors
of N very very inefficiently, especially if N = 10551425 as in my case.
*/

func part2(ip int, instructions []Instruction) int {
	register := [6]int{}
	register[0] = 1
	// run through program until it starts main loop, which is when it goes to line 1 (0-index)
	for register[ip] != 1 {
		if register[ip] > len(instructions)-1 {
			break
		}
		execute(instructions[register[ip]], &register)
		register[ip]++
	}
	// number N is in register[3] ** FOR MY INPUT, POSSIBLY DIFFERENT FOR OTHERS **
	sum := 0
	N := register[3]
	k := int(math.Sqrt(float64(N)))
	for i := 1; i <= k; i++ { // technically double counts if N is a perfect square
		if N%i == 0 {
			sum += i
			sum += N / i
		}
	}
	return sum
}
