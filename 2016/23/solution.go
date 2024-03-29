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
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

type Instruction struct {
	typ  string
	args []string
}

func part1(input []string) int {
	register := map[string]int{}
	instructions := []*Instruction{}
	for _, line := range input {
		parts := strings.Split(line, " ")
		instructions = append(instructions, &Instruction{
			typ:  parts[0],
			args: parts[1:],
		})
		for i := 1; i < len(parts); i++ {
			if n, err := strconv.Atoi(parts[i]); err == nil {
				register[parts[i]] = n
			}
		}
	}
	register["a"] = 7
	i := 0
	for i < len(instructions) {
		curr := instructions[i]
		switch curr.typ {
		case "cpy":
			if _, err := strconv.Atoi(curr.args[1]); err != nil {
				register[curr.args[1]] = register[curr.args[0]]
			}
		case "inc":
			register[curr.args[0]]++
		case "dec":
			register[curr.args[0]]--
		case "jnz":
			if register[curr.args[0]] != 0 {
				i += register[curr.args[1]] - 1
			}
		case "tgl":
			j := i + register[curr.args[0]]
			if j < len(instructions) {
				switch instructions[j].typ {
				case "inc":
					instructions[j].typ = "dec"
				case "dec":
					instructions[j].typ = "inc"
				case "tgl":
					instructions[j].typ = "inc"
				case "jnz":
					instructions[j].typ = "cpy"
				case "cpy":
					instructions[j].typ = "jnz"
				}
			}
		}
		i++
	}
	return register["a"]
}

func factorial(i int) int {
	if i > 1 {
		return i * factorial(i-1)
	}
	return 1
}

func part2(input []string) int {
	/*
		Analysis of the assembly code reveals that the first 19 lines calculate
		the factorial of `a`, during which lines 25, 23, 21, 19 are toggled (in that order, 1-indexed).
		After line 19 is toggled, the rest of the code is finally run, which is just adding 73*90
		to the result of `a`.	So the final result is a! + 73 * 90.
		Most likely, only lines 20 and 21 are different for other users.
	*/
	parts := strings.Split(input[19], " ")
	a, _ := strconv.Atoi(parts[1])
	parts = strings.Split(input[20], " ")
	b, _ := strconv.Atoi(parts[1])
	return factorial(12) + a*b
}
