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
	input := strings.Split(string(data), "\n")
	parsedInput := parseInput(input)
	fmt.Println(part1(parsedInput))
}

func parseInput(input []string) map[string][]int {
	re := regexp.MustCompile(`\d+`)
	row := 0
	edges := map[string][]int{}
	for row < len(input) {
		num, _ := strconv.Atoi(re.FindString(input[row]))
		row++
		top := input[row]
		var left strings.Builder
		var right strings.Builder
		for row < len(input) && len(input[row]) > 0 {
			left.WriteString(string(input[row][0]))
			right.WriteString(string(input[row][len(input[row])-1]))
			row++
		}
		bottom := input[row-1]
		for _, edge := range []string{top, left.String(), right.String(), bottom} {
			revEdge := reverse(edge)
			if _, ok := edges[revEdge]; ok {
				edges[revEdge] = append(edges[revEdge], num)
			} else {
				edges[edge] = append(edges[edge], num)
			}
		}
		row++
	}
	return edges
}

func reverse(s string) string {
	rev := []rune(s)
	for i, j := 0, len(rev)-1; i < len(rev)/2; i, j = i+1, j-1 {
		rev[i], rev[j] = rev[j], rev[i]
	}
	return string(rev)
}

func part1(edges map[string][]int) int {
	edgePieces := map[int]int{}
	prod := 1
	for _, tiles := range edges {
		if len(tiles) == 1 {
			edgePieces[tiles[0]]++
			if edgePieces[tiles[0]] == 2 {
				prod *= tiles[0]
			}
		}
	}
	return prod
}
