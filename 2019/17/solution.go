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

var directions = [][2]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}

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

func part1(puzzleInput []int) int {
	ch := make(chan int)
	ic := intcode.NewIntcodeProgram(puzzleInput, ch)
	go ic.Run()
	var sb strings.Builder
	for {
		output := <-ch
		if output == intcode.ENDSIGNAL {
			break
		}
		sb.WriteRune(rune(output))
	}

	area := strings.Split(sb.String(), "\n")
	res := 0

	for r, row := range area {
	loop:
		for c, chr := range row {
			if chr == '#' {
				for _, dir := range directions {
					r2, c2 := r+dir[0], c+dir[1]
					if r2 < 0 || c2 < 0 || r2 >= len(area) || c2 >= len(area[0]) || area[r2][c2] == '.' {
						continue loop
					}
				}
				res += r * c
			}
		}
	}

	return res
}
