package main

import (
	"container/heap"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/weitecklee/adventofcode/utils"
)

var (
	directions = [][2]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
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
	startTile, endTile := parseInput(puzzleInput)
	fmt.Println(solve(puzzleInput, startTile, endTile))
}

func parseInput(puzzleInput []string) ([2]int, [2]int) {
	var startTile, endTile [2]int
	for r, row := range puzzleInput {
		for c, chr := range row {
			if chr == 'S' {
				startTile = [2]int{r, c}
			} else if chr == 'E' {
				endTile = [2]int{r, c}
			}
		}
		if startTile != [2]int{0, 0} && endTile != [2]int{0, 0} {
			break
		}
	}
	return startTile, endTile
}

func dist(endTile, p [2]int) int {
	return utils.AbsInt(endTile[0]-p[0]) + utils.AbsInt((endTile[1] - p[1]))
}

type Value struct {
	score    int
	pos      [2]int
	dirIndex int
	visited  map[[2]int]struct{}
}

func solve(puzzleInput []string, startTile, endTile [2]int) (int, int) {
	queue := utils.NewMinHeap[Value]()
	visited := make(map[[2]int]struct{})
	visited[startTile] = struct{}{}
	heap.Push(queue, &utils.Item[Value]{Value: Value{0, startTile, 0, visited}, Priority: 0})
	visitedScores := make(map[[3]int]int)
	visitedScores[[3]int{startTile[0], startTile[1], 0}] = 0
	part1 := 0
	var part2 map[[2]int]struct{}
	for len(queue.PriorityQueue) > 0 {
		item := heap.Pop(queue).(*utils.Item[Value])
		score := item.Value.score
		pos := item.Value.pos
		if pos == endTile {
			if part1 == 0 {
				part1 = score
				part2 = item.Value.visited
			} else if score == part1 {
				for k := range item.Value.visited {
					part2[k] = struct{}{}
				}
			} else {
				break
			}
			continue
		}
		dirIndex := item.Value.dirIndex
		visited := item.Value.visited
		for i, dir := range directions {
			if i == 3-dirIndex {
				continue
			}
			pos2 := [2]int{pos[0] + dir[0], pos[1] + dir[1]}

			if puzzleInput[pos2[0]][pos2[1]] == '#' {
				continue
			}
			score2 := score + 1
			if i != dirIndex {
				score2 += 1000
			}
			if v, ok := visitedScores[[3]int{pos2[0], pos2[1], i}]; ok && v < score2 {
				continue
			}
			visitedScores[[3]int{pos2[0], pos2[1], i}] = score2
			visited2 := make(map[[2]int]struct{})
			for k := range visited {
				visited2[k] = struct{}{}
			}
			visited2[pos2] = struct{}{}
			heap.Push(queue, &utils.Item[Value]{Value: Value{score2, pos2, i, visited2}, Priority: score2 + dist(endTile, pos2)})
		}
	}

	return part1, len(part2)
}
