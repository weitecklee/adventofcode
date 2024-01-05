package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	_, filename, _, _ := runtime.Caller(0)
	dirname := filepath.Dir(filename)
	inputFilePath := filepath.Join(dirname, "input.txt")
	data, err := os.ReadFile(inputFilePath)
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), ",")
	part1, part2 := solve(input)
	fmt.Println(part1)
	fmt.Println(part2)
}

func calculateSteps(pos [2]float64) int {
	a := int(math.Abs(pos[0]))
	b := int(math.Abs(pos[1]))
	if a > b {
		return (a-b)/2 + b
	} else {
		return (b-a)/2 + a
	}
}

func solve(input []string) (int, int) {
	pos := [2]float64{0, 0}
	maxDist := 0
	for _, step := range input {
		switch step {
		case "n":
			pos[1] += 2
		case "ne":
			pos[0]++
			pos[1]++
		case "nw":
			pos[0]--
			pos[1]++
		case "s":
			pos[1] -= 2
		case "se":
			pos[0]++
			pos[1]--
		case "sw":
			pos[0]--
			pos[1]--
		default:
			panic("What is this? " + step)
		}
		dist := calculateSteps(pos)
		if dist > maxDist {
			maxDist = dist
		}
	}
	return calculateSteps(pos), maxDist
}
