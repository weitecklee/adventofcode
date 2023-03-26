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

func program(prog int, i int, snd0 *[]int, snd1 *[]int, instructions *[]Instruction, registers *map[string]int, count int) (int, int) {
	instruction := (*instructions)[i]
	switch instruction.action {
	case "snd":
		if prog == 0 {
			*snd0 = append(*snd0, (*registers)[instruction.subject])
		} else {
			*snd1 = append(*snd1, (*registers)[instruction.subject])
			count++
		}
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
		if prog == 0 {
			(*registers)[instruction.subject] = (*snd1)[0]
			*snd1 = (*snd1)[1:]
		} else {
			(*registers)[instruction.subject] = (*snd0)[0]
			*snd0 = (*snd0)[1:]
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
	return i + 1, count
}

func part2(instructions *[]Instruction) int {
	registers0 := map[string]int{}
	registers1 := map[string]int{}
	registers0["1"] = 1 // easiest way to account for input line "jgz 1 3",
	registers1["1"] = 1 // only occurrence of subject not pointing to register
	registers1["p"] = 1
	snd0 := []int{}
	snd1 := []int{}
	count := 0
	i := 0
	j := 0
loop:
	for i < len(*instructions) && i >= 0 && j < len(*instructions) && j >= 0 {
		for (*instructions)[i].action == "rcv" && len(snd1) == 0 {
			if (*instructions)[j].action == "rcv" && len(snd0) == 0 {
				break loop
			}
			j, count = program(1, j, &snd0, &snd1, instructions, &registers1, count)
		}
		for (*instructions)[j].action == "rcv" && len(snd0) == 0 {
			if (*instructions)[i].action == "rcv" && len(snd1) == 0 {
				break loop
			}
			i, count = program(0, i, &snd0, &snd1, instructions, &registers0, count)
		}
		i, count = program(0, i, &snd0, &snd1, instructions, &registers0, count)
		j, count = program(1, j, &snd0, &snd1, instructions, &registers1, count)
	}
	return count
}
