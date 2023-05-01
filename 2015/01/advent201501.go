package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(data)
	fmt.Println(solve(input))
}

func solve(input string) (int, int) {
	floor := 0
	part2 := -1
	for i, c := range input {
		if c == '(' {
			floor++
		} else {
			floor--
		}
		if floor < 0 && part2 < 0 {
			part2 = i + 1
		}
	}
	return floor, part2
}
