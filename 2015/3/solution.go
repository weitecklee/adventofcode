package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	_, filename, _, _ := runtime.Caller(0)
	dirname := filepath.Dir(filename)
	inputFilePath := filepath.Join(dirname, "input.txt")
	data, err := os.ReadFile(inputFilePath)
	if err != nil {
		panic(err)
	}
	input := string(data)
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input string) int {
	houses := map[[2]int]bool{}
	loc := [2]int{0, 0}
	houses[loc] = true
	for _, dir := range input {
		switch dir {
		case '^':
			loc[1]++
		case 'v':
			loc[1]--
		case '<':
			loc[0]--
		case '>':
			loc[0]++
		}
		houses[loc] = true
	}
	return len(houses)
}

func part2(input string) int {
	houses := map[[2]int]bool{}
	locs := [2][2]int{
		{0, 0},
		{0, 0},
	}
	houses[locs[0]] = true
	for i, dir := range input {
		switch dir {
		case '^':
			locs[i%2][1]++
		case 'v':
			locs[i%2][1]--
		case '<':
			locs[i%2][0]--
		case '>':
			locs[i%2][0]++
		}
		houses[locs[i%2]] = true
	}
	return len(houses)
}
