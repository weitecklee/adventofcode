package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n")
	lights := parseInput(input)
	fmt.Println(part1(lights))
}

func parseInput(input []string) map[[2]int]bool {
	lights := map[[2]int]bool{}
	for j, line := range input {
		for i, c := range line {
			if c == '#' {
				lights[[2]int{i, j}] = true
			}
		}
	}
	return lights
}

func part1(lights map[[2]int]bool) int {
	for i := 0; i < 100; i++ {
		neighbors := map[[2]int]int{}
		for coord := range lights {
			for j := -1; j <= 1; j++ {
				for k := -1; k <= 1; k++ {
					if (j == 0 && k == 0) || coord[0]+j < 0 || coord[0]+j >= 100 || coord[1]+k < 0 || coord[1]+k >= 100 {
						continue
					}
					coord2 := [2]int{coord[0] + j, coord[1] + k}
					neighbors[coord2]++
				}
			}
		}
		lights2 := map[[2]int]bool{}
		for coord := range lights {
			if neighbors[coord] == 2 || neighbors[coord] == 3 {
				lights2[coord] = true
			}
		}
		for coord, n := range neighbors {
			if n == 3 && !lights[coord] {
				lights2[coord] = true
			}
		}
		lights = lights2
	}
	return len(lights)
}
