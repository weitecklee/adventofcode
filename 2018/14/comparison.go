package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input, _ := strconv.Atoi(string(data))
	fmt.Println(part1(input))
	fmt.Println(part1v2(input))
}

// Comparison between original method and using channel.
// Channel method actually slightly slower (~108ms to ~135ms).

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
	for i := 2; i < input; i++ {
		<-recipeChan
	}
	res := ""
	for i := 0; i < 10; i++ {
		res += strconv.Itoa(<-recipeChan)
	}
	return res
}
