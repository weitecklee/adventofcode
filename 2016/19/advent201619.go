package main

import (
	"container/ring"
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

func part1(nElves int) int {
	elves := ring.New(nElves)
	for i := 0; i < nElves; i++ {
		elves.Value = i + 1
		elves = elves.Next()
	}
	for i := 0; i < nElves; i++ {
		elves.Unlink(1)
		elves = elves.Next()
	}
	return elves.Value.(int)
}
