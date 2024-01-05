package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
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
	fmt.Println(part1v2(input))
	fmt.Println(part1vA(input))
	fmt.Println(part1v2A(input))
	fmt.Println(part2(string(data)))
	fmt.Println(part2vA(string(data)))
}

// Comparison between original method and using channel.
// Channel method actually slightly slower (~108ms to ~135ms).

// Added comparison between using maps and arrays.
// Arrays much faster. (~21ms and ~43ms, compare to above).
// Part 2 goes from 8.2s to 2.4s

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	fmt.Printf("%v: %v\n", msg, time.Since(start))
}

func part1(input int) string {
	defer duration(track("part1"))
	elf1 := 0
	elf2 := 1
	recipes := map[int]int{}
	recipes[0] = 3
	recipes[1] = 7
	recipeLen := 2
	for recipeLen < input+10 {
		currentRecipe := recipes[elf1] + recipes[elf2]
		for _, c := range strconv.Itoa(currentRecipe) {
			n, _ := strconv.Atoi(string(c))
			recipes[recipeLen] = n
			recipeLen++
		}
		elf1 = (elf1 + recipes[elf1] + 1) % recipeLen
		elf2 = (elf2 + recipes[elf2] + 1) % recipeLen
	}
	res := ""
	for i := input; i < input+10; i++ {
		res += strconv.Itoa(recipes[i])
	}
	return res
}

func recipeChannel(recipeChan chan int) {
	elf1 := 0
	elf2 := 1
	recipes := map[int]int{}
	recipes[0] = 3
	recipes[1] = 7
	recipeChan <- 3
	recipeChan <- 7
	recipeLen := 2
	for {
		currentRecipe := recipes[elf1] + recipes[elf2]
		for _, c := range strconv.Itoa(currentRecipe) {
			n, _ := strconv.Atoi(string(c))
			recipeChan <- n
			recipes[recipeLen] = n
			recipeLen++
		}
		elf1 = (elf1 + recipes[elf1] + 1) % recipeLen
		elf2 = (elf2 + recipes[elf2] + 1) % recipeLen
	}
}

func part1v2(input int) string {
	defer duration(track("part1v2"))
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

func part1vA(input int) string {
	defer duration(track("part1vA"))
	elf1 := 0
	elf2 := 1
	recipes := []int{3, 7}
	for len(recipes) < input+10 {
		currentRecipe := recipes[elf1] + recipes[elf2]
		for _, c := range strconv.Itoa(currentRecipe) {
			n, _ := strconv.Atoi(string(c))
			recipes = append(recipes, n)
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

func recipeChannel2(recipeChan chan int) {
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

func part1v2A(input int) string {
	defer duration(track("part1v2A"))
	recipeChan := make(chan int, 10000)
	go recipeChannel2(recipeChan)
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
	defer duration(track("part2"))
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

func part2vA(input string) int {
	defer duration(track("part2vA"))
	recipeChan := make(chan int, 10000)
	go recipeChannel2(recipeChan)
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
