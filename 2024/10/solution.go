package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

var (
	directions = [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
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
	trailMap := parseInput(&puzzleInput)
	fmt.Println(solve(trailMap))
}

func parseInput(puzzleInput *[]string) *[][]int {
	trailMap := make([][]int, len(*puzzleInput))
	for r, row := range *puzzleInput {
		for _, chr := range row {
			if n, err := strconv.Atoi(string(chr)); err == nil {
				trailMap[r] = append(trailMap[r], n)
			}
		}
	}
	return &trailMap
}

func findTrailheads(trailMap *[][]int) *[][2]int {
	var trailheads [][2]int
	for r := range *trailMap {
		for c, n := range (*trailMap)[r] {
			if n == 0 {
				trailheads = append(trailheads, [2]int{r, c})
			}
		}
	}
	return &trailheads
}

func calcNumTrails(trailMap *[][]int, start *[2]int, isPart2 bool) int {
	var queue [][2]int
	visited := make(map[[2]int]struct{})
	res := 0
	queue = append(queue, *start)
	for len(queue) > 0 {
		pos := queue[0]
		r, c := pos[0], pos[1]
		queue = queue[1:]
		for _, dir := range directions {
			r2, c2 := pos[0]+dir[0], pos[1]+dir[1]
			if r2 < 0 || c2 < 0 || r2 > len(*trailMap)-1 || c2 > len((*trailMap)[0])-1 {
				continue
			}
			if (*trailMap)[r2][c2]-(*trailMap)[r][c] != 1 {
				continue
			}
			if _, ok := visited[[2]int{r2, c2}]; ok && !isPart2 {
				continue
			}
			visited[[2]int{r2, c2}] = struct{}{}
			if (*trailMap)[r2][c2] == 9 {
				res++
				continue
			}
			queue = append(queue, [2]int{r2, c2})
		}
	}
	return res
}

func solve(trailMap *[][]int) (int, int) {
	trailheads := findTrailheads(trailMap)
	part1 := 0
	part2 := 0
	for _, head := range *trailheads {
		part1 += calcNumTrails(trailMap, &head, false)
		part2 += calcNumTrails(trailMap, &head, true)
	}
	return part1, part2
}
