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
	a := part1(instructions)
	fmt.Println(a)
	fmt.Println(part2(instructions, a))
}

type Instruction struct {
	in    map[string]string
	op    string
	out   string
	left  int
	right int
}

func parseInput(input []string) []Instruction {
	instructions := []Instruction{}
	for _, line := range input {
		parts := strings.Split(line, " -> ")
		curr := Instruction{}
		curr.out = parts[1]
		curr.in = map[string]string{}
		leftParts := strings.Split(parts[0], " ")
		l := len(leftParts)
		if l == 1 {
			curr.op = "SET"
			curr.right = 0
			if n, err := strconv.Atoi(leftParts[0]); err == nil {
				curr.left = n
			} else {
				curr.in["left"] = leftParts[0]
				curr.left = -1
			}
		} else if l == 2 {
			curr.op = leftParts[0]
			curr.left = 65535
			if n, err := strconv.Atoi(leftParts[1]); err == nil {
				curr.right = n
			} else {
				curr.in["right"] = leftParts[1]
				curr.right = -1
			}
		} else {
			curr.op = leftParts[1]
			if n, err := strconv.Atoi(leftParts[0]); err == nil {
				curr.left = n
			} else {
				curr.in["left"] = leftParts[0]
				curr.left = -1
			}
			if n, err := strconv.Atoi(leftParts[2]); err == nil {
				curr.right = n
			} else {
				curr.in["right"] = leftParts[2]
				curr.right = -1
			}
		}
		instructions = append(instructions, curr)
	}
	return instructions
}

func part1(instructions []Instruction) int {
	wires := map[string]int{}
	for len(instructions) > 0 {
		instructions2 := []Instruction{}
		for _, instruc := range instructions {
			for side, wire := range instruc.in {
				if n, ok := wires[wire]; ok {
					if side == "left" {
						instruc.left = n
					} else {
						instruc.right = n
					}
				}
			}
			if instruc.left >= 0 && instruc.right >= 0 {
				switch instruc.op {
				case "SET":
					wires[instruc.out] = instruc.left
				case "AND":
					wires[instruc.out] = instruc.left & instruc.right
				case "OR":
					wires[instruc.out] = instruc.left | instruc.right
				case "NOT":
					wires[instruc.out] = instruc.left ^ instruc.right
				case "LSHIFT":
					wires[instruc.out] = instruc.left << instruc.right
				case "RSHIFT":
					wires[instruc.out] = instruc.left >> instruc.right
				default:
					panic(instruc.op)
				}
			} else {
				instructions2 = append(instructions2, instruc)
			}
		}
		instructions = instructions2
	}
	return wires["a"]
}

func part2(instructions []Instruction, a int) int {
	wires := map[string]int{}
	wires["b"] = a
	for i, instruc := range instructions {
		if instruc.out == "b" {
			instructions = append(instructions[:i], instructions[i+1:]...)
			break
		}
	}
	for len(instructions) > 0 {
		instructions2 := []Instruction{}
		for _, instruc := range instructions {
			for side, wire := range instruc.in {
				if n, ok := wires[wire]; ok {
					if side == "left" {
						instruc.left = n
					} else {
						instruc.right = n
					}
				}
			}
			if instruc.left >= 0 && instruc.right >= 0 {
				switch instruc.op {
				case "SET":
					wires[instruc.out] = instruc.left
				case "AND":
					wires[instruc.out] = instruc.left & instruc.right
				case "OR":
					wires[instruc.out] = instruc.left | instruc.right
				case "NOT":
					wires[instruc.out] = instruc.left ^ instruc.right
				case "LSHIFT":
					wires[instruc.out] = instruc.left << instruc.right
				case "RSHIFT":
					wires[instruc.out] = instruc.left >> instruc.right
				default:
					panic(instruc.op)
				}
			} else {
				instructions2 = append(instructions2, instruc)
			}
		}
		instructions = instructions2
	}
	return wires["a"]
}
