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
	fmt.Println(part1(puzzleInput))
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

func factorial(n int) int {
	res := 1
	for i := range n {
		res *= (i + 1)
	}
	return res
}

func generatePermutations(sl []int) [][]int {
	if len(sl) == 1 {
		return [][]int{sl}
	}
	res := make([][]int, 0, factorial(len(sl)))
	for i, n := range sl {
		others := make([]int, 0, len(sl)-1)
		others = append(others, sl[:i]...)
		others = append(others, sl[i+1:]...)
		for _, perm := range generatePermutations(others) {
			res = append(res, append([]int{n}, perm...))
		}
	}
	return res
}

func amplifierOutput(program, sequence []int) int {
	var wg sync.WaitGroup
	inChanA := make(chan int)
	outChanA := make(chan int)
	inChanB := make(chan int)
	outChanB := make(chan int)
	inChanC := make(chan int)
	outChanC := make(chan int)
	inChanD := make(chan int)
	outChanD := make(chan int)
	inChanE := make(chan int)
	outChanE := make(chan int)
	ampA := intcode.NewIntcodeProgram(program, inChanA, outChanA, &wg)
	ampB := intcode.NewIntcodeProgram(program, inChanB, outChanB, &wg)
	ampC := intcode.NewIntcodeProgram(program, inChanC, outChanC, &wg)
	ampD := intcode.NewIntcodeProgram(program, inChanD, outChanD, &wg)
	ampE := intcode.NewIntcodeProgram(program, inChanE, outChanE, &wg)
	wg.Add(5)
	go ampA.Run()
	inChanA <- sequence[0]
	inChanA <- 0
	go ampB.Run()
	inChanB <- sequence[1]
	inChanB <- <-outChanA
	go ampC.Run()
	inChanC <- sequence[2]
	inChanC <- <-outChanB
	go ampD.Run()
	inChanD <- sequence[3]
	inChanD <- <-outChanC
	go ampE.Run()
	inChanE <- sequence[4]
	inChanE <- <-outChanD
	return <-outChanE
}

func part1(puzzleInput []int) int {
	sl := []int{0, 1, 2, 3, 4}
	perms := generatePermutations(sl)
	res := make([]int, len(perms))
	for i, perm := range perms {
		res[i] = amplifierOutput(puzzleInput, perm)
	}
	max := 0
	for _, n := range res {
		if n > max {
			max = n
		}
	}
	return max
}
