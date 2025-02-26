package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/weitecklee/adventofcode/2019/intcode"
)

func main() {
	_, filename, _, _ := runtime.Caller(0)
	dirname := filepath.Dir(filename)
	inputFilePath := filepath.Join(dirname, "input.txt")
	data, err := os.ReadFile(inputFilePath)
	if err != nil {
		panic(err)
	}
	puzzleInput := parseInput(strings.Split(string(data), ","))
	fmt.Println(part1(puzzleInput))
	fmt.Println(part2(puzzleInput))
}

func parseInput(data []string) []int {
	numbers := make([]int, 0, len(data))
	for _, s := range data {
		if n, err := strconv.Atoi(s); err != nil {
			panic(err)
		} else {
			numbers = append(numbers, n)
		}
	}
	return numbers
}

func part1(puzzleInput []int) int {
	puzzleInput[1] = 12
	puzzleInput[2] = 2
	ch := make(chan int)
	ic := intcode.NewIntcodeProgram(puzzleInput, ch)
	go ic.Run()
	<-ch
	return ic.Program[0]
}

func part2(puzzleInput []int) int {
	for noun := range 100 {
		for verb := range 100 {
			puzzleInput[1] = noun
			puzzleInput[2] = verb
			ch := make(chan int)
			ic := intcode.NewIntcodeProgram(puzzleInput, ch)
			go ic.Run()
			<-ch
			if ic.Program[0] == 19690720 {
				return noun*100 + verb
			}
		}
	}
	return -1
}
