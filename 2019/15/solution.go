package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"strconv"
	"strings"

	"github.com/weitecklee/adventofcode/2019/intcode"
)

var directions = [][2]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}

func main() {
	_, filename, _, _ := runtime.Caller(0)
	dirname := filepath.Dir(filename)
	inputFilePath := filepath.Join(dirname, "input.txt")
	data, err := os.ReadFile(inputFilePath)
	if err != nil {
		panic(err)
	}
	puzzleInput := parseInput(strings.Split(string(data), ","))
	fmt.Println(solve(puzzleInput))
}

func parseInput(data []string) []int {
	numbers := make([]int, 0, len(data))
	for _, s := range data {
		if n, err := strconv.Atoi(s); err != nil {
			panic(err)
		} else {
			numbers = append(numbers, n)
		}
	}
	return numbers
}

type QueueEntry struct {
	pos        [2]int
	ch         chan int
	cmdHistory []int
}

type QueueEntry2 struct {
	pos   [2]int
	steps int
}

func cloneDroid(puzzleInput, cmdHistory []int) chan int {
	ch := make(chan int)
	droid := intcode.NewIntcodeProgram(puzzleInput, ch)
	go droid.Run()
	for _, cmd := range cmdHistory {
		<-ch
		ch <- cmd
		<-ch
	}
	return ch
}

func reverse(dirIndex int) int {
	if dirIndex < 3 {
		return 3 - dirIndex
	}
	return 7 - dirIndex
}

func solve(puzzleInput []int) (int, int) {
	var part1, part2 int
	ch := make(chan int)
	droid := intcode.NewIntcodeProgram(puzzleInput, ch)
	go droid.Run()

	var queue []QueueEntry
	queue = append(queue, QueueEntry{[2]int{0, 0}, ch, []int{}})
	visited := make(map[[2]int]int)
	visited[[2]int{0, 0}] = 1

	var oxygenPos [2]int

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		pos := item.pos
		ch := item.ch
		cmdHistory := item.cmdHistory
		candidates := make([]int, 0, 4)

		for i, dir := range directions {
			pos2 := [2]int{pos[0] + dir[0], pos[1] + dir[1]}
			if _, ok := visited[pos2]; ok {
				continue
			}
			<-ch
			ch <- i + 1
			resp := <-ch
			visited[pos2] = resp
			if resp == 1 {
				candidates = append(candidates, i)
				<-ch
				ch <- reverse(i + 1)
				<-ch
			} else if resp == 2 {
				oxygenPos = pos2
				part1 = len(cmdHistory) + 1
			}
		}

		if len(candidates) == 0 {
			<-ch
			ch <- intcode.STOPSIGNAL
			continue
		}

		for _, i := range candidates[1:] {
			pos2 := [2]int{pos[0] + directions[i][0], pos[1] + directions[i][1]}
			ch2 := cloneDroid(puzzleInput, cmdHistory)
			<-ch2
			ch2 <- i + 1
			<-ch2
			queue = append(queue, QueueEntry{pos2, ch2, slices.Concat(cmdHistory, []int{i + 1})})
		}
		i := candidates[0]
		<-ch
		ch <- i + 1
		<-ch
		pos2 := [2]int{pos[0] + directions[i][0], pos[1] + directions[i][1]}
		queue = append(queue, QueueEntry{pos2, ch, append(cmdHistory, i+1)})
	}

	queue2 := []QueueEntry2{{oxygenPos, 0}}

	for len(queue2) > 0 {
		item := queue2[0]
		queue2 = queue2[1:]
		pos := item.pos
		steps := item.steps
		if steps > part2 {
			part2 = steps
		}
		for _, dir := range directions {
			pos2 := [2]int{pos[0] + dir[0], pos[1] + dir[1]}
			if visited[pos2] != 0 {
				visited[pos2] = 0
				queue2 = append(queue2, QueueEntry2{pos2, steps + 1})
			}
		}
	}

	return part1, part2
}
