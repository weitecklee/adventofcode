package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"

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
	fmt.Println(runBOOST(puzzleInput, 1))
	fmt.Println(runBOOST(puzzleInput, 2))
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

func runBOOST(puzzleInput []int, initialInput int) int {
	inChan := make(chan int)
	defer close(inChan)
	outChan := make(chan int)
	var wg sync.WaitGroup
	ic := intcode.NewIntcodeProgram(puzzleInput, inChan, outChan, &wg)
	wg.Add(1)
	go ic.Run()
	inChan <- initialInput
	return <-outChan
}
