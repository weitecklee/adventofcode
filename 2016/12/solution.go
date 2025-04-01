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
	puzzleInput := parseInput(strings.Split(string(data), "\n"))
	fmt.Println(part1(puzzleInput))
	fmt.Println(part2(puzzleInput))
}

type Instruction struct {
	instruc    string
	isAddress1 bool
	isAddress2 bool
	address1   string
	address2   string
	value1     int
	value2     int
}

func NewInstruction(s string) (*Instruction, error) {
	parts := strings.Split(s, " ")
	var isAddress1, isAddress2 bool
	var address1, address2 string
	var value1, value2 int
	if parts[0] == "cpy" || parts[0] == "jnz" {
		n1, err := strconv.Atoi(parts[1])
		if err != nil {
			isAddress1 = true
			address1 = parts[1]
		} else {
			value1 = n1
		}
		n2, err := strconv.Atoi(parts[2])
		if err != nil {
			isAddress2 = true
			address2 = parts[2]
		} else {
			value2 = n2
		}
	} else if parts[0] == "inc" || parts[0] == "dec" {
		isAddress1 = true
		address1 = parts[1]
	} else {
		return nil, fmt.Errorf("unexpected instruction: %s", s)
	}

	return &Instruction{parts[0], isAddress1, isAddress2, address1, address2, value1, value2}, nil
}

func parseInput(data []string) []*Instruction {
	puzzleInput := make([]*Instruction, len(data))
	for i, s := range data {
		instruction, err := NewInstruction(s)
		if err != nil {
			panic(err)
		}
		puzzleInput[i] = instruction
	}
	return puzzleInput
}

func execute(puzzleInput []*Instruction, cValue int) int {
	registers := make(map[string]int, 4)
	registers["c"] = cValue
	i := 0
	for i >= 0 && i < len(puzzleInput) {
		curr := puzzleInput[i]
		x := curr.value1
		if curr.isAddress1 {
			x = registers[curr.address1]
		}
		y := curr.value2
		if curr.isAddress2 {
			y = registers[curr.address2]
		}
		switch curr.instruc {
		case "cpy":
			registers[curr.address2] = x
		case "inc":
			registers[curr.address1]++
		case "dec":
			registers[curr.address1]--
		case "jnz":
			if x != 0 {
				i += y
				continue
			}
		}
		i++
	}
	return registers["a"]
}

func part1(puzzleInput []*Instruction) int {
	return execute(puzzleInput, 0)
}

func part2(puzzleInput []*Instruction) int {
	return execute(puzzleInput, 1)
}
