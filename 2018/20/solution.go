package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	_, filename, _, _ := runtime.Caller(0)
	dirname := filepath.Dir(filename)
	inputFilePath := filepath.Join(dirname, "input.txt")
	data, err := os.ReadFile(inputFilePath)
	if err != nil {
		panic(err)
	}
	puzzleInput := string(data)
	facilityMap := createMap(puzzleInput)
	fmt.Println(solve(facilityMap))
}

type Point struct {
	pos          [2]int
	isRoom       bool
	doorsToReach int
}

type QueueEntry struct {
	pos   [2]int
	doors int
}

var directions = [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// use stack to keep track of branching points.
// when `(` is encountered, a branch is started, add current pos to stack.
// when `|` is encountered, go back to start of branch, top of stack becomes current pos again.
// when `)` is encounted, branch is done, pop value off stack to become new current pos.

func createMap(puzzleInput string) map[[2]int]*Point {
	facilityMap := make(map[[2]int]*Point)
	curr := [2]int{0, 0}
	facilityMap[curr] = &Point{curr, true, 0}
	currStack := [][2]int{}
	for _, ch := range puzzleInput {
		switch ch {
		case 'N':
			curr[1]++
			facilityMap[curr] = &Point{curr, false, -1}
			curr[1]++
			facilityMap[curr] = &Point{curr, true, -1}
		case 'E':
			curr[0]++
			facilityMap[curr] = &Point{curr, false, -1}
			curr[0]++
			facilityMap[curr] = &Point{curr, true, -1}
		case 'W':
			curr[0]--
			facilityMap[curr] = &Point{curr, false, -1}
			curr[0]--
			facilityMap[curr] = &Point{curr, true, -1}
		case 'S':
			curr[1]--
			facilityMap[curr] = &Point{curr, false, -1}
			curr[1]--
			facilityMap[curr] = &Point{curr, true, -1}
		case '|':
			curr = currStack[len(currStack)-1]
		case '(':
			currStack = append(currStack, curr)
		case ')':
			curr = currStack[len(currStack)-1]
			currStack = currStack[:len(currStack)-1]
		}
	}
	return facilityMap
}

func solve(facilityMap map[[2]int]*Point) (int, int) {
	queue := []QueueEntry{{[2]int{0, 0}, 0}}
	part1 := 0
	part2 := 0
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		pos := curr.pos
		doors := curr.doors
		if facilityMap[pos].isRoom {
			if doors > part1 {
				part1 = doors
			}
			if doors >= 1000 {
				part2++
			}
		} else {
			doors++
		}
		for _, d := range directions {
			pos2 := [2]int{pos[0] + d[0], pos[1] + d[1]}
			if _, exists := facilityMap[pos2]; !exists {
				// any pos not in facilityMap is not room or door, i.e., it's a wall
				continue
			}
			if facilityMap[pos2].doorsToReach >= 0 {
				continue
			}
			facilityMap[pos2].doorsToReach = doors
			queue = append(queue, QueueEntry{pos2, doors})
		}
	}
	return part1, part2
}
