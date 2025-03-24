package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/weitecklee/adventofcode/utils"
)

func main() {
	_, filename, _, _ := runtime.Caller(0)
	dirname := filepath.Dir(filename)
	inputFilePath := filepath.Join(dirname, "input.txt")
	data, err := os.ReadFile(inputFilePath)
	if err != nil {
		panic(err)
	}
	instructions := parseInput(strings.Split(string(data), "\n"))
	fmt.Println(solve(instructions))
}

type Instruction struct {
	ins string
	val int
}

func parseInput(data []string) []*Instruction {
	instructions := make([]*Instruction, len(data))
	for i, line := range data {
		if line == "noop" {
			instructions[i] = &Instruction{"noop", 0}
		} else {
			parts := strings.Split(line, " ")
			n, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}
			instructions[i] = &Instruction{parts[0], n}
		}
	}
	return instructions
}

func solve(instructions []*Instruction) (int, string) {
	registerX := 1
	cycle := 0
	cycleHistory := make(map[int]int)
	cycleHistory[cycle] = registerX
	for _, instruction := range instructions {
		cycle++
		cycleHistory[cycle] = registerX
		if instruction.ins == "addx" {
			cycle++
			registerX += instruction.val
			cycleHistory[cycle] = registerX
		}
	}
	res := 0
	for i := 20; i <= 220; i += 40 {
		res += i * cycleHistory[i-1]
	}
	var sb strings.Builder
	for cycle := 0; cycle < len(cycleHistory); cycle++ {
		if cycle%40 == 0 {
			sb.WriteByte('\n')
		}
		if utils.AbsInt(cycleHistory[cycle]-(cycle%40)) <= 1 {
			sb.WriteRune('█')
		} else {
			sb.WriteByte(' ')
		}
	}
	return res, sb.String()
}
