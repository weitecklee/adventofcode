package main

import (
	"container/heap"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/weitecklee/adventofcode/utils"
)

const (
	XMIN = 0
	YMIN = 0
	XMAX = 70
	YMAX = 70
)

var (
	directions = [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
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
	fmt.Println(part1(puzzleInput))
	fmt.Println(part2(puzzleInput))
}

func createMaze(puzzleInput []string, nBytes int) map[[2]int]struct{} {
	maze := make(map[[2]int]struct{})
	for i := 0; i < nBytes; i++ {
		nums := strings.Split(puzzleInput[i], ",")
		x, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}
		maze[[2]int{x, y}] = struct{}{}
	}

	return maze
}

type Value struct {
	steps int
	pos   [2]int
}

func heuristic(a [2]int) int {
	return utils.AbsInt(a[0]-XMAX) + utils.AbsInt(a[1]-YMAX)
}

func solve(puzzleInput []string, nBytes int) int {
	maze := createMaze(puzzleInput, nBytes)
	visited := make(map[[2]int]int)
	visited[[2]int{0, 0}] = 0
	queue := utils.NewMinHeap[Value]()
	queue.Push(&utils.Item[Value]{Value: Value{0, [2]int{0, 0}}, Priority: 0})
	for len(queue.PriorityQueue) > 0 {
		item := heap.Pop(queue).(*utils.Item[Value])
		value := item.Value
		if value.pos == [2]int{XMAX, YMAX} {
			return value.steps
		}
		for _, dir := range directions {
			newPos := [2]int{value.pos[0] + dir[0], value.pos[1] + dir[1]}
			if newPos[0] < XMIN || newPos[0] > XMAX || newPos[1] < YMIN || newPos[1] > YMAX {
				continue
			}
			if _, ok := maze[newPos]; ok {
				continue
			}
			if v, ok := visited[newPos]; ok && v <= value.steps+1 {
				continue
			}
			visited[newPos] = value.steps + 1
			heap.Push(queue, &utils.Item[Value]{Value: Value{value.steps + 1, newPos}, Priority: value.steps + 1 + heuristic(newPos)})
		}
	}
	return -1
}

func part1(puzzleInput []string) int {
	return solve(puzzleInput, 1024)
}

func part2(puzzleInput []string) string {
	lo, hi := 1025, len(puzzleInput)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if solve(puzzleInput, mid) < 0 {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return puzzleInput[lo]
}
