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
	input := strings.Split(string(data), "\n")
	fmt.Println(part1(input))
}

type Instruction struct {
	typ  string
	args []string
}

// func part1(input []string) int {
// 	register := map[string]int{}
// 	instructions := []*Instruction{}
// 	for _, line := range input {
// 		parts := strings.Split(line, " ")
// 		instructions = append(instructions, &Instruction{
// 			typ:  parts[0],
// 			args: parts[1:],
// 		})
// 		for i := 1; i < len(parts); i++ {
// 			if n, err := strconv.Atoi(parts[i]); err == nil {
// 				register[parts[i]] = n
// 			}
// 		}
// 	}
// 	register["a"] = 1
// 	i := 0
// loop:
// 	for i < len(instructions) {
// 		curr := instructions[i]
// 		switch curr.typ {
// 		case "cpy":
// 			if _, err := strconv.Atoi(curr.args[1]); err != nil {
// 				register[curr.args[1]] = register[curr.args[0]]
// 			}
// 		case "inc":
// 			register[curr.args[0]]++
// 		case "dec":
// 			register[curr.args[0]]--
// 		case "jnz":
// 			if register[curr.args[0]] != 0 {
// 				i += register[curr.args[1]] - 1
// 			}
// 		case "out":
// 			fmt.Println(register["a"], register["b"])
// 			if register["a"] == 0 {
// 				break loop
// 			}
// 		}
// 		i++
// 	}
// 	return register["a"]
// }

/*
	Playing around with assembly code shows it is repeatedly outputting the binary
	representation of a number, this number being `a + 643 * 4` (depending on your puzzle input)
	The answer then is the smallest `a` so that `a + 643 * 4` is equal to
	0b1010...1010
*/

func part1(input []string) int {
	parts := strings.Split(input[1], " ")
	c, _ := strconv.Atoi(parts[1])
	parts = strings.Split(input[2], " ")
	b, _ := strconv.Atoi(parts[1])
	i := 2
	a := 0
	for a < b*c {
		a += i
		i *= 4
	}
	return a - b*c
}
