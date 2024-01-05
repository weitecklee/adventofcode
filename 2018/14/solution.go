package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
)

func main() {
	_, filename, _, _ := runtime.Caller(0)
	dirname := filepath.Dir(filename)
	inputFilePath := filepath.Join(dirname, "input.txt")
	data, err := os.ReadFile(inputFilePath)
	if err != nil {
		panic(err)
	}
	input, _ := strconv.Atoi(string(data))
	fmt.Println(part1(input))
	fmt.Println(part2(string(data)))
}

func recipeChannel(recipeChan chan int) {
	elf1 := 0
	elf2 := 1
	recipes := []int{3, 7}
	recipeChan <- 3
	recipeChan <- 7
	for {
		currentRecipe := recipes[elf1] + recipes[elf2]
		for _, c := range strconv.Itoa(currentRecipe) {
			n, _ := strconv.Atoi(string(c))
			recipeChan <- n
			recipes = append(recipes, n)
		}
		elf1 = (elf1 + recipes[elf1] + 1) % len(recipes)
		elf2 = (elf2 + recipes[elf2] + 1) % len(recipes)
	}
}

func part1(input int) string {
	recipeChan := make(chan int, 10000)
	go recipeChannel(recipeChan)
	for i := 0; i < input; i++ {
		<-recipeChan
	}
	res := ""
	for i := 0; i < 10; i++ {
		res += strconv.Itoa(<-recipeChan)
	}
	return res
}

func part2(input string) int {
	recipeChan := make(chan int, 10000)
	go recipeChannel(recipeChan)
	sequence := ""
	for i := 0; i < len(input); i++ {
		sequence += strconv.Itoa(<-recipeChan)
	}
	i := 0
	for sequence != input {
		i++
		sequence = sequence[1:] + strconv.Itoa(<-recipeChan)
	}
	return i
}
