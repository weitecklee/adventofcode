package main

import (
	"fmt"
	"math"
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
	instructions := parseInput(input)
	fmt.Println(part1(instructions))
	fmt.Println(part2(instructions))
}

type Instruction struct {
	action      string
	subject     string
	valueString string
	valueInt    int
}

func parseInput(input []string) *[]Instruction {
	instructions := []Instruction{}
	for _, line := range input {
		instruction := Instruction{}
		parts := strings.Split(line, " ")
		instruction.action = parts[0]
		instruction.subject = parts[1]
		if len(parts) > 2 {
			if n, err := strconv.Atoi(parts[2]); err != nil {
				instruction.valueString = parts[2]
			} else {
				instruction.valueInt = n
			}
		}
		instructions = append(instructions, instruction)
	}
	return &instructions
}

func execute(i int, mul *int, instructions *[]Instruction, registers *map[string]int) int {
	instruction := (*instructions)[i]
	switch instruction.action {
	case "set":
		if instruction.valueString != "" {
			(*registers)[instruction.subject] = (*registers)[instruction.valueString]
		} else {
			(*registers)[instruction.subject] = instruction.valueInt
		}
	case "sub":
		if instruction.valueString != "" {
			(*registers)[instruction.subject] -= (*registers)[instruction.valueString]
		} else {
			(*registers)[instruction.subject] -= instruction.valueInt
		}
	case "mul":
		if instruction.valueString != "" {
			(*registers)[instruction.subject] *= (*registers)[instruction.valueString]
		} else {
			(*registers)[instruction.subject] *= instruction.valueInt
		}
		*mul++
	case "jnz":
		if (*registers)[instruction.subject] != 0 {
			if instruction.valueString != "" {
				i += (*registers)[instruction.valueString] - 1
			} else {
				i += instruction.valueInt - 1
			}
		}
	}
	return i + 1
}

func part1(instructions *[]Instruction) int {
	registers := map[string]int{}
	registers["1"] = 1
	mul := 0
	i := 0
	for i < len(*instructions) && i >= 0 {
		i = execute(i, &mul, instructions, &registers)
	}
	return mul
}

func part2(instructions *[]Instruction) int {
	// analysis of assembly code is necessary
	// (if you try and run the actual code it would take forever)
	// run the first 8 lines of codes to get the starting values
	registers := map[string]int{}
	registers["1"] = 1
	registers["a"] = 1
	i := 0
	mul := 0
	for i < 8 {
		i = execute(i, &mul, instructions, &registers)
	}
	// now have starting values for b, c
	b := registers["b"]
	c := registers["c"]
	h := 0

	// code loops from b to c in increments of 17
	// in each loop, initialize d = 2, e = 2, f = 1
	// then it's two nested loops from e to b inside d to b
	// checking if d * e == b then set f = 0
	// at the end, if f == 0, increment h
	// basically, incredibly inefficient method to check if b is prime or composite
	// loop is counting the number of composite numbers from b to c (inclusive) with increments of 17

	for b <= c {
		if b%2 != 0 {
			max := int(math.Sqrt(float64(b)))
			for j := 3; j <= max; j += 2 {
				if b%j == 0 {
					h++
					break
				}
			}
		} else {
			h++
		}
		b += 17
	}

	return h
}
