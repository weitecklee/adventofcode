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
	ch := make(chan int)
	ic := intcode.NewIntcodeProgram(puzzleInput, ch)
	go ic.Run()
	<-ch
	ch <- 1
	var output int
	for {
		tmp := <-ch
		if tmp == intcode.ENDSIGNAL {
			return output
		}
		output = tmp
	}
}

func part2(puzzleInput []int) int {
	ch := make(chan int)
	ic := intcode.NewIntcodeProgram(puzzleInput, ch)
	go ic.Run()
	<-ch
	ch <- 5
	return <-ch
}
