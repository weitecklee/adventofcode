package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
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
	fmt.Println(part1(puzzleInput))
	fmt.Println(part2(puzzleInput))
}

var directionConversion = map[byte]byte{
	'0': 'R',
	'1': 'D',
	'2': 'L',
	'3': 'U',
	'R': '0',
	'D': '1',
	'L': '2',
	'U': '3',
}

type Step struct {
	direction byte
	distance  int
	color     string
}

func (s *Step) Swap() *Step {
	distance, err := strconv.ParseInt(s.color[:5], 16, 0)
	if err != nil {
		panic(err)
	}
	color := fmt.Sprintf("%05x%c", s.distance, directionConversion[s.direction])
	return &Step{directionConversion[s.color[5]], int(distance), color}
}

func parseInput(data []string) []*Step {
	steps := make([]*Step, len(data))
	stepRegex := regexp.MustCompile(`(\w) (\d+) \(#(\w{6})\)`)
	for i, line := range data {
		match := stepRegex.FindStringSubmatch(line)
		if match == nil {
			panic(fmt.Sprintf("Could not match regex with: %s", line))
		}
		distance, err := strconv.Atoi(match[2])
		if err != nil {
			panic(err)
		}
		steps[i] = &Step{match[1][0], distance, match[3]}
	}
	return steps
}

func calculateArea(steps []*Step) int {
	// Combination of Shoelace formula and Pick's theorem
	// Shoelace formula calculates area of polygon given vertex coordinates
	// Pick's theorem calculates area of polygon given number of boundary points and number of interior points
	// We have vertices (and boundary points) from input, and ultimately want to find sum of interior points and boundary points
	// Shoelace formula
	// https://en.wikipedia.org/wiki/Shoelace_formula
	// Pick's theorem
	// https://en.wikipedia.org/wiki/Pick%27s_theorem

	var x0, y0, x1, y1, sum, nBoundaryPoints int
	// `x0`, `y0`, `x1`, `y1` are vertex coordinates x_i, y_i, x_i+1, y_i+1,
	// 		with (x0, y0) starting from origin
	// `sum` is sum of (x_i * y_i+1 - x_i+1 * y_i) terms from Shoelace formula
	// `nBoundaryPoints` is count of boundary points

	for _, step := range steps {
		x1, y1 = x0, y0
		switch step.direction {
		case 'U':
			y1 += step.distance
		case 'D':
			y1 -= step.distance
		case 'L':
			x1 -= step.distance
		case 'R':
			x1 += step.distance
		}
		nBoundaryPoints += step.distance
		sum += x0*y1 - x1*y0
		x0, y0 = x1, y1
	}
	area := utils.AbsInt(sum) / 2                   // from Shoelance formula
	nInteriorPoints := area - nBoundaryPoints/2 + 1 // Pick's theorem rearranged
	return nInteriorPoints + nBoundaryPoints
}

func part1(puzzleInput []*Step) int {
	return calculateArea(puzzleInput)
}

func part2(puzzleInput []*Step) int {
	swappedSteps := make([]*Step, len(puzzleInput))
	for i, step := range puzzleInput {
		swappedSteps[i] = step.Swap()
	}
	return calculateArea(swappedSteps)
}
