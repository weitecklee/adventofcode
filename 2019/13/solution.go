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
	res := 0
	arcade := intcode.NewIntcodeProgram(puzzleInput, ch)
	var tileId int
	go arcade.Run()
	for range ch {
		<-ch
		tileId = <-ch
		if tileId == 2 {
			res++
		}
	}
	return res
}

func part2(puzzleInput []int) int {
	ch := make(chan int)
	puzzleInput[0] = 2
	arcade := intcode.NewIntcodeProgram(puzzleInput, ch)
	var x, tileId, score, ballPos, paddlePos int
	go arcade.Run()
	for x = range ch {
		if x == intcode.REQUESTSIGNAL {
			if ballPos > paddlePos {
				ch <- 1
			} else if ballPos < paddlePos {
				ch <- -1
			} else {
				ch <- 0
			}
		} else {
			<-ch
			tileId = <-ch
			if tileId == 3 {
				paddlePos = x
			} else if tileId == 4 {
				ballPos = x
			} else if x == -1 {
				score = tileId
			}
		}
	}
	return score
}
