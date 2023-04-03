package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input, _ := strconv.Atoi(string(data))
	fmt.Println(part1(input))
}

func part1(input int) string {
	elf1 := 0
	elf2 := 1
	recipes := map[int]int{}
	recipes[0] = 3
	recipes[1] = 7
	for len(recipes) < input+10 {
		currentRecipe := recipes[elf1] + recipes[elf2]
		for _, c := range strconv.Itoa(currentRecipe) {
			n, _ := strconv.Atoi(string(c))
			recipes[len(recipes)] = n
		}
		elf1 = (elf1 + recipes[elf1] + 1) % len(recipes)
		elf2 = (elf2 + recipes[elf2] + 1) % len(recipes)
	}
	res := ""
	for i := input; i < input+10; i++ {
		res += strconv.Itoa(recipes[i])
	}
	return res
}
