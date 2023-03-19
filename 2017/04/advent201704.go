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

func part1(input []string) int {
	res := 0
	for _, row := range input {
		words := strings.Split(row, " ")
		wordMap := map[string]bool{}
		valid := true
		for _, word := range words {
			if wordMap[word] {
				valid = false
				break
			}
			wordMap[word] = true
		}
		if valid {
			res++
		}
	}
	return res
}
