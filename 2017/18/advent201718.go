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

func execute(i int, snd int, instructions *[]Instruction, registers *map[string]int) (int, int) {
	instruction := (*instructions)[i]
	switch instruction.action {
	case "snd":
		snd = (*registers)[instruction.subject]
	case "set":
		if instruction.valueString != "" {
			(*registers)[instruction.subject] = (*registers)[instruction.valueString]
		} else {
			(*registers)[instruction.subject] = instruction.valueInt
		}
	case "add":
		if instruction.valueString != "" {
			(*registers)[instruction.subject] += (*registers)[instruction.valueString]
		} else {
			(*registers)[instruction.subject] += instruction.valueInt
		}
	case "mul":
		if instruction.valueString != "" {
			(*registers)[instruction.subject] *= (*registers)[instruction.valueString]
		} else {
			(*registers)[instruction.subject] *= instruction.valueInt
		}
	case "mod":
		if instruction.valueString != "" {
			(*registers)[instruction.subject] %= (*registers)[instruction.valueString]
		} else {
			(*registers)[instruction.subject] %= instruction.valueInt
		}
	case "rcv":
		if (*registers)[instruction.subject] != 0 {
			i = -2
		}
	case "jgz":
		if (*registers)[instruction.subject] > 0 {
			if instruction.valueString != "" {
				i += (*registers)[instruction.valueString] - 1
			} else {
				i += instruction.valueInt - 1
			}
		}
	}
	return i + 1, snd
}

func part1(instructions *[]Instruction) int {
	registers := map[string]int{}
	snd := 0
	i := 0
	for i < len(*instructions)-1 && i >= 0 {
		i, snd = execute(i, snd, instructions, &registers)
	}
	return snd
}
