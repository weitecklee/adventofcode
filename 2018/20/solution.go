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

type QueueEntry struct {
	pos   [2]int
	doors int
}

var directions = [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// use stack to keep track of branching points.
// when `(` is encountered, a branch is started, add current pos to stack.
// when `|` is encountered, go back to start of branch, top of stack becomes current pos again.
// when `)` is encounted, branch is done, pop value off stack to become new current pos.

// facilityMap is map of <pos, isRoom>
// isRoom = !isDoor

func createMap(puzzleInput string) map[[2]int]bool {
	facilityMap := make(map[[2]int]bool)
	pos := [2]int{0, 0}
	facilityMap[pos] = true
	posStack := [][2]int{}
	for _, ch := range puzzleInput {
		switch ch {
		case 'N':
			pos[1]++
			facilityMap[pos] = false
			pos[1]++
			facilityMap[pos] = true
		case 'E':
			pos[0]++
			facilityMap[pos] = false
			pos[0]++
			facilityMap[pos] = true
		case 'W':
			pos[0]--
			facilityMap[pos] = false
			pos[0]--
			facilityMap[pos] = true
		case 'S':
			pos[1]--
			facilityMap[pos] = false
			pos[1]--
			facilityMap[pos] = true
		case '|':
			pos = posStack[len(posStack)-1]
		case '(':
			posStack = append(posStack, pos)
		case ')':
			pos = posStack[len(posStack)-1]
			posStack = posStack[:len(posStack)-1]
		}
	}
	return facilityMap
}

func solve(facilityMap map[[2]int]bool) (int, int) {
	queue := []QueueEntry{{[2]int{0, 0}, 0}}
	visited := make(map[[2]int]struct{})
	visited[[2]int{0, 0}] = struct{}{}
	part1 := 0
	part2 := 0
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		pos := curr.pos
		doors := curr.doors
		if facilityMap[pos] {
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
			if _, exists := visited[pos2]; exists {
				continue
			}
			visited[pos2] = struct{}{}
			queue = append(queue, QueueEntry{pos2, doors})
		}
	}
	return part1, part2
}
