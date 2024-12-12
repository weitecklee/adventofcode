package main

import (
	"container/ring"
	"fmt"
	"os"
	"path/filepath"
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
	input := strings.Split(string(data), "\n")
	fmt.Println(part1(parseInput(input)))
	fmt.Println(part2(parseInput(input)))
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

func mix(originalOrder *map[int]*ring.Ring, length int) {
	for i := 0; i < length; i++ {
		prev := (*originalOrder)[i].Prev()
		curr := prev.Unlink(1)
		prev = prev.Move(curr.Value.(int) % (length - 1))
		prev.Link(curr)
	}
}

func part1(file *ring.Ring, length int) int {
	originalOrder := map[int]*ring.Ring{}
	valueMap := map[int]*ring.Ring{}
	for i := 0; i < length; i++ {
		originalOrder[i] = file
		valueMap[file.Value.(int)] = file
		file = file.Next()
	}
	mix(&originalOrder, length)
	fileZero := valueMap[0]
	sum := 0
	for i := 0; i < 3; i++ {
		fileZero = fileZero.Move(1000 % length)
		sum += fileZero.Value.(int)
	}
	return sum
}

func part2(file *ring.Ring, length int) int {
	k := 811589153
	originalOrder := map[int]*ring.Ring{}
	valueMap := map[int]*ring.Ring{}
	for i := 0; i < length; i++ {
		file.Value = file.Value.(int) * k
		originalOrder[i] = file
		valueMap[file.Value.(int)] = file
		file = file.Next()
	}
	for j := 0; j < 10; j++ {
		mix(&originalOrder, length)
	}
	fileZero := valueMap[0]
	sum := 0
	for i := 0; i < 3; i++ {
		fileZero = fileZero.Move(1000 % length)
		sum += fileZero.Value.(int)
	}
	return sum
}
