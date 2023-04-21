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

	fmt.Println(part1(input))
}

type Instruction struct {
	typ  string
	args []string
}

func (i *Instruction) execute() {

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
