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
	var wg sync.WaitGroup
	inChan := make(chan int)
	defer close(inChan)
	outChan := make(chan int)
	res := 0
	arcade := intcode.NewIntcodeProgram(puzzleInput, inChan, outChan, &wg)
	var tileId int
	wg.Add(1)
	go arcade.Run()
	for range outChan {
		<-outChan
		tileId = <-outChan
		if tileId == 2 {
			res++
		}
	}
	return res
}

func part2(puzzleInput []int) int {
	var wg sync.WaitGroup
	inChan := make(chan int)
	defer close(inChan)
	outChan := make(chan int)
	puzzleInput[0] = 2
	arcade := intcode.NewIntcodeProgram(puzzleInput, inChan, outChan, &wg)
	var x, tileId, score, ballPos, paddlePos int
	wg.Add(1)
	go arcade.Run()
	for {
		for x = range outChan {
			if x == intcode.REQUESTSIGNAL || x == intcode.ENDSIGNAL {
				break
			}
			<-outChan
			tileId = <-outChan
			if tileId == 3 {
				paddlePos = x
			} else if tileId == 4 {
				ballPos = x
			} else if x == -1 {
				score = tileId
			}
		}
		if x == intcode.ENDSIGNAL {
			break
		}
		if ballPos > paddlePos {
			inChan <- 1
		} else if ballPos < paddlePos {
			inChan <- -1
		} else {
			inChan <- 0
		}
	}
	return score
}
