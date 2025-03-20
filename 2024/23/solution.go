package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
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
	computerMap := parseInput(strings.Split(string(data), "\n"))
	fmt.Println(part1(computerMap))
}

type Computer struct {
	name      string
	neighbors map[*Computer]struct{}
}

func NewComputer(name string) *Computer {
	return &Computer{name, make(map[*Computer]struct{})}
}

func LinkComputers(c1, c2 *Computer) {
	c1.neighbors[c2] = struct{}{}
	c2.neighbors[c1] = struct{}{}
}

func parseInput(data []string) map[string]*Computer {
	computerMap := make(map[string]*Computer)
	for _, line := range data {
		parts := strings.Split(line, "-")
		if _, ok := computerMap[parts[0]]; !ok {
			computerMap[parts[0]] = NewComputer(parts[0])
		}
		if _, ok := computerMap[parts[1]]; !ok {
			computerMap[parts[1]] = NewComputer(parts[1])
		}
		LinkComputers(computerMap[parts[0]], computerMap[parts[1]])
	}
	return computerMap
}

func part1(computerMap map[string]*Computer) int {
	groupsOfThree := make(map[[3]*Computer]struct{})
	for _, computer := range computerMap {
		for neighbor := range computer.neighbors {
			for third := range neighbor.neighbors {
				if _, ok := computer.neighbors[third]; ok {
					slice := []*Computer{computer, neighbor, third}
					sort.SliceStable(slice, func(i, j int) bool {
						return slice[i].name < slice[j].name
					})
					var group [3]*Computer
					copy(group[:], slice)
					groupsOfThree[group] = struct{}{}
				}
			}
		}
	}

	res := 0
	for group := range groupsOfThree {
		for _, computer := range group {
			if computer.name[0] == 't' {
				res++
				break
			}
		}
	}
	return res
}
