package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

func main() {
	_, filename, _, _ := runtime.Caller(0)
	dirname := filepath.Dir(filename)
	inputFilePath := filepath.Join(dirname, "input.txt")
	data, err := os.ReadFile(inputFilePath)
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n")
	fmt.Println(part1(parseInput(input)))
	fmt.Println(part1v2(parseInputv2(input)))
	fmt.Println(part2(parseInput2(input)))
	fmt.Println(part2v2(parseInput2v2(input)))
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	fmt.Printf("%v: %v\n", msg, time.Since(start))
}

type Carrierv2 struct {
	pos  [2]int
	dir  [2]int
	grid *map[[2]int]bool
}

func (c *Carrierv2) burst() bool {
	if (*c.grid)[c.pos] {
		if c.dir[0] == 0 {
			c.dir[0] = -c.dir[1]
			c.dir[1] = 0
		} else {
			c.dir[1] = c.dir[0]
			c.dir[0] = 0
		}
		(*c.grid)[c.pos] = false
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
		(*c.grid)[c.pos] = true
		c.pos[0] += c.dir[0]
		c.pos[1] += c.dir[1]
		return true
	}
}

func parseInputv2(input []string) (*map[[2]int]bool, [2]int) {
	grid := map[[2]int]bool{}
	center := [2]int{}
	center[0] = len(input[0]) / 2
	center[1] = len(input) / 2
	for i := 0; i < len(input[0]); i++ {
		for j := 0; j < len(input); j++ {
			if input[j][i] == '#' {
				grid[[2]int{i, j}] = true
			}
		}
	}
	return &grid, center
}

func part1v2(grid *map[[2]int]bool, center [2]int) int {
	defer duration(track("part1v2"))
	virus := Carrierv2{
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

type Carrier2v2 struct {
	pos  [2]int
	dir  [2]int
	grid *map[[2]int]string
}

func (c *Carrier2v2) burst() bool {
	infection := false
	switch (*c.grid)[c.pos] {
	case "":
		fallthrough
	case "clean":
		if c.dir[0] == 0 {
			c.dir[0] = c.dir[1]
			c.dir[1] = 0
		} else {
			c.dir[1] = -c.dir[0]
			c.dir[0] = 0
		}
		(*c.grid)[c.pos] = "weakened"
	case "weakened":
		(*c.grid)[c.pos] = "infected"
		infection = true
	case "flagged":
		c.dir[0] = -c.dir[0]
		c.dir[1] = -c.dir[1]
		(*c.grid)[c.pos] = "clean"
	case "infected":
		if c.dir[0] == 0 {
			c.dir[0] = -c.dir[1]
			c.dir[1] = 0
		} else {
			c.dir[1] = c.dir[0]
			c.dir[0] = 0
		}
		(*c.grid)[c.pos] = "flagged"
	}
	c.pos[0] += c.dir[0]
	c.pos[1] += c.dir[1]
	return infection
}

func parseInput2v2(input []string) (*map[[2]int]string, [2]int) {
	grid := map[[2]int]string{}
	center := [2]int{}
	center[0] = len(input[0]) / 2
	center[1] = len(input) / 2
	for i := 0; i < len(input[0]); i++ {
		for j := 0; j < len(input); j++ {
			if input[j][i] == '#' {
				grid[[2]int{i, j}] = "infected"
			}
		}
	}
	return &grid, center
}

func part2v2(grid *map[[2]int]string, center [2]int) int {
	defer duration(track("part2v2"))
	virus := Carrier2v2{
		pos:  center,
		dir:  [2]int{0, -1},
		grid: grid,
	}
	count := 0
	for i := 0; i < 10000000; i++ {
		if virus.burst() {
			count++
		}
	}
	return count
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
	defer duration(track("part1"))
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

type Carrier2 struct {
	pos  [2]int
	dir  [2]int
	grid *map[string]string
}

func (c *Carrier2) burst() bool {
	infection := false
	switch (*c.grid)[coord(c.pos)] {
	case "":
		fallthrough
	case "clean":
		if c.dir[0] == 0 {
			c.dir[0] = c.dir[1]
			c.dir[1] = 0
		} else {
			c.dir[1] = -c.dir[0]
			c.dir[0] = 0
		}
		(*c.grid)[coord(c.pos)] = "weakened"
	case "weakened":
		(*c.grid)[coord(c.pos)] = "infected"
		infection = true
	case "flagged":
		c.dir[0] = -c.dir[0]
		c.dir[1] = -c.dir[1]
		(*c.grid)[coord(c.pos)] = "clean"
	case "infected":
		if c.dir[0] == 0 {
			c.dir[0] = -c.dir[1]
			c.dir[1] = 0
		} else {
			c.dir[1] = c.dir[0]
			c.dir[0] = 0
		}
		(*c.grid)[coord(c.pos)] = "flagged"
	}
	c.pos[0] += c.dir[0]
	c.pos[1] += c.dir[1]
	return infection
}

func parseInput2(input []string) (*map[string]string, [2]int) {
	grid := map[string]string{}
	center := [2]int{}
	center[0] = len(input[0]) / 2
	center[1] = len(input) / 2
	for i := 0; i < len(input[0]); i++ {
		for j := 0; j < len(input); j++ {
			if input[j][i] == '#' {
				grid[coord([2]int{i, j})] = "infected"
			}
		}
	}
	return &grid, center
}

func part2(grid *map[string]string, center [2]int) int {
	defer duration(track("part2"))
	virus := Carrier2{
		pos:  center,
		dir:  [2]int{0, -1},
		grid: grid,
	}
	count := 0
	for i := 0; i < 10000000; i++ {
		if virus.burst() {
			count++
		}
	}
	return count
}
