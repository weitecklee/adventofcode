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
	fmt.Println(part1(&puzzleInput))
	fmt.Println(part2(&puzzleInput))
}

func calcEnergizedTiles(puzzleInput *[]string, startingBeam [4]int) int {
	energized := make(map[[2]int]struct{})
	visited := make(map[[4]int]struct{})
	beamQueue := [][4]int{startingBeam}
	var beam [4]int
	var r, c, dr, dc int
	var tile byte
	rMax := len(*puzzleInput) - 1
	cMax := len((*puzzleInput)[0]) - 1
	for len(beamQueue) > 0 {
		beam = beamQueue[0]
		beamQueue = beamQueue[1:]
		if _, ok := visited[beam]; ok {
			continue
		}
		visited[beam] = struct{}{}
		r = beam[0]
		c = beam[1]
		dr = beam[2]
		dc = beam[3]
		if r < 0 || c < 0 || r > rMax || c > cMax {
			continue
		}
		energized[[2]int{r, c}] = struct{}{}
		tile = (*puzzleInput)[r][c]
		switch tile {
		case '.':
			beamQueue = append(beamQueue, [4]int{r + dr, c + dc, dr, dc})
		case '\\':
			dr, dc = dc, dr
			beamQueue = append(beamQueue, [4]int{r + dr, c + dc, dr, dc})
		case '/':
			dr, dc = -dc, -dr
			beamQueue = append(beamQueue, [4]int{r + dr, c + dc, dr, dc})
		case '|':
			if dc == 0 {
				beamQueue = append(beamQueue, [4]int{r + dr, c + dc, dr, dc})
			} else {
				dc = 0
				beamQueue = append(beamQueue, [4]int{r + 1, c, 1, dc})
				beamQueue = append(beamQueue, [4]int{r - 1, c, -1, dc})
			}
		case '-':
			if dr == 0 {
				beamQueue = append(beamQueue, [4]int{r + dr, c + dc, dr, dc})
			} else {
				dr = 0
				beamQueue = append(beamQueue, [4]int{r + dr, c + 1, dr, 1})
				beamQueue = append(beamQueue, [4]int{r + dr, c - 1, dr, -1})
			}
		default:
			panic(fmt.Sprintf("Unknown tile: %c", tile))
		}
	}
	return len(energized)
}

func part1(puzzleInput *[]string) int {
	return calcEnergizedTiles(puzzleInput, [4]int{0, 0, 0, 1})
}

func part2(puzzleInput *[]string) int {
	var tiles, maxTiles int
	rMax := len(*puzzleInput) - 1
	cMax := len((*puzzleInput)[0]) - 1
	for r := range rMax + 1 {
		tiles = calcEnergizedTiles(puzzleInput, [4]int{r, 0, 0, 1})
		if tiles > maxTiles {
			maxTiles = tiles
		}
		tiles = calcEnergizedTiles(puzzleInput, [4]int{r, cMax, 0, -1})
		if tiles > maxTiles {
			maxTiles = tiles
		}
	}
	for c := range cMax + 1 {
		tiles = calcEnergizedTiles(puzzleInput, [4]int{0, c, 1, 0})
		if tiles > maxTiles {
			maxTiles = tiles
		}
		tiles = calcEnergizedTiles(puzzleInput, [4]int{rMax, c, -1, 0})
		if tiles > maxTiles {
			maxTiles = tiles
		}
	}
	return maxTiles
}
