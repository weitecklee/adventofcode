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
	fmt.Println(part1(parseInput(input)))
}

type Carrier struct {
	pos  [2]int
	dir  [2]int
	grid *map[string]bool
}

func (c *Carrier) burst() bool {
	if (*c.grid)[coord(c.pos)] {
		if c.dir[0] == 0 {
			c.dir[0] = -c.dir[1]
			c.dir[1] = 0
		} else {
			c.dir[1] = c.dir[0]
			c.dir[0] = 0
		}
		(*c.grid)[coord(c.pos)] = false
		c.pos[0] += c.dir[0]
		c.pos[1] += c.dir[1]
		return false
	} else {
		if c.dir[0] == 0 {
			c.dir[0] = c.dir[1]
			c.dir[1] = 0
		} else {
			c.dir[1] = -c.dir[0]
			c.dir[0] = 0
		}
		(*c.grid)[coord(c.pos)] = true
		c.pos[0] += c.dir[0]
		c.pos[1] += c.dir[1]
		return true
	}
}

func coord(pos [2]int) string {
	return fmt.Sprintf("%d,%d", pos[0], pos[1])
}

func parseInput(input []string) (*map[string]bool, [2]int) {
	grid := map[string]bool{}
	center := [2]int{}
	center[0] = len(input[0]) / 2
	center[1] = len(input) / 2
	for i := 0; i < len(input[0]); i++ {
		for j := 0; j < len(input); j++ {
			if input[j][i] == '#' {
				grid[coord([2]int{i, j})] = true
			}
		}
	}
	return &grid, center
}

func part1(grid *map[string]bool, center [2]int) int {
	virus := Carrier{
		pos:  center,
		dir:  [2]int{0, -1},
		grid: grid,
	}
	count := 0
	for i := 0; i < 10000; i++ {
		if virus.burst() {
			count++
		}
	}
	return count
}
