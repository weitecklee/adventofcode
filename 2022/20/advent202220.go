package main

import (
	"container/ring"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n")
	fmt.Println(part1(parseInput(input)))
}

func parseInput(input []string) (*ring.Ring, int) {
	file := ring.New(len(input))
	for _, line := range input {
		n, _ := strconv.Atoi(line)
		file.Value = n
		file = file.Next()
	}
	return file, len(input)
}

func part1(file *ring.Ring, length int) int {
	originalOrder := map[int]*ring.Ring{}
	valueMap := map[int]*ring.Ring{}
	for i := 0; i < length; i++ {
		originalOrder[i] = file
		valueMap[file.Value.(int)] = file
		file = file.Next()
	}
	for i := 0; i < length; i++ {
		prev := originalOrder[i].Prev()
		curr := prev.Unlink(1)
		prev = prev.Move(curr.Value.(int))
		prev.Link(curr)
	}
	fileZero := valueMap[0]
	sum := 0
	for i := 0; i < 3; i++ {
		fileZero = fileZero.Move(1000)
		sum += fileZero.Value.(int)
	}
	return sum
}
