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
	fmt.Println(part2(input))
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

/*
	Part 2 solved by analyzing solution pattern for nElves.
	First 10 solutions:
		1: 1
		2: 1
		3: 3
		4: 1
		5: 2
		6: 3
		7: 5
		8: 7
		9: 9
		10: 1
	Pattern is to increment by 1 until lastElf is equal to half of nElves, then increment by 2.
	When lastElf = nElves, reset lastElf and increment to 1 for the next one.
	I'm sure there is an actual O(1) equation instead of iterating through but this will do for now.
*/

func part2(nElves int) int {
	lastElf := 1
	inc := 1
	for i := 1; i < nElves; i++ {
		if i == lastElf {
			lastElf = 1
			inc = 1
		} else {
			lastElf += inc
		}
		if i+1 == 2*lastElf {
			inc++
		}
	}
	return lastElf
}
