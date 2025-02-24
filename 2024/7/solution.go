package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
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
	puzzleInput := strings.Split((string(data)), "\n")
	equations := parseInput(puzzleInput)
	fmt.Println(solve(equations))
}

type Equation struct {
	result  int
	numbers []int
}

func (e *Equation) canBeTrue() (bool, bool) {
	curr := []int{e.numbers[0]}
	for _, n2 := range e.numbers[1:] {
		curr2 := make([]int, 0, len(curr)*2)
		for _, n := range curr {
			sum := n + n2
			prod := n * n2
			if sum <= e.result {
				curr2 = append(curr2, sum)
			}
			if prod <= e.result {
				curr2 = append(curr2, prod)
			}
		}
		curr = curr2
	}
	for _, n := range curr {
		if n == e.result {
			return true, true
		}
	}
	curr = []int{e.numbers[0]}
	for _, n2 := range e.numbers[1:] {
		curr2 := make([]int, 0, len(curr)*3)
		for _, n := range curr {
			sum := n + n2
			prod := n * n2
			concatString := strconv.Itoa(n) + strconv.Itoa(n2)
			if sum <= e.result {
				curr2 = append(curr2, sum)
			}
			if prod <= e.result {
				curr2 = append(curr2, prod)
			}
			if concat, err := strconv.Atoi(concatString); err != nil {
				panic(err)
			} else if concat <= e.result {
				curr2 = append(curr2, concat)
			}
		}
		curr = curr2
	}
	for _, n := range curr {
		if n == e.result {
			return false, true
		}
	}
	return false, false
}

func parseInput(puzzleInput []string) []Equation {
	equations := make([]Equation, len(puzzleInput))
	numRegex := regexp.MustCompile(`\d+`)
	for i, line := range puzzleInput {
		matches := numRegex.FindAllString(line, -1)
		nums := make([]int, len(matches))
		for j, s := range matches {
			if n, err := strconv.Atoi(s); err != nil {
				panic(err)
			} else {
				nums[j] = n
			}
		}
		equations[i] = Equation{nums[0], nums[1:]}
	}
	return equations
}

func solve(equations []Equation) (int, int) {
	part1 := 0
	part2 := 0
	for _, eq := range equations {
		withoutConcat, withConcat := eq.canBeTrue()
		if withoutConcat {
			part1 += eq.result
			part2 += eq.result
		} else if withConcat {
			part2 += eq.result
		}
	}
	return part1, part2
}
