package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
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

func program(prog int, inChan chan int, outChan chan int, instructions *[]Instruction, count *int, wg *sync.WaitGroup) {
	registers := map[string]int{}
	registers["1"] = 1 // easiest way to account for input line "jgz 1 3",
	registers["p"] = prog
	i := 0
	for i >= 0 && i < len(*instructions) {
		instruction := (*instructions)[i]
		switch instruction.action {
		case "snd":
			outChan <- registers[instruction.subject]
			if prog == 1 {
				*count++
			}
		case "set":
			if instruction.valueString != "" {
				registers[instruction.subject] = registers[instruction.valueString]
			} else {
				registers[instruction.subject] = instruction.valueInt
			}
		case "add":
			if instruction.valueString != "" {
				registers[instruction.subject] += registers[instruction.valueString]
			} else {
				registers[instruction.subject] += instruction.valueInt
			}
		case "mul":
			if instruction.valueString != "" {
				registers[instruction.subject] *= registers[instruction.valueString]
			} else {
				registers[instruction.subject] *= instruction.valueInt
			}
		case "mod":
			if instruction.valueString != "" {
				registers[instruction.subject] %= registers[instruction.valueString]
			} else {
				registers[instruction.subject] %= instruction.valueInt
			}
		case "rcv":
			select {
			case registers[instruction.subject] = <-inChan:
			case <-time.After(time.Second):
				wg.Done()
				return
			}
		case "jgz":
			if registers[instruction.subject] > 0 {
				if instruction.valueString != "" {
					i += registers[instruction.valueString] - 1
				} else {
					i += instruction.valueInt - 1
				}
			}
		}
		i++
	}
}

func part2(instructions *[]Instruction) int {
	snd0 := make(chan int, 10000)
	snd1 := make(chan int, 10000)
	count := 0
	var wg sync.WaitGroup
	wg.Add(2)
	go program(0, snd1, snd0, instructions, &count, &wg)
	go program(1, snd0, snd1, instructions, &count, &wg)
	wg.Wait()
	return count
}
