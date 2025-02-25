package main

import (
	"fmt"
	"math"
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
	part2(puzzleInput)
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

func runRobot(puzzleInput []int, startingPaint int) map[[2]int]int {
	inChan := make(chan int, 1)
	defer close(inChan)
	outChan := make(chan int)
	var wg sync.WaitGroup
	robot := intcode.NewIntcodeProgram(puzzleInput, inChan, outChan, &wg)
	wg.Add(1)
	go robot.Run()

	panels := make(map[[2]int]int)
	robotPos := [2]int{0, 0}
	robotDir := [2]int{-1, 0}
	panels[robotPos] = startingPaint
	inChan <- panels[robotPos]
	for paint := range outChan {
		panels[robotPos] = paint
		if <-outChan == 1 {
			robotDir[0], robotDir[1] = robotDir[1], -robotDir[0]
		} else {
			robotDir[0], robotDir[1] = -robotDir[1], robotDir[0]
		}
		robotPos[0], robotPos[1] = robotPos[0]+robotDir[0], robotPos[1]+robotDir[1]
		inChan <- panels[robotPos]
	}
	return panels

}

func part1(puzzleInput []int) int {
	return len(runRobot(puzzleInput, 0))
}

func part2(puzzleInput []int) {
	panels := runRobot(puzzleInput, 1)
	minRow := math.MaxInt
	maxRow := math.MinInt
	minCol := minRow
	maxCol := maxRow
	for k := range panels {
		r, c := k[0], k[1]
		if r < minRow {
			minRow = r
		}
		if r > maxRow {
			maxRow = r
		}
		if c < minCol {
			minCol = c
		}
		if c > maxCol {
			maxCol = c
		}
	}

	rows := maxRow - minRow + 1
	cols := maxCol - minCol + 1
	hull := make([][]int, rows)
	for i := range hull {
		hull[i] = make([]int, cols)
	}

	for k, v := range panels {
		r, c := k[0]-minRow, k[1]-minCol
		hull[r][c] = v
	}

	var sb strings.Builder
	for _, row := range hull {
		for _, v := range row {
			if v == 1 {
				sb.WriteByte('#')
			} else {
				sb.WriteByte(' ')
			}
		}
		sb.WriteByte('\n')
	}
	fmt.Println(sb.String())
}
