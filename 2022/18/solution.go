package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"runtime"
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
	cubeSet := parseInput(strings.Split(string(data), "\n"))
	fmt.Println(part1(cubeSet))
	fmt.Println(part2(cubeSet))
}

var directions = [][3]int{{-1, 0, 0}, {1, 0, 0}, {0, -1, 0}, {0, 1, 0}, {0, 0, -1}, {0, 0, 1}}

func parseInput(data []string) map[[3]int]struct{} {
	cubeSet := make(map[[3]int]struct{}, len(data))
	for _, line := range data {
		parts := utils.ExtractInts(line)
		cubeSet[[3]int{parts[0], parts[1], parts[2]}] = struct{}{}
	}
	return cubeSet
}

func part1(cubeSet map[[3]int]struct{}) int {
	res := 0
	for cube := range cubeSet {
		for _, d := range directions {
			tmp := [3]int{cube[0] + d[0], cube[1] + d[1], cube[2] + d[2]}
			if _, exists := cubeSet[tmp]; !exists {
				res++
			}
		}
	}
	return res
}

// Use outer bounds to find surrounding air, add to area when droplet cube is encountered

func part2(cubeSet map[[3]int]struct{}) int {
	xMin := math.MaxInt
	xMax := math.MinInt
	yMin := math.MaxInt
	yMax := math.MinInt
	zMin := math.MaxInt
	zMax := math.MinInt
	for cube := range cubeSet {
		if cube[0] < xMin {
			xMin = cube[0]
		}
		if cube[0] > xMax {
			xMax = cube[0]
		}

		if cube[1] < yMin {
			yMin = cube[1]
		}
		if cube[1] > yMax {
			yMax = cube[1]
		}
		if cube[2] < zMin {
			zMin = cube[2]
		}
		if cube[2] > zMax {
			zMax = cube[2]
		}
	}
	res := 0
	queue := [][3]int{{xMin - 1, yMin - 1, zMin - 1}}
	checked := make(map[[3]int]struct{})
	checked[queue[0]] = struct{}{}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		toCheck := make([][3]int, 0, 6)
		if curr[0] > xMin-1 {
			toCheck = append(toCheck, [3]int{curr[0] - 1, curr[1], curr[2]})
		}
		if curr[0] < xMax+1 {
			toCheck = append(toCheck, [3]int{curr[0] + 1, curr[1], curr[2]})
		}
		if curr[1] > yMin-1 {
			toCheck = append(toCheck, [3]int{curr[0], curr[1] - 1, curr[2]})
		}
		if curr[1] < yMax+1 {
			toCheck = append(toCheck, [3]int{curr[0], curr[1] + 1, curr[2]})
		}
		if curr[2] > zMin-1 {
			toCheck = append(toCheck, [3]int{curr[0], curr[1], curr[2] - 1})
		}
		if curr[2] < zMax+1 {
			toCheck = append(toCheck, [3]int{curr[0], curr[1], curr[2] + 1})
		}
		for _, cube := range toCheck {
			if _, exists := cubeSet[cube]; exists {
				res++
			} else if _, exists := checked[cube]; !exists {
				queue = append(queue, cube)
				checked[cube] = struct{}{}
			}
		}
	}
	return res
}
