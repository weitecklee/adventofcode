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

func inputCommand(ch chan int, command string) {
	for _, c := range command {
		ch <- int(c)
		<-ch
	}
	ch <- 10
	<-ch
}

func displayMessage(ch chan int, valueToWatch int, display bool) {
	var ret int
	for {
		ret = <-ch
		if ret == intcode.ENDSIGNAL ||
			ret == intcode.REQUESTSIGNAL ||
			ret == valueToWatch {
			break
		}
		if display {
			fmt.Printf("%c", ret)
		}
	}

}

func runCommands(puzzleInput []int, commands []string, debugMode bool) int {
	ch := make(chan int)
	ic := intcode.NewIntcodeProgram(puzzleInput, ch)
	go ic.Run()
	displayMessage(ch, -1, false)
	for _, command := range commands {
		inputCommand(ch, command)
	}

	if debugMode {
		displayMessage(ch, -1, true)
		return -1
	}

	displayMessage(ch, 10, false)
	<-ch
	return <-ch

}

func part1(puzzleInput []int) int {

	// If any of A/B/C is false, J is set to true. (There is a hole to jump over.)
	// T is set to D. (If true, D is ground and safe to jump to.)
	// if both T and J are true, J is true. (There is a hole to jump over and ground to land on.)

	commands := []string{
		"NOT A J",
		"NOT B T",
		"OR T J",
		"NOT C T",
		"OR T J",
		"NOT D T",
		"NOT T T",
		"AND T J",
		"WALK",
	}
	return runCommands(puzzleInput, commands, false)
}

func part2(puzzleInput []int) int {

	// same as above, but now check if E and H are both holes
	// set T to (E || H) and do another "AND T J"

	commands := []string{
		"NOT A J",
		"NOT B T",
		"OR T J",
		"NOT C T",
		"OR T J",
		"NOT D T",
		"NOT T T",
		"AND T J",
		"NOT E T",
		"NOT T T",
		"OR H T",
		"AND T J",
		"RUN",
	}
	return runCommands(puzzleInput, commands, false)
}
