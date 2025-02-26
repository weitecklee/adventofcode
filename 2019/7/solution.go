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

func amplifierOutput(program, sequence []int, withFeedback bool) int {
	chA := make(chan int)
	chB := make(chan int)
	chC := make(chan int)
	chD := make(chan int)
	chE := make(chan int)
	ampA := intcode.NewIntcodeProgram(program, chA)
	ampB := intcode.NewIntcodeProgram(program, chB)
	ampC := intcode.NewIntcodeProgram(program, chC)
	ampD := intcode.NewIntcodeProgram(program, chD)
	ampE := intcode.NewIntcodeProgram(program, chE)
	go ampA.Run()
	go ampB.Run()
	go ampC.Run()
	go ampD.Run()
	go ampE.Run()
	<-chA
	<-chB
	<-chC
	<-chD
	<-chE
	chA <- sequence[0]
	chB <- sequence[1]
	chC <- sequence[2]
	chD <- sequence[3]
	chE <- sequence[4]
	var outputE int
	for {
		if <-chA == intcode.ENDSIGNAL {
			break
		}
		<-chB
		<-chC
		<-chD
		<-chE
		chA <- outputE
		chB <- <-chA
		chC <- <-chB
		chD <- <-chC
		chE <- <-chD
		outputE = <-chE
		if !withFeedback {
			return outputE
		}
	}
	return outputE
}

func part1(puzzleInput []int) int {
	sl := []int{0, 1, 2, 3, 4}
	perms := generatePermutations(sl)
	max := 0
	for _, perm := range perms {
		res := amplifierOutput(puzzleInput, perm, false)
		if res > max {
			max = res
		}
	}
	return max
}

func part2(puzzleInput []int) int {
	sl := []int{5, 6, 7, 8, 9}
	perms := generatePermutations(sl)
	max := 0
	for _, perm := range perms {
		res := amplifierOutput(puzzleInput, perm, true)
		if res > max {
			max = res
		}
	}
	return max
}
