package main

import (
	"fmt"
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
	puzzleInput := strings.Split(string(data), "\n")
	fmt.Println(part1(parseInput(puzzleInput)))
	fmt.Println(part2(parseInput(puzzleInput)))
}

type Octopus struct {
	energyLevel int
	flashed     bool
}

var directions = [][2]int{{-1, 0}, {-1, -1}, {-1, 1}, {0, -1}, {0, 1}, {1, 0}, {1, -1}, {1, 1}}

func parseInput(data []string) [][]*Octopus {
	octopusArray := make([][]*Octopus, len(data))
	for i, line := range data {
		row := make([]*Octopus, len(line))
		for j, s := range line {
			n := int(s - '0')
			row[j] = &Octopus{n, false}
		}
		octopusArray[i] = row
	}
	return octopusArray
}

func simulate(octopusArray [][]*Octopus) int {
	var flashQueue [][2]int
	for r, row := range octopusArray {
		for c, octopus := range row {
			octopus.energyLevel++
			if octopus.energyLevel > 9 {
				flashQueue = append(flashQueue, [2]int{r, c})
				octopus.flashed = true
			}
		}
	}
	i := 0
	for i < len(flashQueue) {
		r, c := flashQueue[i][0], flashQueue[i][1]
		for _, d := range directions {
			r2, c2 := r+d[0], c+d[1]
			if r2 < 0 || c2 < 0 || r2 >= len(octopusArray) || c2 >= len(octopusArray[0]) {
				continue
			}
			octopusArray[r2][c2].energyLevel++
			if octopusArray[r2][c2].energyLevel > 9 && !octopusArray[r2][c2].flashed {
				flashQueue = append(flashQueue, [2]int{r2, c2})
				octopusArray[r2][c2].flashed = true
			}
		}
		i++
	}
	for _, pos := range flashQueue {
		r, c := pos[0], pos[1]
		octopusArray[r][c].energyLevel = 0
		octopusArray[r][c].flashed = false
	}
	return len(flashQueue)
}

func part1(octopusArray [][]*Octopus) int {
	res := 0
	for range 100 {
		res += simulate(octopusArray)
	}
	return res
}

func part2(octopusArray [][]*Octopus) int {
	i := 1
	nOctopuses := len(octopusArray) * len(octopusArray[0])
	for nOctopuses != simulate(octopusArray) {
		i++
	}
	return i
}
