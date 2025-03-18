package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"runtime"
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
	fmt.Println(solve(puzzleInput))
}

func solve(puzzleInput []string) (string, string) {
	positionCountsMap := make(map[int]map[rune]int, len(puzzleInput[0]))
	for i := range len(puzzleInput[0]) {
		positionCountsMap[i] = make(map[rune]int, 26)
	}
	for _, s := range puzzleInput {
		for i, ch := range s {
			positionCountsMap[i][ch]++
		}
	}
	var sb1 strings.Builder
	var sb2 strings.Builder
	var maxCh, minCh rune
	var maxN, minN int
	for i := range len(puzzleInput[0]) {
		// unexpected order when using `for _, countMap := range positionCountsMap`
		// have to force it to iterate in numeric order
		countMap := positionCountsMap[i]
		maxN = 0
		minN = math.MaxInt
		for ch, n := range countMap {
			if n > maxN {
				maxN = n
				maxCh = ch
			}
			if n < minN {
				minN = n
				minCh = ch
			}
		}
		sb1.WriteRune(maxCh)
		sb2.WriteRune(minCh)
	}
	return sb1.String(), sb2.String()
}
