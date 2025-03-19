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

func isInTractorBeam(icProgram []int, x, y int) bool {
	ch := make(chan int)
	ic := intcode.NewIntcodeProgram(icProgram, ch)
	go ic.Run()
	<-ch
	ch <- x
	<-ch
	ch <- y
	res := <-ch
	<-ch
	return res == 1
}

func part1(puzzleInput []int) int {
	res := 0
	for x := range 50 {
		for y := range 50 {
			if isInTractorBeam(puzzleInput, x, y) {
				res++
			}
		}
	}
	return res
}

func part2(puzzleInput []int) int {
	yMap := make(map[int][2]int)
	var x0, x1 int
	y := 100
	h, w := 100-1, 100-1
	for x0 > 0 && isInTractorBeam(puzzleInput, x0, y) {
		x0--
	}
	for !isInTractorBeam(puzzleInput, x0, y) {
		x0++
	}
	x1 = x0
	for isInTractorBeam(puzzleInput, x1, y) {
		x1++
	}
	for !isInTractorBeam(puzzleInput, x1, y) {
		x1--
	}
	yMap[y] = [2]int{x0, x1}
	for {
		y++
		for x0 > 0 && isInTractorBeam(puzzleInput, x0, y) {
			x0--
		}
		for !isInTractorBeam(puzzleInput, x0, y) {
			x0++
		}
		for isInTractorBeam(puzzleInput, x1, y) {
			x1++
		}
		for !isInTractorBeam(puzzleInput, x1, y) {
			x1--
		}
		yMap[y] = [2]int{x0, x1}
		if x1-x0 >= w {
			pair := yMap[y-h]
			x00 := pair[0]
			x01 := pair[1]
			if x0 >= x00 && x0 <= x01 && x0+w >= x00 && x0+w <= x01 {
				return x0*10000 + y - h
			}
		}
	}
}
