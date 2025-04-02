package main

import (
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
	elves := parseInput(strings.Split(string(data), "\n\n"))
	fmt.Println(solve(elves))
}

type Elf struct {
	calories []int
}

func (e *Elf) TotalCalories() int {
	res := 0
	for _, cal := range e.calories {
		res += cal
	}
	return res
}

func parseInput(data []string) []*Elf {
	elves := make([]*Elf, len(data))
	for i, s := range data {
		lines := strings.Split(s, "\n")
		calories := make([]int, len(lines))
		for j, line := range lines {
			n, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			calories[j] = n
		}
		elves[i] = &Elf{calories}
	}
	return elves
}

func solve(elves []*Elf) (int, int) {
	var top3 [3]int
	for _, elf := range elves {
		total := elf.TotalCalories()
		for i := range top3 {
			if total > top3[i] {
				copy(top3[i+1:], top3[i:])
				top3[i] = total
				break
			}
		}
	}
	return top3[0], top3[0] + top3[1] + top3[2]
}
