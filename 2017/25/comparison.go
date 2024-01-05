package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func main() {
	_, filename, _, _ := runtime.Caller(0)
	dirname := filepath.Dir(filename)
	inputFilePath := filepath.Join(dirname, "input.txt")
	data, err := os.ReadFile(inputFilePath)
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n")
	fmt.Println(part1(parseInput(input)))
	fmt.Println(part1v2(parseInput(input)))
}

// comparison between setting to zero and deleting value from map
// deleting method is actually slightly slower (~837ms to ~921ms)

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	fmt.Printf("%v: %v\n", msg, time.Since(start))
}

type Instruction struct {
	zeroWrite int
	zeroMove  int
	zeroState string
	oneWrite  int
	oneMove   int
	oneState  string
}

func parseInput(input []string) (*map[string]*Instruction, string, int) {
	reState := regexp.MustCompile(`state [A-Z]`)
	reMove := regexp.MustCompile(`right`)
	reNum := regexp.MustCompile(`\d+`)
	beginningState := reState.FindString(input[0])
	steps, _ := strconv.Atoi(reNum.FindString(input[1]))
	i := 3
	instructions := map[string]*Instruction{}
	for i < len(input) {
		state := reState.FindString(input[i])
		zeroWrite, _ := strconv.Atoi(reNum.FindString(input[i+2]))
		zeroMove := -1
		if reMove.MatchString(input[i+3]) {
			zeroMove = 1
		}
		zeroState := reState.FindString(input[i+4])
		oneWrite, _ := strconv.Atoi(reNum.FindString(input[i+6]))
		oneMove := -1
		if reMove.MatchString(input[i+7]) {
			oneMove = 1
		}
		oneState := reState.FindString(input[i+8])
		instruction := Instruction{
			zeroWrite: zeroWrite,
			zeroMove:  zeroMove,
			zeroState: zeroState,
			oneWrite:  oneWrite,
			oneMove:   oneMove,
			oneState:  oneState,
		}
		instructions[state] = &instruction
		i += 10
	}
	return &instructions, beginningState, steps
}

func runMachine(slot *int, state *string, tape *map[int]int, instructions *map[string]*Instruction) {
	if (*tape)[*slot] == 0 {
		(*tape)[*slot] = (*instructions)[*state].zeroWrite
		*slot += (*instructions)[*state].zeroMove
		*state = (*instructions)[*state].zeroState
	} else {
		(*tape)[*slot] = (*instructions)[*state].oneWrite
		*slot += (*instructions)[*state].oneMove
		*state = (*instructions)[*state].oneState
	}
}

func part1(instructions *map[string]*Instruction, state string, steps int) int {
	defer duration(track("part1"))
	tape := map[int]int{}
	slot := 0
	for i := 0; i < steps; i++ {
		runMachine(&slot, &state, &tape, instructions)
	}
	count := 0
	for _, v := range tape {
		if v == 1 {
			count++
		}
	}
	return count
}

func runMachinev2(slot *int, state *string, tape *map[int]int, instructions *map[string]*Instruction) {
	if (*tape)[*slot] == 1 {
		if (*instructions)[*state].oneWrite == 0 {
			delete(*tape, *slot)
		}
		*slot += (*instructions)[*state].oneMove
		*state = (*instructions)[*state].oneState
	} else {
		if (*instructions)[*state].zeroWrite == 1 {
			(*tape)[*slot] = 1
		}
		*slot += (*instructions)[*state].zeroMove
		*state = (*instructions)[*state].zeroState
	}
}

func part1v2(instructions *map[string]*Instruction, state string, steps int) int {
	defer duration(track("part1v2"))
	tape := map[int]int{}
	slot := 0
	for i := 0; i < steps; i++ {
		runMachinev2(&slot, &state, &tape, instructions)
	}
	return len(tape)
}
