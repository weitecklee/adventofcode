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
	puzzleInput := parseInput(strings.Split(string(data), "\n"))
	fmt.Println(part1(puzzleInput))
	fmt.Println(part2(puzzleInput))
}

type Orbiter struct {
	parent   *Orbiter
	children []*Orbiter
}

type QueueEntry struct {
	orbiter *Orbiter
	depth   int
}

func parseInput(data []string) map[string]*Orbiter {
	orbiterMap := make(map[string]*Orbiter)
	var name1, name2 string
	var obj1, obj2 *Orbiter
	for _, line := range data {
		objects := strings.Split(line, ")")
		name1 = objects[0]
		name2 = objects[1]
		if _, ok := orbiterMap[name1]; !ok {
			orbiterMap[name1] = &Orbiter{nil, []*Orbiter{}}
		}
		if _, ok := orbiterMap[name2]; !ok {
			orbiterMap[name2] = &Orbiter{nil, []*Orbiter{}}
		}
		obj1 = orbiterMap[name1]
		obj2 = orbiterMap[name2]
		obj1.children = append(obj1.children, obj2)
		obj2.parent = obj1
	}
	return orbiterMap
}

func part1(orbiterMap map[string]*Orbiter) int {
	res := 0
	var queue []QueueEntry
	queue = append(queue, QueueEntry{orbiterMap["COM"], 0})
	for len(queue) > 0 {
		entry := queue[0]
		queue = queue[1:]
		orbiter := entry.orbiter
		depth := entry.depth
		res += depth
		for _, child := range orbiter.children {
			queue = append(queue, QueueEntry{child, depth + 1})
		}
	}
	return res
}

func part2(orbiterMap map[string]*Orbiter) int {
	visited := make(map[*Orbiter]struct{})
	var queue []QueueEntry
	queue = append(queue, QueueEntry{orbiterMap["YOU"], 0})
	visited[orbiterMap["YOU"]] = struct{}{}
	for len(queue) > 0 {
		entry := queue[0]
		queue = queue[1:]
		orbiter := entry.orbiter
		depth := entry.depth
		if orbiter == orbiterMap["SAN"] {
			return depth - 2
		}
		for _, child := range orbiter.children {
			if _, ok := visited[child]; ok {
				continue
			}
			visited[child] = struct{}{}
			queue = append(queue, QueueEntry{child, depth + 1})
		}
		if orbiter.parent == nil {
			continue
		}
		if _, ok := visited[orbiter.parent]; !ok {
			visited[orbiter.parent] = struct{}{}
			queue = append(queue, QueueEntry{orbiter.parent, depth + 1})
		}
	}
	return -1
}
