package main

import (
	"fmt"
	"os"
	"regexp"
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
	cat  string
	valr string
	vali int
}

func parseInput(input []string) []Instruction {
	instructions := []Instruction{}
	re1 := regexp.MustCompile(`\w+`)
	re2 := regexp.MustCompile(`-?\d+`)
	for _, line := range input {
		matches1 := re1.FindAllString(line, -1)
		curr := Instruction{}
		if matches1 != nil {
			curr.cat = matches1[0]
			if len(matches1) > 1 {
				curr.valr = matches1[1]
			}
		}
		matches2 := re2.FindAllString(line, -1)
		if matches2 != nil {
			n, _ := strconv.Atoi(matches2[0])
			curr.vali = n
		}
		instructions = append(instructions, curr)
	}
	return instructions
}

func (I *Instruction) execute(register *map[string]int, i *int) {
	switch I.cat {
	case "hlf":
		(*register)[I.valr] /= 2
	case "tpl":
		(*register)[I.valr] *= 3
	case "inc":
		(*register)[I.valr]++
	case "jmp":
		*i += I.vali - 1
	case "jie":
		if (*register)[I.valr]%2 == 0 {
			*i += I.vali - 1
		}
	case "jio":
		if (*register)[I.valr] == 1 {
			*i += I.vali - 1
		}
	}
	*i++
}

func part1(instructions []Instruction) int {
	register := map[string]int{}
	i := 0
	for i < len(instructions) {
		instructions[i].execute(&register, &i)
	}
	return register["b"]
}

func part2(instructions []Instruction) int {
	register := map[string]int{}
	register["a"] = 1
	i := 0
	for i < len(instructions) {
		instructions[i].execute(&register, &i)
	}
	return register["b"]
}
