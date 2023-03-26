package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n")
	fmt.Println(part1(input))
	fmt.Println(part1v2(parseInput(input)))
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	fmt.Printf("%v: %v\n", msg, time.Since(start))
}

func part1(input []string) int {
	defer duration(track("part1"))
	registers := map[string]int{}
	snd := 0
	i := 0
loop:
	for i < len(input)-1 && i >= 0 {
		parts := strings.Split(input[i], " ")
		switch parts[0] {
		case "snd":
			snd = registers[parts[1]]
		case "set":
			if n, err := strconv.Atoi(parts[2]); err != nil {
				registers[parts[1]] = registers[parts[2]]
			} else {
				registers[parts[1]] = n
			}
		case "add":
			if n, err := strconv.Atoi(parts[2]); err != nil {
				registers[parts[1]] += registers[parts[2]]
			} else {
				registers[parts[1]] += n
			}
		case "mul":
			if n, err := strconv.Atoi(parts[2]); err != nil {
				registers[parts[1]] *= registers[parts[2]]
			} else {
				registers[parts[1]] *= n
			}
		case "mod":
			if n, err := strconv.Atoi(parts[2]); err != nil {
				registers[parts[1]] %= registers[parts[2]]
			} else {
				registers[parts[1]] %= n
			}
		case "rcv":
			if registers[parts[1]] != 0 {
				break loop
			}
		case "jgz":
			if registers[parts[1]] > 0 {
				if n, err := strconv.Atoi(parts[2]); err != nil {
					i += registers[parts[2]]
				} else {
					i += n
				}
				continue loop
			}
		}
		i++
	}
	return snd
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

func part1v2(instructions *[]Instruction) int {
	defer duration(track("part1v2"))
	registers := map[string]int{}
	snd := 0
	i := 0
	for i < len(*instructions)-1 && i >= 0 {
		i, snd = execute(i, snd, instructions, &registers)
	}
	return snd
}
