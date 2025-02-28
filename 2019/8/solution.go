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
	puzzleInput := parseInput(string(data), 25, 6)
	fmt.Println(part1(puzzleInput))
	fmt.Println(part2(puzzleInput))
}

func parseInput(data string, w, h int) [][][]int {
	var image [][][]int
	var layer [][]int
	var row []int
	for i := 0; i*w*h < len(data); i++ {
		layer = make([][]int, h)
		for r := range h {
			row = make([]int, w)
			for c := range w {
				row[c] = int(data[i*w*h+r*w+c] - '0')
			}
			layer[r] = row
		}
		image = append(image, layer)
	}
	return image
}

func part1(image [][][]int) int {
	res := 0
	fewestZeros := math.MaxInt
	var counts map[int]int
	for _, layer := range image {
		counts = make(map[int]int)
		for _, row := range layer {
			for _, n := range row {
				counts[n]++
			}
		}
		if counts[0] < fewestZeros {
			fewestZeros = counts[0]
			res = counts[1] * counts[2]
		}
	}
	return res
}

func part2(image [][][]int) string {
	h := len(image[0])
	w := len(image[0][0])
	colorMap := map[int]rune{0: ' ', 1: '█'}
	var sb strings.Builder
	var i, color int
	for r := range h {
		for c := range w {
			color = 2
			i = 0
			for color == 2 {
				color = image[i][r][c]
				i++
			}
			sb.WriteRune(colorMap[color])
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}
