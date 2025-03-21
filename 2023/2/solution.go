package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
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
	games := parseInput(strings.Split(string(data), "\n"))
	fmt.Println(part1(games))
	fmt.Println(part2(games))
}

var (
	numRegex  = regexp.MustCompile(`\d+`)
	cubeRegex = regexp.MustCompile(`(\d+) (\w+)`)
	colors    = []string{"red", "green", "blue"}
)

type Game struct {
	id            int
	revealedCubes []map[string]int
}

func (g *Game) IsPossible(testBag map[string]int) bool {
	for _, cubeSet := range g.revealedCubes {
		for col, n := range testBag {
			if cubeSet[col] > n {
				return false
			}
		}
	}
	return true
}

func (g *Game) MinimumSet() map[string]int {
	minSet := make(map[string]int)
	for _, cubeSet := range g.revealedCubes {
		for col, n := range cubeSet {
			if n > minSet[col] {
				minSet[col] = n
			}
		}
	}
	return minSet
}

func PowerOfSet(set map[string]int) int {
	res := 1
	for _, col := range colors {
		res *= set[col]
	}
	return res
}

func NewGame(line string) *Game {
	parts := strings.Split(line, ":")
	cubeSets := strings.Split(parts[1], ";")
	revealedCubes := make([]map[string]int, len(cubeSets))
	match := numRegex.FindString(parts[0])
	id, err := strconv.Atoi(match)
	if err != nil {
		panic(err)
	}
	for i, set := range cubeSets {
		matches := cubeRegex.FindAllStringSubmatch(set, -1)
		cubes := make(map[string]int, len(matches))
		for _, match := range matches {
			n, err := strconv.Atoi(match[1])
			if err != nil {
				panic(err)
			}
			cubes[match[2]] = n
		}
		revealedCubes[i] = cubes
	}
	return &Game{id, revealedCubes}
}

func parseInput(data []string) []*Game {
	games := make([]*Game, len(data))
	for i, line := range data {
		games[i] = NewGame(line)
	}
	return games
}

func part1(games []*Game) int {
	res := 0
	testBag := map[string]int{"red": 12, "green": 13, "blue": 14}
	for _, game := range games {
		if game.IsPossible(testBag) {
			res += game.id
		}
	}
	return res
}

func part2(games []*Game) int {
	res := 0
	for _, game := range games {
		res += PowerOfSet(game.MinimumSet())
	}
	return res
}
