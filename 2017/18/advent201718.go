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

func part1(input []string) int {
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
