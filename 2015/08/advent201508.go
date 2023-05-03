package main

import (
	"fmt"
	"os"
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

func memoryLen(line string) int {
	i := 1
	count := 0
	for i < len(line)-1 {
		if line[i] == '\\' {
			if line[i+1] == 'x' {
				i += 3
			} else {
				i++
			}
		}
		count++
		i++
	}
	return count
}

func part1(input []string) int {
	count := 0
	for _, line := range input {
		count += len(line) - memoryLen(line)
	}
	return count
}
