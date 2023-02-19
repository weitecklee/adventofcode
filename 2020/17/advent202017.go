package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input202017.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n")
	parsedInput := parseInput(input)
	fmt.Println(part1(parsedInput))
	fmt.Println(part2(parsedInput))
}

func convertCoords(x int, y int, z int, w int) string {
	return strconv.Itoa(x) + "," + strconv.Itoa(y) + "," + strconv.Itoa(z) + "," + strconv.Itoa(w)
}

func parseInput(input []string) map[string]bool {
	cubes := map[string]bool{}
	for i, row := range input {
		for j, c := range row {
			if string(c) == "#" {
				cubes[convertCoords(i, j, 0, 0)] = true
			}
		}
	}
	return cubes
}

func coordsOfInterest(cubes map[string]bool, forPart2 bool) (map[string]bool, map[string]bool) {
	neighbors := map[string]int{}
	twoActiveNeighbors := map[string]bool{}
	threeActiveNeighbors := map[string]bool{}
	for coords := range cubes {
		splitCoords := strings.Split(coords, ",")
		x, _ := strconv.Atoi(splitCoords[0])
		y, _ := strconv.Atoi(splitCoords[1])
		z, _ := strconv.Atoi(splitCoords[2])
		w, _ := strconv.Atoi(splitCoords[3])
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				for k := -1; k <= 1; k++ {
					if forPart2 {
						for h := -1; h <= 1; h++ {
							if i == 0 && j == 0 && k == 0 && h == 0 {
								continue
							}
							x2 := x + i
							y2 := y + j
							z2 := z + k
							w2 := w + h
							convCoords := convertCoords(x2, y2, z2, w2)
							neighbors[convCoords]++
							if neighbors[convCoords] == 2 {
								twoActiveNeighbors[convCoords] = true
							} else if neighbors[convCoords] == 3 {
								threeActiveNeighbors[convCoords] = true
								delete(twoActiveNeighbors, convCoords)
							} else if neighbors[convCoords] == 4 {
								delete(threeActiveNeighbors, convCoords)
							}
						}
					} else {
						if i == 0 && j == 0 && k == 0 {
							continue
						}
						x2 := x + i
						y2 := y + j
						z2 := z + k
						convCoords := convertCoords(x2, y2, z2, 0)
						neighbors[convCoords]++
						if neighbors[convCoords] == 2 {
							twoActiveNeighbors[convCoords] = true
						} else if neighbors[convCoords] == 3 {
							threeActiveNeighbors[convCoords] = true
							delete(twoActiveNeighbors, convCoords)
						} else if neighbors[convCoords] == 4 {
							delete(threeActiveNeighbors, convCoords)
						}
					}
				}
			}
		}
	}
	return twoActiveNeighbors, threeActiveNeighbors
}

func simulate(cubes map[string]bool, forPart2 bool) map[string]bool {
	cubes2 := map[string]bool{}
	twoActiveNeighbors, threeActiveNeighbors := coordsOfInterest(cubes, forPart2)
	for coords := range cubes {
		if twoActiveNeighbors[coords] || threeActiveNeighbors[coords] {
			cubes2[coords] = true
		}
	}
	for coords := range threeActiveNeighbors {
		cubes2[coords] = true
	}
	return cubes2
}

func part1(input map[string]bool) int {
	cubes := input
	i := 0
	for i < 6 {
		cubes = simulate(cubes, false)
		i++
	}
	return len(cubes)
}

func part2(input map[string]bool) int {
	cubes := input
	i := 0
	for i < 6 {
		cubes = simulate(cubes, true)
		i++
	}
	return len(cubes)
}
