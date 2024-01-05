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
	input := strings.Split(string(data), "\n")
	fmt.Println(part1(parseInput(input)))
	fmt.Println(part2(parseInput(input)))
}

type Elf struct {
	pos         [2]int
	proposedPos [2]int
	allAdj      int
	northAdj    int
	eastAdj     int
	southAdj    int
	westAdj     int
}

func (e *Elf) CountAdj(elfMap *map[[2]int]bool) {
	e.northAdj = 0
	e.eastAdj = 0
	e.southAdj = 0
	e.westAdj = 0
	e.allAdj = 0
	if (*elfMap)[[2]int{e.pos[0] + 1, e.pos[1] + 1}] {
		e.eastAdj++
		e.southAdj++
		e.allAdj++
	}
	if (*elfMap)[[2]int{e.pos[0], e.pos[1] + 1}] {
		e.southAdj++
		e.allAdj++
	}
	if (*elfMap)[[2]int{e.pos[0] - 1, e.pos[1] + 1}] {
		e.westAdj++
		e.southAdj++
		e.allAdj++
	}
	if (*elfMap)[[2]int{e.pos[0] - 1, e.pos[1]}] {
		e.westAdj++
		e.allAdj++
	}
	if (*elfMap)[[2]int{e.pos[0] - 1, e.pos[1] - 1}] {
		e.westAdj++
		e.northAdj++
		e.allAdj++
	}
	if (*elfMap)[[2]int{e.pos[0], e.pos[1] - 1}] {
		e.northAdj++
		e.allAdj++
	}
	if (*elfMap)[[2]int{e.pos[0] + 1, e.pos[1] - 1}] {
		e.eastAdj++
		e.northAdj++
		e.allAdj++
	}
	if (*elfMap)[[2]int{e.pos[0] + 1, e.pos[1]}] {
		e.eastAdj++
		e.allAdj++
	}
}

func (e *Elf) Propose(dir int, elfMap *map[[2]int]bool) {
	e.CountAdj(elfMap)
	e.proposedPos = e.pos
	if e.allAdj == 0 {
		return
	}
	switch dir {
	case 0:
		if e.northAdj == 0 {
			e.proposedPos[1]--
		} else if e.southAdj == 0 {
			e.proposedPos[1]++
		} else if e.westAdj == 0 {
			e.proposedPos[0]--
		} else if e.eastAdj == 0 {
			e.proposedPos[0]++
		}
	case 1:
		if e.southAdj == 0 {
			e.proposedPos[1]++
		} else if e.westAdj == 0 {
			e.proposedPos[0]--
		} else if e.eastAdj == 0 {
			e.proposedPos[0]++
		} else if e.northAdj == 0 {
			e.proposedPos[1]--
		}
	case 2:
		if e.westAdj == 0 {
			e.proposedPos[0]--
		} else if e.eastAdj == 0 {
			e.proposedPos[0]++
		} else if e.northAdj == 0 {
			e.proposedPos[1]--
		} else if e.southAdj == 0 {
			e.proposedPos[1]++
		}
	case 3:
		if e.eastAdj == 0 {
			e.proposedPos[0]++
		} else if e.northAdj == 0 {
			e.proposedPos[1]--
		} else if e.southAdj == 0 {
			e.proposedPos[1]++
		} else if e.westAdj == 0 {
			e.proposedPos[0]--
		}
	}
}

func parseInput(input []string) (*[]*Elf, map[[2]int]bool) {
	elves := []*Elf{}
	elfMap := map[[2]int]bool{}
	for j, row := range input {
		for i, c := range row {
			if c == '#' {
				pos := [2]int{i, j}
				elves = append(elves, &Elf{
					pos: pos,
				})
				elfMap[pos] = true
			}
		}
	}
	return &elves, elfMap
}

func part1(elves *[]*Elf, elfMap map[[2]int]bool) int {
	dir := 0
	for rounds := 0; rounds < 10; rounds++ {
		proposedMap := map[[2]int]int{}
		for _, elf := range *elves {
			elf.Propose(dir, &elfMap)
			proposedMap[elf.proposedPos]++
		}
		elfMap2 := map[[2]int]bool{}
		for _, elf := range *elves {
			if proposedMap[elf.proposedPos] == 1 {
				elf.pos = elf.proposedPos
			}
			elfMap2[elf.pos] = true
		}
		elfMap = elfMap2
		dir++
		if dir == 4 {
			dir = 0
		}
	}
	xMin, yMin, xMax, yMax := math.MaxInt, math.MaxInt, math.MinInt, math.MinInt
	for _, elf := range *elves {
		if elf.pos[0] < xMin {
			xMin = elf.pos[0]
		}
		if elf.pos[0] > xMax {
			xMax = elf.pos[0]
		}
		if elf.pos[1] < yMin {
			yMin = elf.pos[1]
		}
		if elf.pos[1] > yMax {
			yMax = elf.pos[1]
		}
	}
	return int((math.Abs(float64(yMax-yMin))+1)*(math.Abs(float64(xMax-xMin))+1)) - len(*elves)
}

func part2(elves *[]*Elf, elfMap map[[2]int]bool) int {
	dir := 0
	rounds := 0
	for {
		rounds++
		noMoves := 0
		proposedMap := map[[2]int]int{}
		for _, elf := range *elves {
			elf.Propose(dir, &elfMap)
			if elf.allAdj == 0 {
				noMoves++
			}
			proposedMap[elf.proposedPos]++
		}
		if noMoves == len(*elves) {
			break
		}
		elfMap2 := map[[2]int]bool{}
		for _, elf := range *elves {
			if proposedMap[elf.proposedPos] == 1 {
				elf.pos = elf.proposedPos
			}
			elfMap2[elf.pos] = true
		}
		elfMap = elfMap2
		dir++
		if dir == 4 {
			dir = 0
		}
	}
	return rounds
}
