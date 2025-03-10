package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
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
	puzzleInput := strings.Split(string(data), "\n")
	illegalCharacters, incompletes := parseInput(puzzleInput)
	fmt.Println(part1(illegalCharacters))
	fmt.Println(part2(incompletes))
}

var bracketMatches = map[rune]rune{
	'{': '}',
	'[': ']',
	'<': '>',
	'(': ')',
}

var illegalScores = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var incompleteScores = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

func parseInput(puzzleInput []string) ([]rune, [][]rune) {
	var illegalCharacters []rune
	var incompletes [][]rune
	for _, line := range puzzleInput {
		if ok, illegalCh, leftovers := isComplete(line); !ok {
			if leftovers != nil {
				incompletes = append(incompletes, leftovers)
			} else {
				illegalCharacters = append(illegalCharacters, illegalCh)
			}
		}
	}
	return illegalCharacters, incompletes
}

func isComplete(line string) (bool, rune, []rune) {
	brackets := make([]rune, len(line)/2)
	i := -1
	for _, ch := range line {
		if _, ok := bracketMatches[ch]; ok {
			i++
			brackets[i] = ch
		} else if bracketMatches[brackets[i]] == ch {
			i--
		} else {
			return false, ch, nil
		}
	}
	if i >= 0 {
		return false, 0, brackets[:i+1]
	}
	return true, 0, nil
}

func part1(illegalCharacters []rune) int {
	res := 0
	for _, ch := range illegalCharacters {
		res += illegalScores[ch]
	}
	return res
}

func part2(incompleteLines [][]rune) int {
	res := make([]int, len(incompleteLines))
	var score int
	var closer rune
	for k, line := range incompleteLines {
		score = 0
		for i := len(line) - 1; i >= 0; i-- {
			score *= 5
			closer = bracketMatches[line[i]]
			score += incompleteScores[closer]
		}
		res[k] = score
	}
	sort.Ints(res)
	return res[len(res)/2]
}
