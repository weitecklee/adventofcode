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
	fmt.Println(part2(computerMap))
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

func isSupersetOf[T comparable](mapA, mapB map[T]struct{}) bool {
	for k := range mapB {
		if _, ok := mapA[k]; !ok {
			return false
		}
	}
	return true
}

func part2(computerMap map[string]*Computer) string {
	largestSet := make(map[*Computer]struct{})
	computers := make([]*Computer, 0, len(computerMap))
	for _, computer := range computerMap {
		computers = append(computers, computer)
	}
	for i, computer := range computers {
		checked := make(map[*Computer]struct{})
		for j := i + 1; j < len(computers); j++ {
			if _, ok := computer.neighbors[computers[j]]; !ok {
				continue
			}
			if _, ok := checked[computers[j]]; ok {
				continue
			}
			checked[computers[j]] = struct{}{}
			network := make(map[*Computer]struct{})
			network[computer] = struct{}{}
			network[computers[j]] = struct{}{}
			for k := j + 1; k < len(computers); k++ {
				if isSupersetOf(computers[k].neighbors, network) {
					checked[computers[k]] = struct{}{}
					network[computers[k]] = struct{}{}
				}
			}
			if len(network) > len(largestSet) {
				largestSet = network
			}
		}

	}
	names := make([]string, 0, len(largestSet))
	for computer := range largestSet {
		names = append(names, computer.name)
	}
	sort.Strings(names)
	return strings.Join(names, ",")
}
