package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/weitecklee/adventofcode/utils"
)

func main() {
	_, filename, _, _ := runtime.Caller(0)
	dirname := filepath.Dir(filename)
	inputFilePath := filepath.Join(dirname, "input.txt")
	data, err := os.ReadFile(inputFilePath)
	if err != nil {
		panic(err)
	}
	puzzleInput := parseInput(strings.Split(string(data), "\n"))
	fmt.Println(solve(puzzleInput))
}

type Segment struct {
	direction string
	distance  int
}

var directions = map[string][2]int{"U": {0, 1}, "D": {0, -1}, "L": {-1, 0}, "R": {1, 0}}

func parseInput(data []string) [2][]Segment {
	var wires [2][]Segment
	for i, s := range data {
		parts := strings.Split(s, ",")
		var wire []Segment
		for _, part := range parts {
			n, err := strconv.Atoi(part[1:])
			if err != nil {
				panic(err)
			}
			wire = append(wire, Segment{string(part[0]), n})
		}
		wires[i] = wire
	}
	return wires
}

func solve(puzzleInput [2][]Segment) (int, int) {
	wire1Map := make(map[[2]int]int)
	pos := [2]int{0, 0}
	wire1Map[pos] = 0
	d := 1

	for _, segment := range puzzleInput[0] {
		direction := directions[segment.direction]
		for range segment.distance {
			pos[0] += direction[0]
			pos[1] += direction[1]
			if _, ok := wire1Map[pos]; !ok {
				wire1Map[pos] = d
			}
			d++
		}
	}

	pos = [2]int{0, 0}
	wire2Map := make(map[[2]int]int)
	wire2Map[pos] = 0
	d = 1

	var intersections [][2]int
	for _, segment := range puzzleInput[1] {
		direction := directions[segment.direction]
		for range segment.distance {
			pos[0] += direction[0]
			pos[1] += direction[1]
			if _, ok := wire1Map[pos]; ok {
				intersections = append(intersections, pos)
			}
			if _, ok := wire2Map[pos]; !ok {
				wire2Map[pos] = d
			}
			d++
		}
	}

	part1 := math.MaxInt
	part2 := math.MaxInt
	for _, pos := range intersections {
		manDist := utils.AbsInt(pos[0]) + utils.AbsInt(pos[1])
		if manDist < part1 {
			part1 = manDist
		}
		steps := wire1Map[pos] + wire2Map[pos]
		if steps < part2 {
			part2 = steps
		}
	}

	return part1, part2
}
